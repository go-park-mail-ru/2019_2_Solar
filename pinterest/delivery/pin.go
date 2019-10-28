package delivery

import (
	"bytes"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (h *HandlersStruct) HandleCreatePin(ctx echo.Context) (Err error) {
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
	user := getUser.(models.User)

	file, header, err := ctx.Request().FormFile("pinPicture")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	jsonPin := ctx.Request().FormValue("pin")
	newPin := new(models.NewPin)

	if err := json.Unmarshal([]byte(jsonPin), newPin); err != nil {
		return err
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

	newPin.PinDir = fileName

	if err := h.PUsecase.CheckPinData(*newPin); err != nil {
		return err
	}
	pin := models.Pin{
		OwnerID:     user.ID,
		AuthorID:    user.ID,
		BoardID:     newPin.BoardID,
		Title:       newPin.Title,
		Description: newPin.Description,
		PinDir:      newPin.PinDir,
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
			Pin  models.Pin `json:"pin"`
			Info string     `json:"info"`
		} `json:"body"`
	}{Body: struct {
		Pin  models.Pin `json:"pin"`
		Info string     `json:"info"`
	}{Info: "data successfully saved", Pin: pin}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) HandleGetPin(ctx echo.Context) (Err error) {
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
	//user := getUser.(models.User)

	id := ctx.Param("id")
	if id == "" {
		return errors.New("incorrect id")
	}
	pinID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	pin, err := h.PUsecase.GetPin(uint64(pinID))
	if err != nil {
		return err
	}

	data := struct {
		Body struct {
			Pin  models.Pin `json:"pin"`
			Info string     `json:"info"`
		} `json:"body"`
	}{Body: struct {
		Pin  models.Pin `json:"pin"`
		Info string     `json:"info"`
	}{Info: "OK", Pin: pin}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) HandleGetNewPins(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	var pins []models.PinForMainPage
	pins, err := h.PUsecase.GetNewPins()
	if err != nil {
		return nil
	}
	json := models.JSONResponse{Body: pins}
	if err := ctx.JSON(200, json); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleGetSubscribePins(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	user := getUser.(models.User)
	var pins []models.PinForMainPage
	pins, err := h.PUsecase.GetNewSubscribePins(user.ID)
	if err != nil {
		return nil
	}
	json := models.JSONResponse{Body: pins}
	if err := ctx.JSON(200, json); err != nil {
		return err
	}
	return nil
}
