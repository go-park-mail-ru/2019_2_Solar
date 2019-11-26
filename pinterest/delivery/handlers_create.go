package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
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

	cookie, err := h.AuthSessManager.Client.RegUser(ctx2, &sUserReg)
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
	cookie, err := h.AuthSessManager.Client.LoginUser(ctx2, &userLogin)
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

