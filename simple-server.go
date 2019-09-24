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
	ID        uint64 `json:"-"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	Status    string `json:"status"`
	AvatarDir string `json:"-"`
	IsActive  string `json:"isactive"`
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

func CreateNewUserSession(h *Handlers, w http.ResponseWriter, user User) error {

	expiration := time.Now().Add(100 * time.Hour)
	value, err := rand.Int(rand.Reader, big.NewInt(80))
	if err != nil {
		return err
	}
	sessionValue := int((value.Int64()))
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   strconv.Itoa(sessionValue),
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	var id uint64 = 0
	if len(h.sessions) > 0 {
		id = h.sessions[len(h.sessions)-1].ID + 1
	}

	newUserSession := UserSession{
		ID:     id,
		UserID: user.ID,
		UserCookie: UserCookie{
			Value:      strconv.Itoa(sessionValue),
			Expiration: expiration,
		},

		//SessionValue: strconv.Itoa(*sessionValue),
	}
	h.sessions = append(h.sessions, newUserSession)
	return nil
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

	h.mu.Lock()
	if err := CreateNewUserSession(h, w, newUser); err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
	}
	h.mu.Unlock()

	return
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
	user, ok := value.(User)
	if !ok {
		log.Printf("email was not found")
		w.Write([]byte(`{"errorMessage":"incorrect combination of Email and Password"}`))
		return
	} else if user.Password != newUserLogin.Password {
		log.Printf("incorrect password")
			w.Write([]byte(`{"errorMessage":"incorrect combination of Email and Password"}`))
		return
	}
	if err := CreateNewUserSession(h, w, user); err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
	}
	h.mu.Unlock()
	return
}

func (h *Handlers) HandleEditProfileUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newProfileUser := new(User)
	err := decoder.Decode(newProfileUser)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"incorrect json"}`))
		return
	}

	fmt.Println(newProfileUser)
	idUser, err := FindIdUserByCookie()
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		w.Write([]byte(`{"errorMessage":"invalid cookie"}`))
	}
	h.mu.Lock()
	value := SearchUserByIdUser(h.users, idUser)
	user, ok := value.(User)
	if !ok {
		log.Printf("email was not found")
		w.Write([]byte(`{"errorMessage":"incorrect combination of Email and Password"}`))
		return
	}
	SaveNewProfileUser(user, newProfileUser)
	h.mu.Unlock()
	w.Write([]byte(`{message":"data successfully saved"}`))
	return
}

func FindIdUserByCookie() (uint64, error) {
	return 1, nil
}

func SearchUserByIdUser(users []User, idUser uint64) interface{} {
	for _, user := range users {
		if user.ID == idUser {
			return user
		}
	}
	return ""
}

func SaveNewProfileUser(user *User, newUser User) {
	user.Age = newUser.Age
	user.Email = newUser.Email
	user.Name = newUser.Name
	user.Password = newUser.Password
	user.Status = newUser.Status
	user.Surname = newUser.Surname
	user.Username = newUser.Username
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

	http.HandleFunc("/profile/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			handlers.HandleEditProfileUser(w, r)
			return
		}

		handlers.HandleEmpty(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
