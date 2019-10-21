package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"strconv"
)

func (h *HandlersStruct) HandleGetProfileUserData(ctx echo.Context) (Err error) {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	user := ctx.Get("User")
	if user == nil {
		return errors.New("not authorized")
	}
	data := h.PUsecase.SetJsonData(user.(models.User), "OK")

	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleEditProfileUserData(ctx echo.Context) (Err error) {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	user := getUser.(models.User)

	decoder := json.NewDecoder(ctx.Request().Body)

	newUserProfile := new(models.EditUserProfile)

	if err := decoder.Decode(newUserProfile); err != nil {
		return err
	}

	if err := h.PUsecase.EditProfileDataValidationCheck(newUserProfile); err != nil {
		return err
	}
	if err := h.PUsecase.EditUsernameEmailIsUnique(newUserProfile.Username, newUserProfile.Email, user.Username, user.Email, user.ID); err != nil {
		return err
	}

	editStrings, err := h.PUsecase.SetUser(*newUserProfile, user);
	if err != nil {
		return err
	}
	if editStrings != 1 {
		return errors.New("several notes edit")
	}

	data := h.PUsecase.SetJsonData(nil, "data successfully saved")

	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleEditProfileUserPicture(ctx echo.Context) (Err error) {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	if err := ctx.Request().ParseMultipartForm(5 * 1024 * 1024); err != nil {
		return err
	}
	user := getUser.(models.User)
	file, header, err := ctx.Request().FormFile("profilePicture")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	defer func() {
		if err := file.Close(); err != nil {
			Err = err
		}
	}()
	var buf bytes.Buffer
	tee := io.TeeReader(file, &buf)
	fileHash, err := h.PUsecase.CalculateMD5FromFile(tee)
	if err != nil {
		return err
	}
	if err = h.PUsecase.CreateDir("static/picture/" + fileHash[:2]); err != nil {
		return err
	}
	formatFile, err := h.PUsecase.ExtractFormatFile(header.Filename)
	if err != nil {
		return err
	}
	fileName := "static/picture/" + fileHash[:2] + "/" + fileHash + formatFile
	if err = h.PUsecase.CreatePictureFile(fileName, &buf); err != nil {
		return
	}
	if _, err := h.PUsecase.SetUserAvatarDir(strconv.Itoa(int(user.ID)), fileName); err != nil {
		return err
	}
	data := h.PUsecase.SetJsonData(nil, "profile picture has been successfully saved")
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}
