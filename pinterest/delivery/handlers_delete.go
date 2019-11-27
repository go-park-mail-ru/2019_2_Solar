package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	user_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/user-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo/v4"
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
	_, err = h.AuthSessManager.LogoutUser(ctx2, &cookie)
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

func (h *HandlersStruct) ServiceDeleteSubscribe(ctx echo.Context) (Err error) {
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
	if _, err := h.UserService.DeleteSubscribe(context.Background(),
		&user_service.UserIDAndFolloweeUsername{
			UserID: user.ID,
			FolloweeUsername: followeeName}); err != nil {
		return err
	}
	body := struct {
		Info string `json:"info"`
	}{"data successfully saved"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string), body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}
