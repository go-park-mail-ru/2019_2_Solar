package handls

import (
	"2019_2_Solar/pkg/structs"
	"log"
	"net/http"
	"sync"
)

var handlers = Handlers{
	Users:    make([]structs.User, 0),
	Sessions: make([]structs.UserSession, 0),
	Mu:       &sync.Mutex{},
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", "http://solar-env.v2zxh2s3me.us-east-2.elasticbeanstalk.com")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
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
