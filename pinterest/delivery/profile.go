package delivery

import (
	"2019_2_Solar/pkg/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (h *Handlers) HandleGetProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	h.Mu.Lock()
	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "invalid cookie or user", err)
		return
	}

	h.Mu.Lock()
	data := h.PUsecase.SetJsonData(h.Users[h.PUsecase.GetUserIndexByID(h.Users, idUser)], "OK")
	h.Mu.Unlock()

	err = encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleGetProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	h.Mu.Lock()
	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "invalid cookie or user", err)
		return
	}
	h.Mu.Lock()
	filename := h.Users[h.PUsecase.GetUserIndexByID(h.Users, idUser)].AvatarDir
	h.Mu.Unlock()
	openFile, err := os.Open(filename)
	defer openFile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		w.WriteHeader(http.StatusNotFound)
		h.PUsecase.SetResponseError(encoder, "file not found", err)
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

	newProfileUser := new(models.EditUserProfile)
	err := decoder.Decode(newProfileUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return
	}

	if err := h.PUsecase.EditProfileDataCheck(newProfileUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, err.Error(), err)
		return
	}

	defer h.Mu.Unlock()
	h.Mu.Lock()

	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "invalid cookie or user", err)
		return
	}
	if !h.PUsecase.EditEmailIsUnique(h.Users, newProfileUser.Email, idUser) {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return
	}
	if !h.PUsecase.EditUsernameIsUnique(h.Users, newProfileUser.Username, idUser) {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return
	}

	h.PUsecase.SaveNewProfileUser(&h.Users[h.PUsecase.GetUserIndexByID(h.Users, idUser)], newProfileUser)

	data := h.PUsecase.SetJsonData(nil, "data successfully saved")
	encoder.Encode(data)
}

func (h *Handlers) HandleEditProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	r.ParseMultipartForm(5 * 1024 * 1025)
	h.Mu.Lock()
	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "user not found or not valid cookies", err)
		return
	}
	file, header, err := r.FormFile("profilePicture")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Cannot read profile picture", err)
		return
	}

	defer file.Close()
	formatFile, err := h.PUsecase.ExtractFormatFile(header.Filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Cannot read profile picture", err)
		return
	}
	fileName := strconv.FormatUint(idUser, 10) + "_picture" + formatFile
	newFile, err := os.Create(fileName)
	h.Mu.Lock()
	h.Users[h.PUsecase.GetUserIndexByID(h.Users, idUser)].AvatarDir = fileName
	h.Mu.Unlock()
	defer newFile.Close()
	_, err = io.Copy(newFile, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "File recording has failed", err)
		return
	}

	data := h.PUsecase.SetJsonData(nil, "profile picture has been successfully saved")
	encoder.Encode(data)
}
