package delivery

import (
	"bytes"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (h *HandlersStruct) HandleGetProfileUserData(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	user := ctx.Get("User")
	if user == nil {
		return errors.New("not authorized")
	}
	data := h.PUsecase.SetJSONData(user.(models.User), ctx.Get("token").(string),"OK")

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

	if err := h.PUsecase.CheckProfileData(newUserProfile); err != nil {
		return err
	}
	if err := h.PUsecase.CheckUsernameEmailIsUnique(newUserProfile.Username, newUserProfile.Email, user.Username, user.Email, user.ID); err != nil {
		return err
	}

	editStrings, err := h.PUsecase.SetUser(*newUserProfile, user)
	if err != nil {
		return err
	}
	if editStrings != 1 {
		return errors.New("several notes edit")
	}

	data := h.PUsecase.SetJSONData(nil, ctx.Get("token").(string),"data successfully saved")

	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleEditProfileUserPicture(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
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
	if err = h.PUsecase.AddDir("static/ava/" + fileHash[:2]); err != nil {
		return err
	}
	formatFile, err := h.PUsecase.ExtractFormatFile(header.Filename)
	if err != nil {
		return err
	}
	fileName := "static/ava/" + fileHash[:2] + "/" + fileHash + formatFile
	if err = h.PUsecase.AddPictureFile(fileName, &buf); err != nil {
		return err
	}
	if _, err := h.PUsecase.SetUserAvatarDir(user.ID, fileName); err != nil {
		return err
	}
	data := h.PUsecase.SetJSONData(nil, ctx.Get("token").(string),"profile picture has been successfully saved")
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}
