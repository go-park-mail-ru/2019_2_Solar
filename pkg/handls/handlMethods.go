package handls

import (
	"2019_2_Solar/pkg/functions"
	"2019_2_Solar/pkg/structs"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Handlers struct {
	Users    []structs.User
	Sessions []structs.UserSession
	Mu       *sync.Mutex
}

func (h *Handlers) HandleEmpty(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	data := functions.SetJsonData(nil, "Empty handler has been done")
	encoder.Encode(data)
	log.Printf("Empty handler has been done")
}

func (h *Handlers) HandleRegUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	newUserReg := new(structs.UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		functions.SetResponseError(encoder, "incorrect json", err)
		return
	}

	if err := functions.EmailCheck(newUserReg.Email); err != nil {
		functions.SetResponseError(encoder, "incorrect Email", err)
		return
	}
	if err := functions.UsernameCheck(newUserReg.Username); err != nil {
		functions.SetResponseError(encoder, "incorrect Username", err)
		return
	}
	if err := functions.PasswordCheck(newUserReg.Password); err != nil {
		functions.SetResponseError(encoder, err.Error(), err)
		return
	}

	defer h.Mu.Unlock()

	h.Mu.Lock()
	if !functions.RegEmailIsUnique(h.Users, newUserReg.Email) {
		functions.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return
	}
	if !functions.RegUsernameIsUnique(h.Users, newUserReg.Username) {
		functions.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return
	}

	newUser := functions.CreateNewUser(h.Users, *newUserReg)
	h.Users = append(h.Users, newUser)
	cookies, newSession, err := functions.CreateNewUserSession(h.Sessions, newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	if len(cookies) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "error while generating sessionValue", errors.New("incorrect while create session"))
		return
	}
	http.SetCookie(w, &cookies[0])
	h.Sessions = append(h.Sessions, newSession)

	data := functions.SetJsonData(newUser, "OK")
	err = encoder.Encode(data)
	if err != nil {
		functions.SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	h.Mu.Lock()
	data := functions.SetJsonData(h.Users, "OK")
	h.Mu.Unlock()

	err := encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "error while marshalling JSON", err)
		return
	}
}

func (h *Handlers) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	newUserLogin := new(structs.UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "incorrect json", err)
		return
	}

	defer h.Mu.Unlock()
	h.Mu.Lock()
	value := functions.SearchUserByEmail(h.Users, newUserLogin)
	user, ok := value.(structs.User)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Email"))
		return
	}
	if user.Password != newUserLogin.Password {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Password"))
		return
	}

	//Если пришли валидные куки, значит новую сессию не создаём
	idUser, err := functions.SearchIdUserByCookie(r, h.Sessions)
	fmt.Println(idUser)
	if err == nil {
		data := functions.SetJsonData(user, "Successfully log in yet")
		encoder.Encode(data)
		return
	}

	cookies, newSession, err := functions.CreateNewUserSession(h.Sessions, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	if len(cookies) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "error while generating sessionValue", errors.New("incorrect while create session"))
		return
	}
	http.SetCookie(w, &cookies[0])
	h.Sessions = append(h.Sessions, newSession)

	data := functions.SetJsonData(user, "OK")

	err = encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleGetProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	h.Mu.Lock()
	idUser, err := functions.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		functions.SetResponseError(encoder, "invalid cookie or user", err)
		return
	}

	h.Mu.Lock()
	data := functions.SetJsonData(h.Users[functions.GetUserIndexByID(h.Users, idUser)], "OK")
	h.Mu.Unlock()

	err = encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleGetProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	h.Mu.Lock()
	idUser, err := functions.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		functions.SetResponseError(encoder, "invalid cookie or user", err)
		return
	}
	h.Mu.Lock()
	filename := h.Users[functions.GetUserIndexByID(h.Users, idUser)].AvatarDir
	h.Mu.Unlock()
	openFile, err := os.Open(filename)
	defer openFile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		w.WriteHeader(http.StatusNotFound)
		functions.SetResponseError(encoder, "file not found", err)
		return
	}
	//File is found, create and send the correct headers
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	openFile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := openFile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)
	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	openFile.Seek(0, 0)
	io.Copy(w, openFile) //'Copy' the file to the client
}

func (h *Handlers) HandleEditProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	newProfileUser := new(structs.EditUserProfile)
	err := decoder.Decode(newProfileUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "incorrect json", err)
		return
	}

	if err := functions.EditProfileDataCheck(newProfileUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, err.Error(), err)
		return
	}

	defer h.Mu.Unlock()
	h.Mu.Lock()

	idUser, err := functions.SearchIdUserByCookie(r, h.Sessions)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		functions.SetResponseError(encoder, "invalid cookie or user", err)
		return
	}
	if !functions.EditEmailIsUnique(h.Users, newProfileUser.Email, idUser) {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return
	}
	if !functions.EditUsernameIsUnique(h.Users, newProfileUser.Username, idUser) {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return
	}

	functions.SaveNewProfileUser(&h.Users[functions.GetUserIndexByID(h.Users, idUser)], newProfileUser)

	data := functions.SetJsonData(nil, "data successfully saved")
	encoder.Encode(data)
}

func (h *Handlers) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	sessionKey, err := functions.SearchCookie(r)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "Cookie has not found", err)
		return
	}

	h.Mu.Lock()
	err = functions.DeleteOldUserSession(&h.Sessions, sessionKey.Value)
	h.Mu.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "Session has not found", err)
		return
	}
	sessionKey.Path = "/"
	sessionKey.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(w, sessionKey)

	data := functions.SetJsonData(nil, "Session has been successfully deleted")
	encoder.Encode(data)
}

func (h *Handlers) HandleEditProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	r.ParseMultipartForm(5 * 1024 * 1025)
	h.Mu.Lock()
	idUser, err := functions.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		functions.SetResponseError(encoder, "user not found or not valid cookies", err)
		return
	}
	file, header, err := r.FormFile("profilePicture")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "Cannot read profile picture", err)
		return
	}

	defer file.Close()
	formatFile, err := functions.ExtractFormatFile(header.Filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "Cannot read profile picture", err)
		return
	}
	fileName := strconv.FormatUint(idUser, 10) + "_picture" + formatFile
	newFile, err := os.Create(fileName)
	h.Mu.Lock()
	h.Users[functions.GetUserIndexByID(h.Users, idUser)].AvatarDir = fileName
	h.Mu.Unlock()
	defer newFile.Close()
	_, err = io.Copy(newFile, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		functions.SetResponseError(encoder, "File recording has failed", err)
		return
	}

	data := functions.SetJsonData(nil, "profile picture has been successfully saved")
	encoder.Encode(data)
}
