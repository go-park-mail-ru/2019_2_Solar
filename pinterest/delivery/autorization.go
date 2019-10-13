package delivery

import (
	"2019_2_Solar/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (h *Handlers) HandleRegUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	newUserReg := new(models.UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return
	}

	if err := h.PUsecase.RegDataCheck(newUserReg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, err.Error(), err)
		return
	}

	defer h.Mu.Unlock()

	h.Mu.Lock()
	if !h.PUsecase.RegEmailIsUnique(h.Users, newUserReg.Email) {
		h.PUsecase.SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return
	}
	if !h.PUsecase.RegUsernameIsUnique(h.Users, newUserReg.Username) {
		h.PUsecase.SetResponseError(encoder, "not unique Username", errors.New("not unique Username"))
		return
	}

	newUser := h.PUsecase.CreateNewUser(h.Users, *newUserReg)
	h.Users = append(h.Users, newUser)
	cookies, newSession, err := h.PUsecase.CreateNewUserSession(h.Sessions, newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	if len(cookies) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", errors.New("incorrect while create session"))
		return
	}
	http.SetCookie(w, &cookies[0])
	h.Sessions = append(h.Sessions, newSession)

	data := h.PUsecase.SetJsonData(newUser, "OK")
	err = encoder.Encode(data)
	if err != nil {
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	newUserLogin := new(models.UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect json", err)
		return
	}

	defer h.Mu.Unlock()
	h.Mu.Lock()
	value := h.PUsecase.SearchUserByEmail(h.Users, newUserLogin)
	user, ok := value.(models.User)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Email"))
		return
	}
	if user.Password != newUserLogin.Password {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "incorrect combination of Email and Password", errors.New("incorrect Password"))
		return
	}

	//Если пришли валидные куки, значит новую сессию не создаём
	idUser, err := h.PUsecase.SearchIdUserByCookie(r, h.Sessions)
	fmt.Println(idUser)
	if err == nil {
		data := h.PUsecase.SetJsonData(user, "Successfully log in yet")
		encoder.Encode(data)
		return
	}

	cookies, newSession, err := h.PUsecase.CreateNewUserSession(h.Sessions, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	if len(cookies) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while generating sessionValue", errors.New("incorrect while create session"))
		return
	}
	http.SetCookie(w, &cookies[0])
	h.Sessions = append(h.Sessions, newSession)

	data := h.PUsecase.SetJsonData(user, "OK")

	err = encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	sessionKey, err := h.PUsecase.SearchCookie(r)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Cookie has not found", err)
		return
	}

	h.Mu.Lock()
	err = h.PUsecase.DeleteOldUserSession(&h.Sessions, sessionKey.Value)
	h.Mu.Unlock()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "Session has not found", err)
		return
	}
	sessionKey.Path = "/"
	sessionKey.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(w, sessionKey)

	data := h.PUsecase.SetJsonData(nil, "Session has been successfully deleted")
	encoder.Encode(data)
}
