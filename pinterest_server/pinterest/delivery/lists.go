package delivery

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) HandleListUsers(ctx echo.Context) (Err error) {
	var err error
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	users, err := h.PUsecase.GetAllUsers()
	if err != nil {
		return err
	}

	data := h.PUsecase.SetJSONData(users, ctx.Get("token").(string),"OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}
