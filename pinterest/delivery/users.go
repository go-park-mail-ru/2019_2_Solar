package delivery

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) HandleGetUserByEmail(ctx echo.Context) (Err error) {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	email := ctx.Param("email")
	if email == "" {
		return errors.New("incorrect email")
	}

	userProfile, err := h.PUsecase.ReadUserStructByEmail(email)
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
