package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func (h *HandlersStruct) HandleRegUser(ctx echo.Context) (Err error) {
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
	if err := h.PUsecase.CheckRegDataValidation(newUserReg); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	if err := h.PUsecase.CheckRegUsernameEmailIsUnique(newUserReg.Username, newUserReg.Email); err != nil {
		return err
	}

	newUserID, err := h.PUsecase.AddNewUser(newUserReg.Username, newUserReg.Email, newUserReg.Password)
	if err != nil {
		return err
	}

	cookies, err := h.PUsecase.AddNewUserSession(newUserID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	ctx.SetCookie(&cookies)
	data := h.PUsecase.SetJSONData(newUserReg, "", "OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleLoginUser(ctx echo.Context) (Err error) {
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
	var User models.User
	User, err = h.PUsecase.GetUserByEmail(newUserLogin.Email)
	if err != nil {
		return err
	}

	if err := h.PUsecase.ComparePassword(User.Password, User.Salt, newUserLogin.Password); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	cookies, err := h.PUsecase.AddNewUserSession(User.ID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	ctx.SetCookie(&cookies)
	data := h.PUsecase.SetJSONData(User, "","OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleLogoutUser(ctx echo.Context) (Err error) {
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

	err = h.PUsecase.RemoveOldUserSession(sessionKey.Value)
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
