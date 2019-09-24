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

type UserSession struct {
	ID           uint64 `json:"id"`
	UserID       uint64 `json:"id"`
	SessionValue string `json:"session"`
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

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}

type Handlers struct {
	users    []User
	sessions []UserSession
	mu       *sync.Mutex
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
		w.Write([]byte("{}"))
		return
	}

	for _, user := range h.users {
		if user.Email == newUserReg.Email {
			log.Printf("not unique Email")
			w.Write([]byte(`{"errorMessage":"not unique Email"}`))
			return
		}
	}

	fmt.Println(newUserReg)

	h.mu.Lock()

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

	h.users = append(h.users, newUser)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(newUser)
	h.mu.Unlock()
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	expiration := time.Now().Add(100 * time.Hour)
	value, err := rand.Int(rand.Reader, big.NewInt(80))
	sessionValue := int((value.Int64()))
	if err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte("{}"))
		return
	}
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   strconv.Itoa(sessionValue),
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	if len(h.sessions) > 0 {
		id = h.sessions[len(h.sessions)-1].ID + 1
	}

	newUserSession := UserSession{
		ID:           id,
		UserID:       newUser.ID,
		SessionValue: strconv.Itoa(sessionValue),
	}
	h.sessions = append(h.sessions, newUserSession)

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

	http.ListenAndServe(":8080", nil)
}
