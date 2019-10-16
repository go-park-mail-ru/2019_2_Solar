package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func (h *Handlers) HandleRegUser(ctx echo.Context) error {
	defer func() {
		err:= ctx.Request().Body.Close()
		if err!= nil {
			panic(err)
		}
	}()

	ctx.Response().Header().Set("Content-Type", "application/json")
	DBWorker := repository.DataBaseWorker{}
	DBWorker.NewDataBaseWorker()

	fmt.Println(ctx.Get("User"))
	if ctx.Get("User") != nil {
		return nil
	}

	encoder := json.NewEncoder(ctx.Response())
	decoder := json.NewDecoder(ctx.Request().Body)

	newUserReg := new(models.UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return nil
	}

	if err := h.PUsecase.RegDataCheck(newUserReg); err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, err.Error(), err)
		return nil
	}

	if !h.PUsecase.RegEmailIsUnique(newUserReg.Email) {
		h.PUsecase.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return nil
	}
	if !h.PUsecase.RegUsernameIsUnique(newUserReg.Username) {
		h.PUsecase.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return nil
	}

	newUser := h.PUsecase.CreateNewUser(newUserReg)

	cookies, _, err := h.PUsecase.CreateNewUserSession(newUser)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", err)
		return nil
	}
	if len(cookies) < 1 {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", errors.New("incorrect while create session"))
		return nil
	}
	http.SetCookie(ctx.Response(), &cookies[0])

	data := h.PUsecase.SetJsonData(newUser, "OK")
	err = encoder.Encode(data)
	if err != nil {
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return nil
	}
	return nil
}

func (h *Handlers) HandleLoginUser(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	ctx.Response().Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(ctx.Request().Body)
	encoder := json.NewEncoder(ctx.Response())

	newUserLogin := new(models.UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return nil
	}

	value := h.PUsecase.SearchUserByEmail(newUserLogin)
	user, ok := value.(models.User)
	if !ok {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Email"))
		return nil
	}
	if user.Password != newUserLogin.Password {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Password"))
		return nil
	}

	//Если пришли валидные куки, значит новую сессию не создаём
	idUser, err := h.PUsecase.SearchIdUserByCookie(ctx.Request())
	fmt.Println(idUser)
	if err == nil {
		data := h.PUsecase.SetJsonData(user, "Successfully log in yet")
		encoder.Encode(data)
		return nil
	}

	cookies, _, err := h.PUsecase.CreateNewUserSession(user)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", err)
		return nil
	}
	if len(cookies) < 1 {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", errors.New("incorrect while create session"))
		return nil
	}
	http.SetCookie(ctx.Response(), &cookies[0])

	data := h.PUsecase.SetJsonData(user, "OK")

	err = encoder.Encode(data)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return nil
	}
	return nil
}

func (h *Handlers) HandleLogoutUser(ctx echo.Context) error {
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
