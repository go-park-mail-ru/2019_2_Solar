package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (h *Handlers) HandleEmpty(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	data := SetJsonData(nil, "Empty handler has been done")
	encoder.Encode(data)
	log.Printf("Empty handler has been done")
	return
}

func (h *Handlers) HandleRegUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	newUserReg := new(UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		SetResponseError(encoder, "incorrect json", err)
		return
	}

	defer h.mu.Unlock()

	h.mu.Lock()
	if !EmailIsUnique(h, newUserReg.Email) {
		SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return
	}

	newUser := CreateNewUser(h, *newUserReg)
	h.users = append(h.users, newUser)
	cookie, err := CreateNewUserSession(h, newUser)
	if err != nil {
		SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	http.SetCookie(w, &correctCookie)

	data := SetJsonData(newUser, "OK")
	err = encoder.Encode(data)
	if err != nil {
		SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	h.mu.Lock()
	//users := ForUsersBodyJSON{h.users}
	data := SetJsonData(h.users, "OK")
	err := encoder.Encode(data)
	h.mu.Unlock()
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		data := SetJsonData(nil, "error while marshalling JSON")
		encoder.Encode(data)
		//w.Write([]byte("{}"))
		return
	}
}

func (h *Handlers) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	infMsg := ""

	newUserLogin := new(UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		data := SetJsonData(nil, "incorrect json")
		encoder.Encode(data)
		return
	}

	defer h.mu.Unlock()
	h.mu.Lock()
	value := SearchUserByEmail(h.users, newUserLogin)
	user, ok := value.(User)
	if !ok {
		log.Printf("email was not found")
		data := SetJsonData(nil, "incorrect combination of Email and Password")
		encoder.Encode(data)
		return
	}
	if user.Password != newUserLogin.Password {
		log.Printf("incorrect password")
		data := SetJsonData(nil, "incorrect combination of Email and Password")
		encoder.Encode(data)
		return
	}
	//Если пришли валидные куки, значит новую сессию не создаём
	idUser, err := SearchIdUserByCookie(r, h)
	fmt.Println(idUser)
	if err == nil {
		infMsg = "successfully log in yet"
		data := SetJsonData(user, infMsg)
		encoder.Encode(data)
		return
	}
	cookie, err := CreateNewUserSession(h, user)
	if err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		data := SetJsonData(nil, "error while generating sessionValue")
		encoder.Encode(data)
		return
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		log.Printf("error while generating sessionValue: %s", err)
		data := SetJsonData(nil, "error while generating sessionValue")
		encoder.Encode(data)
	}
	http.SetCookie(w, &correctCookie)

	data := SetJsonData(user, infMsg)

	err = encoder.Encode(data)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		data := SetJsonData(nil, "bad user struct")
		encoder.Encode(data)
		return
	}
	return
}

func (h *Handlers) HandleGetProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	idUser, err := SearchIdUserByCookie(r, h)
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		data := SetJsonData(nil, "invalid cookie or user")
		encoder.Encode(data)
		return
	}
	infMsg := ""
	data := SetJsonData(h.users[GetUserIndexByID(h, idUser)], infMsg)

	err = encoder.Encode(data)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		data := SetJsonData(nil, "bad user struct")
		encoder.Encode(data)
		return
	}
	return
}

func (h *Handlers) HandleGetProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	idUser, err := SearchIdUserByCookie(r, h)
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		data := SetJsonData(nil, "invalid cookie or user")
		encoder.Encode(data)
		return
	}
	filename := strconv.FormatUint(idUser, 10) + "_picture" + ".jpg"

	openfile, err := os.Open(filename)
	defer openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}
	//File is found, create and send the correct headers
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)
	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	openfile.Seek(0, 0)
	io.Copy(w, openfile) //'Copy' the file to the client
	return
}

func (h *Handlers) HandleEditProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	newProfileUser := new(EditUserProfile)
	err := decoder.Decode(newProfileUser)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		data := SetJsonData(nil, "incorrect json")
		encoder.Encode(data)
		return
	}

	defer h.mu.Unlock()
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		data := SetJsonData(nil, "invalid cookie or user")
		encoder.Encode(data)
		return
	}
	if !EmailIsUnique(h, newProfileUser.Email) {
		log.Printf("not unique Email")
		data := SetJsonData(nil, "not unique Email")
		encoder.Encode(data)
		return
	}
	if !UsernameIsUnique(h, newProfileUser.Username) {
		log.Printf("not unique Username")
		data := SetJsonData(nil, "not unique Username")
		encoder.Encode(data)
		return
	}
	SaveNewProfileUser(&h.users[GetUserIndexByID(h, idUser)], newProfileUser)

	data := SetJsonData(nil, "data successfully saved")
	encoder.Encode(data)
	return
}

//Проверено
func (h *Handlers) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	session, err := SearchCookieSession(r)
	if err == http.ErrNoCookie {
		data := SetJsonData(nil, "Cookies have not found")
		encoder.Encode(data)
		return
	}
	h.mu.Lock()
	err = DeleteOldUserSession(h, session.Value)
	if err != nil {
		h.mu.Unlock()
		data := SetJsonData(nil, "Session has not found")
		encoder.Encode(data)
		return
	}
	h.mu.Unlock()
	session.Path = "/"
	session.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(w, session)
	data := SetJsonData(nil, "Session has been successfully deleted")
	encoder.Encode(data)
	//w.Write([]byte(`{"infoMessage":"Session has been successfully deleted"}`))
	return
}

func (h *Handlers) HandleEditProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	r.ParseMultipartForm(5 * 1024 * 1025)
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	h.mu.Unlock()
	if err != nil {
		encoder := json.NewEncoder(w)
		data := SetJsonData(nil, "user not found or not valid cookies")
		encoder.Encode(data)
		//w.Write([]byte(`{"errorMessage":"user not found or not valid cookies"}`))
		return
	}
	file, header, err := r.FormFile("profilePicture")
	if err != nil {
		encoder := json.NewEncoder(w)
		data := SetJsonData(nil, "Cannot read profile picture")
		encoder.Encode(data)
		//w.Write([]byte(`{"errorMessage":"Cannot read profile picture"}`))
		return
	}

	defer file.Close()
	formatFile, err := ExtractFormatFile(header.Filename)
	if err != nil {
		w.Write([]byte(`{"errorMessage"` + err.Error() + "}"))
		return
	}
	newFile, err := os.Create(strconv.FormatUint(idUser, 10) + "_picture" + formatFile)
	defer newFile.Close()
	io.Copy(newFile, file)
	w.Write([]byte(`{"Message":"profile picture has been successfully saved"}`))
	return
}
