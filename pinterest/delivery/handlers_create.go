package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

func (h *HandlersStruct) ServiceRegUser(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	if ctx.Get("User") != nil {
		return errors.New("registration with valid cookie")
	}
	encoder := json.NewEncoder(ctx.Response())
	decoder := json.NewDecoder(ctx.Request().Body)

	newUserReg := new(models.UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		return err
	}
	sUserReg := services.UserReg{
		Email: newUserReg.Email,
		Username: newUserReg.Username,
		Password: newUserReg.Password,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	//serviceCtx := context.WithValue(context.Background(), "userReg", newUserReg)
	ctx2 := context.Background()

	cookie, err := h.AuthSessManager.RegUser(ctx2, &sUserReg)
	if err != nil {
		return err
	}

	cookies := new(http.Cookie)
	cookies.Name = "session_key"
	cookies.Value = cookie.Value
	cookies.Path = "/"
	cookies.Expires = time.Now().Add(365 * 24 * time.Hour)

	ctx.SetCookie(cookies)
	data := h.PUsecase.SetJSONData(newUserReg, "", "OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) ServiceLoginUser(ctx echo.Context) (Err error) {
	var err error
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	if user := ctx.Get("User"); user != nil {
		data := h.PUsecase.SetJSONData(user.(models.User), ctx.Get("token").(string),"OK")
		err := encoder.Encode(data)
		if err != nil {
			return err
		}
		return nil
	}
	decoder := json.NewDecoder(ctx.Request().Body)
	newUserLogin := new(models.UserLogin)
	if err := decoder.Decode(newUserLogin); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	userLogin := services.UserLogin{
		Email:               newUserLogin.Email,
		Password:             newUserLogin.Password,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	ctx2 := context.Background()
	cookie, err := h.AuthSessManager.LoginUser(ctx2, &userLogin)
	if err != nil {
		return err
	}

	cookies := new(http.Cookie)
	cookies.Name = "session_key"
	cookies.Value = cookie.Value
	cookies.Path = "/"
	cookies.Expires = time.Now().Add(365 * 24 * time.Hour)

	ctx.SetCookie(cookies)
	data := h.PUsecase.SetJSONData("", "","OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) ServiceCreateBoard(ctx echo.Context) (Err error) {
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

	decoder := json.NewDecoder(ctx.Request().Body)

	newBoard := new(models.NewBoard)

	if err := decoder.Decode(newBoard); err != nil {
		return err
	}

	newBoardMessage := pinboard_service.NewBoard{
		OwnerID: 				user.ID,
		Title:                newBoard.Title,
		Description:          newBoard.Description,
		Category:             newBoard.Category,
	}

	lastID, err := h.PinBoardService.CreateBoard(context.Background(), &newBoardMessage)
	if err != nil {
		return err
	}

	board := models.Board{
		ID: 		lastID.LastID,
		OwnerID:     user.ID,
		Title:       newBoard.Title,
		Description: newBoard.Description,
		Category:    newBoard.Category,
		CreatedTime: time.Now(),
		IsDeleted: false,
	}

	data := struct {
		CSRFToken string `json:"csrf_token"`
		Body      struct {
			Board models.Board `json:"board"`
			Info  string       `json:"info"`
		} `json:"body"`
	}{CSRFToken: ctx.Get("token").(string),
		Body: struct {
			Board models.Board `json:"board"`
			Info  string       `json:"info"`
		}{Info: "data successfully saved", Board: board}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) ServiceCreatePin(ctx echo.Context) (Err error) {
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

	newPinMessage := pinboard_service.NewPin{
		UserID:					user.ID,
		BoardID:              newPin.BoardID,
		Title:                newPin.Title,
		Description:          newPin.Description,
		PinDir:               newPin.PinDir,
	}

	pinMessage, err := h.PinBoardService.CreatePin(context.Background(), &newPinMessage)
	if err != nil {
		return err
	}

	pin := models.Pin{
		ID:          pinMessage.LastID,
		OwnerID:     user.ID,
		AuthorID:    user.ID,
		BoardID:     newPin.BoardID,
		PinDir:      newPin.PinDir,
		Title:       newPin.Title,
		Description: newPin.Description,
		CreatedTime: time.Time{},
		IsDeleted:   false,
	}

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

