package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type UserCookie struct {
	Value      string    `json:"-"`
	Expiration time.Time `json:"-"`
}

type UserSession struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"id"`
	UserCookie
}

type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserReg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      string `json:"age"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	AvatarDir string `json:"-"`
}

type Handlers struct {
	users    []User
	sessions []UserSession
	mu       *sync.Mutex
}

func CreateNewUser(h *Handlers, newUserReg UserReg) User {
	var id uint64 = 0
	if len(h.users) > 0 {
		id = h.users[len(h.users)-1].ID + 1
	}

	newUser := User{
		ID:       id,
		Name:     "",
		Password: newUserReg.Password,
		Email:    newUserReg.Email,
		Age:      newUserReg.Age,
	}
	return newUser
}

func CreateNewUserSession(h *Handlers, user User, sessionValue *int, expiration time.Time) UserSession {
	var id uint64 = 0
	if len(h.sessions) > 0 {
		id = h.sessions[len(h.sessions)-1].ID + 1
	}

	newUserSession := UserSession{
		ID:     id,
		UserID: user.ID,
		UserCookie: UserCookie{
			Value:      strconv.Itoa(*sessionValue),
			Expiration: expiration,
		},

		//SessionValue: strconv.Itoa(*sessionValue),
	}
	return newUserSession
}

func EmailIsUnique(h *Handlers, newUserReg UserReg) bool {
	for _, user := range h.users {
		if user.Email == newUserReg.Email {
			return false
		}
	}
	return true
}

func SearchUserByEmail(users []User, newUserLogin *UserLogin) interface{} {
	for _, user := range users {
		if user.Email == newUserLogin.Email {
			return user
		}
	}
	return ""
}

func (h *Handlers) HandleEmpty(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Write([]byte("{}"))
	fmt.Println("Empty handler has been done")

	return
}

func (h *Handlers) HandleRegUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newUserReg := new(UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"incorrect json"}`))
		return
	}

	if !EmailIsUnique(h, *newUserReg) {
		log.Printf("not unique Email")
		w.Write([]byte(`{"errorMessage":"not unique Email"}`))
		return
	}

	fmt.Println(newUserReg)

	h.mu.Lock()
	newUser := CreateNewUser(h, *newUserReg)

	h.users = append(h.users, newUser)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(newUser)
	h.mu.Unlock()
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"bad user struct"}`))
		return
	}

	expiration := time.Now().Add(100 * time.Hour)
	value, err := rand.Int(rand.Reader, big.NewInt(80))
	if err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
		return
	}
	sessionValue := int((value.Int64()))
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   strconv.Itoa(sessionValue),
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	h.mu.Lock()

	newUserSession := CreateNewUserSession(h, newUser, &sessionValue, expiration)
	h.sessions = append(h.sessions, newUserSession)

	h.mu.Unlock()

	return
}

func (h *Handlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newUserInput := new(UserInput)
	err := decoder.Decode(newUserInput)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	fmt.Println(newUserInput)
	h.mu.Lock()

	var id uint64 = 0
	if len(h.users) > 0 {
		id = h.users[len(h.users)-1].ID + 1
	}

	h.users = append(h.users, User{
		ID:       id,
		Name:     newUserInput.Name,
		Password: newUserInput.Password,
	})
	h.mu.Unlock()
}

func (h *Handlers) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	h.mu.Lock()
	err := encoder.Encode(h.users)
	h.mu.Unlock()
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}
}

func (h *Handlers) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newUserLogin := new(UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"incorrect json"}`))
		return
	}

	fmt.Println(newUserLogin)

	h.mu.Lock()
	value := SearchUserByEmail(h.users, newUserLogin)
	if user, ok := value.(User); !ok {
		log.Printf("email was not found")
		w.Write([]byte(`{"errorMessage":"incorrect combination of Email and Password"}`))
		return
	} else if user.Password != newUserLogin.Password {
		log.Printf("incorrect password")
		w.Write([]byte(`{"errorMessage":"incorrect combination of Email and Password"}`))
		return
	} else {
		expiration := time.Now().Add(100 * time.Hour)
		value, err := rand.Int(rand.Reader, big.NewInt(80))
		if err != nil {
			log.Printf("error while generating sessionValue: %s", err)
			w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
			return
		}
		sessionValue := int((value.Int64()))
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   strconv.Itoa(sessionValue),
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)

		newUserSession := CreateNewUserSession(h, user, &sessionValue, expiration)
		h.sessions = append(h.sessions, newUserSession)
		w.Write([]byte(`{"infoMessage":"authorization successful"}`))
	}
	h.mu.Unlock()

}

func main() {
	handlers := Handlers{
		users: make([]User, 0),
		mu:    &sync.Mutex{},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			handlers.HandleCreateUser(w, r)
			return
		}

		handlers.HandleListUsers(w, r)
	})

	http.HandleFunc("/registration/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			handlers.HandleRegUser(w, r)
			return
		}

		handlers.HandleEmpty(w, r)
	})

	http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			handlers.HandleLoginUser(w, r)
			return
		}

		handlers.HandleEmpty(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
