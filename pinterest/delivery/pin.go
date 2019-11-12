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
	err = h.PUsecase.AddTags(pin.Description, lastID)
	if err != nil {
		return err
	}
	pin.ID = lastID
	pin.IsDeleted = false

	data := struct {
		CSRFToken string `json:"csrf_token"`
		Body struct {
			Pin  models.Pin `json:"pin"`
			Info string     `json:"info"`
		} `json:"body"`
	}{	CSRFToken: ctx.Get("token").(string),
		Body: struct {
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
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	id := ctx.Param("id")
	if id == "" {
		return errors.New("incorrect id")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	pin, err := h.PUsecase.GetPin(uint64(intId))
	if err != nil {
		return err
	}
	comments, err := h.PUsecase.GetComments(uint64(intId))
	if err != nil {
		return err
	}
	body := struct {
		Pin  models.FullPin `json:"pins"`
		Comments []models.CommentDisplay `json:"comments"`
		Info  string     `json:"info"`
	}{pin, comments ,"OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
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
	ctx.Response().Header().Set("Content-Type", "application/jsonStruct")
	var pins []models.PinDisplay
	pins, err := h.PUsecase.GetNewPins()
	if err != nil {
		return err
	}
	body := struct {
		Pins  []models.PinDisplay `json:"pins"`
		Info  string     `json:"info"`
	}{pins, "OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleGetMyPins(ctx echo.Context) (Err error) {
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
	var pins []models.PinDisplay
	pins, err := h.PUsecase.GetMyPins(user.ID)
	if err != nil {
		return err
	}
	body := struct {
		Pins  []models.PinDisplay `json:"pins"`
		Info  string     `json:"info"`
	}{pins, "OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
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
	var pins []models.PinDisplay
	pins, err := h.PUsecase.GetSubscribePins(user.ID)
	if err != nil {
		return err
	}
	body := struct {
		Pins  []models.PinDisplay `json:"pins"`
		Info  string     `json:"info"`
	}{pins, "OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleCreateComment(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	decoder := json.NewDecoder(ctx.Request().Body)
	ctx.Response().Header().Set("Content-Type", "application/json")
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	user := getUser.(models.User)
	pinID := ctx.Param("id")

	newComment := new(models.NewComment)

	if err := decoder.Decode(newComment); err != nil {
		return err
	}
	comment := models.NewComment{Text: newComment.Text}
	intPinId, err := strconv.Atoi(pinID)
	if err != nil {
		return err
	}
	if err := h.PUsecase.AddComment(uint64(intPinId), user.ID, comment); err != nil {
		return err
	}
	body := struct {
		Info  string     `json:"info"`
	}{"data successfully saved"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}
