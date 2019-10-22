package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"time"
)

func (h *HandlersStruct) HandleCreatePin(ctx echo.Context) (Err error) {
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

	newPin := new(models.NewPin)

	if err := decoder.Decode(newPin); err != nil {
		return err
	}

	if err := h.PUsecase.CheckPinData(*newPin); err != nil {
		return err
	}
	pin := models.Pin{
		OwnerID: user.ID,
		AuthorID: user.ID,
		BoardID: newPin.BoardID,
		Title: newPin.Title,
		Description: newPin.Description,
		PinDir: newPin.PinDir,
		CreatedTime: time.Now(),
	}
	lastID, err := h.PUsecase.AddPin(pin)
	if err != nil {
		return err
	}
	pin.ID = lastID
	pin.IsDeleted = false

	data := struct {
		Body struct {
			Pin models.Pin `json:"pin"`
			Info string `json:"info"`
		} `json:"body"`
	}{Body: struct {
		Pin models.Pin `json:"pin"`
		Info string `json:"info"`
	}{Info: "data successfully saved", Pin: pin}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) HandleCreatePinPicture(ctx echo.Context) (Err error) {
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

	file, header, err := ctx.Request().FormFile("pinPicture")
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
	if err = h.PUsecase.AddDir("static/pin/" + fileHash[:2]); err != nil {
		return err
	}
	formatFile, err := h.PUsecase.ExtractFormatFile(header.Filename)
	if err != nil {
		return err
	}
	fileName := "static/pin/" + fileHash[:2] + "/" + fileHash + formatFile
	if err = h.PUsecase.AddPictureFile(fileName, &buf); err != nil {
		return err
	}

	data := struct {
		Body struct {
			PinDir string `json:"pin_dir"`
			Info string `json:"info"`
		} `json:"body"`
	}{Body: struct {
		PinDir string `json:"pin_dir"`
		Info string `json:"info"`
	}{Info: "pin picture has been successfully saved", PinDir: fileName}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
