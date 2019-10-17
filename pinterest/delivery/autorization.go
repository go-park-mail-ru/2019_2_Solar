package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (h *HandlersStruct) HandleRegUser(ctx echo.Context) (Err error) {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")

	//fmt.Println(ctx.Get("User"))
	if ctx.Get("User") != nil {
		return nil
	}

	encoder := json.NewEncoder(ctx.Response())
	decoder := json.NewDecoder(ctx.Request().Body)

	newUserReg := new(models.UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return err
	}
	fmt.Println(newUserReg.Password)
	if err := h.PUsecase.RegDataValidationCheck(newUserReg); err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, err.Error(), err)
		return err
	}

	if check, err := h.PUsecase.RegUsernameIsUnique(newUserReg.Username); err != nil || !check {
		fmt.Println(err)
		h.PUsecase.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return err
	}

	if check, err := h.PUsecase.RegEmailIsUnique(newUserReg.Email); err != nil || !check {
		h.PUsecase.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return err
	}
	newUserId, err := h.PUsecase.InsertNewUser(newUserReg.Username, newUserReg.Email, newUserReg.Password)
	if err != nil {
		return err
	}

	cookies, err := h.PUsecase.CreateNewUserSession(newUserId)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", err)
		return err
	}
	ctx.SetCookie(&cookies)
	//http.SetCookie(ctx.Response(), &cookies)
	data := h.PUsecase.SetJsonData(newUserReg, "OK")
	err = encoder.Encode(data)
	if err != nil {
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleLoginUser(ctx echo.Context) error {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			panic(err)
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())

	fmt.Println(ctx.Get("User"))
	if user := ctx.Get("User"); user != nil {
		data := h.PUsecase.SetJsonData(user.(models.User), "OK")
		err := encoder.Encode(data)
		if err != nil {
			h.PUsecase.SetResponseError(encoder, "bad user struct", err)
			return err
		}
		return nil
	}
	decoder := json.NewDecoder(ctx.Request().Body)

	newUserLogin := new(models.UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return nil
	}
	var User models.User
	User, err = h.PUsecase.ReadUserStructByEmail(newUserLogin.Email)
	if err != nil {
		return err
	}
	if User.Password != newUserLogin.Password { //Добавить функцию хеша от пароля
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Password"))
		return nil
	}

	cookies, err := h.PUsecase.CreateNewUserSession(strconv.Itoa(int(User.ID)))
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", err)
		return err
	}
	ctx.SetCookie(&cookies)
	data := h.PUsecase.SetJsonData(User, "OK")
	err = encoder.Encode(data)
	if err != nil {
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return err
	}
	return nil
}

/*func (h *HandlersStruct) HandleLogoutUser(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	ctx.Response().Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(ctx.Response())

	sessionKey, err := h.PUsecase.SearchCookie(ctx.Request())
	if err == http.ErrNoCookie {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Cookie has not found", err)
		return nil
	}

	err = h.PUsecase.DeleteOldUserSession(sessionKey.Value)

	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Session has not found", err)
		return nil
	}
	sessionKey.Path = "/"
	sessionKey.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(ctx.Response(), sessionKey)

	data := h.PUsecase.SetJsonData(nil, "Session has been successfully deleted")
	encoder.Encode(data)
	return nil
}*/
