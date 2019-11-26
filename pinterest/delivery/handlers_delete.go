package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func (h *HandlersStruct) ServiceLogoutUser(ctx echo.Context) (Err error) {
	var err error
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	sessionKey, err := ctx.Request().Cookie("session_key")
	if err != nil {
		return err
	}



	cookie := services.Cookie{
		Key:                   sessionKey.Name,
		Value:                sessionKey.Value,
	}

	ctx2 := context.Background()
	_, err = h.AuthSessManager.Client.LogoutUser(ctx2, &cookie)
	if err != nil {
		return err
	}

	sessionKey.Path = "/"
	sessionKey.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(ctx.Response(), sessionKey)

	data := h.PUsecase.SetJSONData(nil, ctx.Get("token").(string),"Session has been successfully deleted")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}