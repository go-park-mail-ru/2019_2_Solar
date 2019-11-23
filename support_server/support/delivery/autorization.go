package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
)

func (h *HandlersStruct) HandleLoginAdmin(ctx echo.Context) (Err error) {
	var err error
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")

	if admin := ctx.Get("Admin"); admin != nil {
		if err := ctx.JSON(400, "already autorized"); err != nil {
			return err
		}
		return nil
	}

	decoder := json.NewDecoder(ctx.Request().Body)
	newAdminAutorize := new(models.AdminAutorize)
	if err := decoder.Decode(newAdminAutorize); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	//var User models.User
	admin, err := h.PUsecase.GetAdminByLogin(newAdminAutorize.Login)
	if err != nil {
		return err
	}

	if err := h.PUsecase.CompareAdminPassword(admin.Password, newAdminAutorize.Password); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	cookies, err := h.PUsecase.AddNewAdminSession(admin.ID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	ctx.SetCookie(&cookies)

	if err := ctx.JSON(200, admin); err != nil {
		return err
	}
	return nil
}
