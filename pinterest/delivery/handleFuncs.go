package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"log"
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

func HandleRoot(ctx echo.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.Response().Write([]byte("{123}"))
	return nil
}

func HandleUsers(ctx echo.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json")
	log.Println(ctx.Request().URL.Path)
	handlers.HandleListUsers(ctx.Response(), ctx.Request())
	return nil
}

/*func HandleRegistration(w http.ResponseWriter, r *http.Request) {
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
*/