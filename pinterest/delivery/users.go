package delivery

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) HandleGetUserByEmail(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	email := ctx.Param("email")
	if email == "" {
		return errors.New("incorrect email")
	}

	userProfile, err := h.PUsecase.GetUserByEmail(email)
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
