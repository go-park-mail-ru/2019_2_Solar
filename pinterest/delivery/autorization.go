package delivery

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func (h *HandlersStruct) HandleRegUser(ctx echo.Context) (Err error) {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
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
	if err := h.PUsecase.RegDataValidationCheck(newUserReg); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	//ОБЪЕДЕНИТЬ ФУНКЦИИ ПРОВЕРКИ УНИКАЛЬНОСТИ (2 sql запроса слишком жирно)
	if check, err := h.PUsecase.RegUsernameIsUnique(newUserReg.Username); err != nil || !check {
		data := h.PUsecase.SetJsonData(newUserReg, "Not unique username")
		err = encoder.Encode(data)
		if err != nil {
			return err
		}
		return nil
	}
	if check, err := h.PUsecase.RegEmailIsUnique(newUserReg.Email); err != nil || !check {
		data := h.PUsecase.SetJsonData(newUserReg, "Not unique email")
		err = encoder.Encode(data)
		if err != nil {
			return err
		}
		return nil
	}

	newUserId, err := h.PUsecase.InsertNewUser(newUserReg.Username, newUserReg.Email, newUserReg.Password)
	if err != nil {
		return err
	}

	cookies, err := h.PUsecase.CreateNewUserSession(newUserId)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	ctx.SetCookie(&cookies)
	data := h.PUsecase.SetJsonData(newUserReg, "OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleLoginUser(ctx echo.Context) error {
	var err error
	defer func() {
		if bodyCloseError := ctx.Request().Body.Close(); bodyCloseError != nil {
			err = bodyCloseError
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	if user := ctx.Get("User"); user != nil {
		data := h.PUsecase.SetJsonData(user.(models.User), "OK")
		err := encoder.Encode(data)
		if err != nil {
			return err
		}
		return nil
	}
	decoder := json.NewDecoder(ctx.Request().Body)

	newUserLogin := new(models.UserLogin)
	err = decoder.Decode(newUserLogin)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	var User models.User
	User, err = h.PUsecase.ReadUserStructByEmail(newUserLogin.Email)
	if err != nil {
		return err
	}
	if User.Password != newUserLogin.Password { //Добавить функцию хеша от пароля
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	cookies, err := h.PUsecase.CreateNewUserSession(strconv.Itoa(int(User.ID)))
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	ctx.SetCookie(&cookies)
	data := h.PUsecase.SetJsonData(User, "OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleLogoutUser(ctx echo.Context) error {
	var err error
	defer func() {
		if bodyCloseError := ctx.Request().Body.Close(); bodyCloseError != nil {
			err = bodyCloseError
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	sessionKey, err := ctx.Request().Cookie("session_key")
	if err != nil {
		return err
	}

	err = h.PUsecase.DeleteOldUserSession(sessionKey.Value)
	if err != nil {
		return err
	}
	sessionKey.Path = "/"
	sessionKey.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(ctx.Response(), sessionKey)

	data := h.PUsecase.SetJsonData(nil, "Session has been successfully deleted")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}
