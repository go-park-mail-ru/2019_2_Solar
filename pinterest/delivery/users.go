package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/pkg/errors"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) HandleGetUserByUsername(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	username := ctx.Param("username")
	if username == "" {
		return errors.New("incorrect name")
	}

	userProfile, err := h.PUsecase.GetUserByUsername(username)
	if err != nil {
		return err
	}

	data := h.PUsecase.SetJsonData(userProfile, "OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleCreateSubscribe(ctx echo.Context) (Err error){
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
	followeeName := ctx.Param("username")
	if err := h.PUsecase.AddSubscribe(string(user.ID), followeeName); err != nil {
		return err
	}
	info := "data successfully saved"
	jsonStruct := models.JSONResponse{Body: info}
	if err := ctx.JSON(200, jsonStruct); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleDeleteSubscribe(ctx echo.Context) (Err error) {

	return nil
}