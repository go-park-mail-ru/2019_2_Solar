package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
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

	//fmt.Println(ctx.Get("User"))
	if ctx.Get("User") != nil {
		return nil
	}

	encoder := json.NewEncoder(ctx.Response())
	decoder := json.NewDecoder(ctx.Request().Body)

	newUserReg := new(models.UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		usecase.SetResponseError(encoder, "incorrect json", err)
		return err
	}

	if err := h.PUsecase.RegDataCheck(newUserReg); err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		usecase.SetResponseError(encoder, err.Error(), err)
		return err
	}

	if check, err := h.PUsecase.RegUsernameIsUnique(newUserReg.Username); err != nil || !check {
		usecase.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return err
	}

	if check, err := h.PUsecase.RegEmailIsUnique(newUserReg.Email); err != nil || !check {
		usecase.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return err
	}
	h.PUsecase.InsertNewUser(newUserReg.Username, newUserReg.Email, newUserReg.Password)

	var str repository.StringSlice
	err = DBWorker.UniversalRead(consts.FindEmailSQLQuery+"'"+newUserReg.Email+"'", &str)
	if err != nil || len(str) > 1 {
		return err
	}
	cookies, err := h.PUsecase.CreateNewUserSession(str[0])
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		usecase.SetResponseError(encoder, "error while generating sessionValue", err)
		return err
	}
	http.SetCookie(ctx.Response(), &cookies)
	data := usecase.SetJsonData(newUserReg, "OK")
	err = encoder.Encode(data)
	if err != nil {
		usecase.SetResponseError(encoder, "bad user struct", err)
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

	fmt.Println(ctx.Get("User"))
	if ctx.Get("User") != nil {
		return nil
	}
	decoder := json.NewDecoder(ctx.Request().Body)
	//encoder := json.NewEncoder(ctx.Response())

	newUserLogin := new(models.UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		//h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return nil
	}
	var User repository.UsersSlice
	DBWorker := repository.DataBaseWorker{}
	DBWorker.NewDataBaseWorker()
	err = DBWorker.UniversalRead(consts.QueryReadUserByEmail+"'"+newUserLogin.Email+"'", &User)
	if err != nil || len(User) != 1 {
		return err
	}
	if User[0].Password != newUserLogin.Password { //Добавить функцию хеша от пароля
		ctx.Response().WriteHeader(http.StatusBadRequest)
		//h.PUsecase.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Password"))
		return nil
	}

	cookies, err := functions.CreateNewUserSession(strconv.Itoa(int(User[0].ID)))
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		//functions.SetResponseError(encoder, "error while generating sessionValue", err)
		return err
	}
	http.SetCookie(ctx.Response(), &cookies)
	//data := h.PUsecase.SetJsonData(newUserReg, "OK")
	/*	err = encoder.Encode(data)
		if err != nil {
			//h.PUsecase.SetResponseError(encoder, "bad user struct", err)
			return err
		}*/
	return nil
}

func (h *HandlersStruct) HandleLogoutUser(ctx echo.Context) error {
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
}
