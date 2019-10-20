package delivery

import (
	"encoding/json"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) HandleListUsers(ctx echo.Context) error {
	var err error
	defer func() {
		if bodyCloseError := ctx.Request().Body.Close(); bodyCloseError != nil {
			err = bodyCloseError
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	users, err := h.PUsecase.GetAllUsers()
	if err != nil {
		return err
	}

	data := h.PUsecase.SetJsonData(users, "OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}