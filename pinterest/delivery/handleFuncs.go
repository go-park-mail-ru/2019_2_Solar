package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"log"
	"net/http"
	"sync"
)

type Handlers struct {
	PUsecase pinterest.Usecase
	Users    []models.User
	Sessions []models.UserSession
	Mu       *sync.Mutex
}

var handlers = Handlers{
	PUsecase: &usecase.PinterestUseCase{},
	Users:    make([]models.User, 0),
	Sessions: make([]models.UserSession, 0),
	Mu:       &sync.Mutex{},
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println(r.URL.Path)
	handlers.HandleListUsers(w, r)
}

func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(r.URL.Path)

	if r.Method == http.MethodPost {
		handlers.HandleRegUser(w, r)
		return
	}

	handlers.HandleEmpty(w, r)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(r.URL.Path)

	if r.Method == http.MethodPost {
		handlers.HandleLoginUser(w, r)
		return
	}

	handlers.HandleEmpty(w, r)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(r.URL.Path)

	if r.Method == http.MethodPost {
		handlers.HandleLogoutUser(w, r)
		return
	}

	handlers.HandleEmpty(w, r)
}

func HandleProfileData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(r.URL.Path)

	if r.Method == http.MethodPost {
		handlers.HandleEditProfileUserData(w, r)
		return
	}
	if r.Method == http.MethodGet {
		handlers.HandleGetProfileUserData(w, r)
		return
	}
	handlers.HandleEmpty(w, r)
}

func HandleProfilePicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println(r.URL.Path)

	if r.Method == http.MethodPost {
		handlers.HandleEditProfileUserPicture(w, r)
		return
	}
	if r.Method == http.MethodGet {
		handlers.HandleGetProfileUserPicture(w, r)
		return
	}
	handlers.HandleEmpty(w, r)
}
