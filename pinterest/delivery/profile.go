package delivery

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (h *Handlers) HandleGetProfileUserData(ctx echo.Context) error {
	r := ctx.Request()
	w := ctx.Response()
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	h.Mu.Lock()
	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "invalid cookie or user", err)
		return nil
	}

	h.Mu.Lock()
	data := h.PUsecase.SetJsonData(h.Users[h.PUsecase.GetUserIndexByID(h.Users, idUser)], "OK")
	h.Mu.Unlock()

	err = encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return nil
	}
	return nil
}

func (h *Handlers) HandleGetProfileUserPicture(ctx echo.Context) error {
	r := ctx.Request()
	w := ctx.Response()
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	h.Mu.Lock()
	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "invalid cookie or user", err)
		return nil
	}
	h.Mu.Lock()
	filename := h.Users[h.PUsecase.GetUserIndexByID(h.Users, idUser)].AvatarDir
	h.Mu.Unlock()
	openFile, err := os.Open(filename)
	defer openFile.Close() //Close after function return nil
	if err != nil {
		//File not found, send 404
		w.WriteHeader(http.StatusNotFound)
		h.PUsecase.SetResponseError(encoder, "file not found", err)
		return nil
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
	return nil
}

func (h *Handlers) HandleEditProfileUserData(ctx echo.Context) error {
	r := ctx.Request()
	w := ctx.Response()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	newProfileUser := new(models.EditUserProfile)
	err := decoder.Decode(newProfileUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return nil
	}

	if err := h.PUsecase.EditProfileDataCheck(newProfileUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, err.Error(), err)
		return nil
	}

	defer h.Mu.Unlock()
	h.Mu.Lock()

	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "invalid cookie or user", err)
		return nil
	}
	if !h.PUsecase.EditEmailIsUnique(h.Users, newProfileUser.Email, idUser) {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return nil
	}
	if !h.PUsecase.EditUsernameIsUnique(h.Users, newProfileUser.Username, idUser) {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return nil
	}

	h.PUsecase.SaveNewProfileUser(&h.Users[h.PUsecase.GetUserIndexByID(h.Users, idUser)], newProfileUser)

	data := h.PUsecase.SetJsonData(nil, "data successfully saved")
	encoder.Encode(data)
	return nil
}

func (h *Handlers) HandleEditProfileUserPicture(ctx echo.Context) error {
	r := ctx.Request()
	w := ctx.Response()
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	r.ParseMultipartForm(5 * 1024 * 1025)
	h.Mu.Lock()
	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	h.Mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.PUsecase.SetResponseError(encoder, "user not found or not valid cookies", err)
		return nil
	}
	file, header, err := r.FormFile("profilePicture")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Cannot read profile picture", err)
		return nil
	}

	defer file.Close()
	formatFile, err := h.PUsecase.ExtractFormatFile(header.Filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Cannot read profile picture", err)
		return nil
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
		return nil
	}

	data := h.PUsecase.SetJsonData(nil, "profile picture has been successfully saved")
	encoder.Encode(data)
	return nil
}