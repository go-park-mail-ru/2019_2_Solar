package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest"
	"github.com/labstack/echo"
)

type Handlers struct {
	PUsecase pinterest.Usecase
}

func HandleRoot(ctx echo.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.Response().Write([]byte("{123}"))
	return nil
}

func NewHandlers(e *echo.Echo, uc pinterest.Usecase) {
	handler := &Handlers{
		PUsecase: uc,
			}

	e.GET("/", HandleRoot)

	e.GET("/users/", handler.HandleListUsers)

	e.POST("/registration/", handler.HandleRegUser)
	e.POST("/login/", handler.HandleLoginUser)
	e.POST("/logout/", handler.HandleLogoutUser)

	e.GET("/profile/data", handler.HandleGetProfileUserData)
	e.GET("/profile/picture", handler.HandleGetProfileUserPicture)

	e.POST("/profile/data", handler.HandleEditProfileUserData)
	e.POST("/profile/picture", handler.HandleEditProfileUserPicture)
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