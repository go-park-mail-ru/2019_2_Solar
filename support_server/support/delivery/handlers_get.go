package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)


func (h *HandlersStruct) HandleGetActiveUsers(ctx echo.Context) (Err error) {
	var err error
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")

	if admin := ctx.Get("Admin"); admin == nil {
		if err := ctx.JSON(400, "not autorized"); err != nil {
			return err
		}
		return nil
	}

	activeUsers, err := h.PUsecase.GetHubListActiveUsers()
	if err != nil {
		return err
	}

	var users []models.User
	for i := 0; i < len(activeUsers); i++ {
		user, err := h.PUsecase.GetUserByID(activeUsers[i])
		if err != nil {
			return err
		}
		users = append(users, user)
	}


	if err := ctx.JSON(200, activeUsers); err != nil {
		return err
	}
	return nil
}