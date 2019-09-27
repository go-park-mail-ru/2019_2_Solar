package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type UserCookie struct {
	Value      string    `json:"-"`
	Expiration time.Time `json:"-"`
}

type UserSession struct {
	ID     uint64 `json:"-"`
	UserID uint64 `json:"-"`
	UserCookie
}

type UserReg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditUserProfile struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      string `json:"age"`
	Status   string `json:"status"`
	IsActive string `json:"isactive"`
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
		Username: newUserReg.Username,
	}
	return newUser
}

func CreateNewUserSession(h *Handlers, user User) (interface{}, error) {

	expiration := time.Now().Add(100 * time.Hour)
	value, err := rand.Int(rand.Reader, big.NewInt(80))
	if err != nil {
		return nil, err
	}
	sessionValue := int((value.Int64()))
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   strconv.Itoa(sessionValue),
		Path:    "/",
		Expires: expiration,
	}

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
	return cookie, nil
}

func DeleteOldUserSession(h *Handlers, value string) error {
	for i, session := range h.sessions {
		if session.Value == value {
			h.sessions = append(h.sessions[:i], h.sessions[i+1:]...)
			return nil
		}
	}
	err := errors.New("session has not found")
	return err
}

func SearchCookieSession(r *http.Request) (*http.Cookie, error) {
	session, err := r.Cookie("session_id")
	return session, err
}

func EmailIsUnique(h *Handlers, email string) bool {
	for _, user := range h.users {
		if user.Email == email {
			return false
		}
	}
	return true
}

func UsernameIsUnique(h *Handlers, username string) bool {
	for _, user := range h.users {
		if user.Username == username {
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

func GetUserIndexByID(h *Handlers, id uint64) int {
	for index, user := range h.users {
		if user.ID == id {
			return index
		}
	}
	return -1
}
func (h *Handlers) HandleEmpty(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Write([]byte("{}"))
	fmt.Println("Empty handler has been done")

	return
}

func (h *Handlers) HandleRegUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newUserReg := new(UserReg)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(newUserReg)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"incorrect json"}`))
		return
	}

	defer h.mu.Unlock()
	h.mu.Lock()
	if !EmailIsUnique(h, newUserReg.Email) {
		log.Printf("not unique Email")
		w.Write([]byte(`{"errorMessage":"not unique Email"}`))
		return
	}

	fmt.Println(newUserReg)

	newUser := CreateNewUser(h, *newUserReg)
	h.users = append(h.users, newUser)
	cookie, err := CreateNewUserSession(h, newUser)
	if err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
		return
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
	}
	http.SetCookie(w, &correctCookie)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(newUser)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"bad user struct"}`))
		return
	}

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
	defer h.mu.Unlock()
	h.mu.Lock()
	value := SearchUserByEmail(h.users, newUserLogin)
	user, ok := value.(User)
	if !ok {
		log.Printf("email was not found")
		w.Write([]byte(`{"errorMessage":"incorrect combination of Email and Password"}`))
		return
	}
	if user.Password != newUserLogin.Password {
		log.Printf("incorrect password")
		w.Write([]byte(`{"errorMessage":"incorrect combination of Email and Password"}`))
		return
	}
	idUser, err := SearchIdUserByCookie(r, h)
	fmt.Println(idUser)
	if err == nil {
		//log.Printf("Invalid cookie: %s", err)
		//w.Write([]byte(`{"errorMessage":"invalid cookie or user"}`))
		encoder := json.NewEncoder(w)
		err = encoder.Encode(user)
		if err != nil {
			log.Printf("error while marshalling JSON: %s", err)
			w.Write([]byte(`{"errorMessage":"bad user struct"}`))
			return
		}
		w.Write([]byte(`{"message":"successfully log in yet"}`))
		return
	}
	cookie, err := CreateNewUserSession(h, user)
	if err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
		return
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		log.Printf("error while generating sessionValue: %s", err)
		w.Write([]byte(`{"errorMessage":"error while generating sessionValue"}`))
	}
	http.SetCookie(w, &correctCookie)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(user)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"bad user struct"}`))
		return
	}
	return
}

//Проверено
func (h *Handlers) HandleEditProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newProfileUser := new(EditUserProfile)
	err := decoder.Decode(newProfileUser)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte(`{"errorMessage":"incorrect json"}`))
		return
	}

	defer h.mu.Unlock()
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		w.Write([]byte(`{"errorMessage":"invalid cookie or user"}`))
		return
	}
	if !EmailIsUnique(h, newProfileUser.Email) {
		log.Printf("not unique Email")
		w.Write([]byte(`{"errorMessage":"not unique Email"}`))
		return
	}
	if !UsernameIsUnique(h, newProfileUser.Username) {
		log.Printf("not unique Username")
		w.Write([]byte(`{"errorMessage":"not unique Username"}`))
		return
	}
	SaveNewProfileUser(&h.users[GetUserIndexByID(h, idUser)], newProfileUser)

	w.Write([]byte(`{"message":"data successfully saved"}`))
	return
}

func SearchIdUserByCookie(r *http.Request, h *Handlers) (uint64, error) {
	idSessionString, err := SearchCookieSession(r)
	if err == http.ErrNoCookie {
		return 0, errors.New("cookies not found")
	}
	fmt.Println(idSessionString)
	for _, oneSession := range h.sessions {
		if oneSession.UserCookie.Value == idSessionString.Value {
			return oneSession.UserID, err
		}
	}
	return 0, errors.New("idUser not found")
}

func SaveNewProfileUser(user *User, newUser *EditUserProfile) {
	if newUser.Age != "" {
		user.Age = newUser.Age
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Name != "" {
		user.Name = newUser.Name
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
	}
	if newUser.Status != "" {
		user.Status = newUser.Status
	}
	if newUser.Surname != "" {
		user.Surname = newUser.Surname
	}
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
}

//Проверено
func (h *Handlers) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	session, err := SearchCookieSession(r)
	if err == http.ErrNoCookie {
		w.Write([]byte(`{"errorMessage":"Cookies have not found"}`))
		return
	}
	h.mu.Lock()
	err = DeleteOldUserSession(h, session.Value)
	if err != nil {
		h.mu.Unlock()
		w.Write([]byte(`{"errorMessage":"Session has not found"}`))
		return
	}
	h.mu.Unlock()
	session.Path = "/"
	session.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(w, session)
	w.Write([]byte(`{"infoMessage":"Session has been successfully deleted"}`))
	return
}

func (h *Handlers) HandleEditProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	/*	session, err := SearchCookieSession(r)
		if err == http.ErrNoCookie {
			w.Write([]byte(`{"errorMessage":"Cookies have not found"}`))
			return
		}*/
	r.ParseMultipartForm(5 * 1024 * 1025)
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	h.mu.Unlock()
	if err != nil {
		w.Write([]byte(`{"errorMessage":"user not found or not valid cookies"}`))
		return
	}
	//Header not used
	file, header, err := r.FormFile("profilePicture")
	if err != nil {
		w.Write([]byte(`{"errorMessage":"Cannot read profile picture"}`))
		return
	}

	defer file.Close()
	formatFile, err := ExtractFormatFile(header.Filename)
	if err != nil {
		w.Write([]byte(`{"errorMessage"` + err.Error() + "}"))
		return
	}
	newFile, err := os.Create(strconv.FormatUint(idUser, 10) + "_picture" + formatFile)
	defer newFile.Close()
	io.Copy(newFile, file)
	w.Write([]byte(`{"Message":"profile picture has been successfully saved"}`))
	return
}

func ExtractFormatFile(FileName string) (string, error) {
	for i := 0; i < len(FileName); i++ {
		if string(FileName[i]) == "." {
			return FileName[i:], nil
		}
	}
	return "", errors.New("Invalid file name")
}

/*func (h *Handlers) HandleCookies(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	expiration := time.Now().Add(100 * time.Hour)
	cookie := http.Cookie{
		Name:    "ses_id",
		Value:   "qwert",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("qwerty"))
}*/

// ================================= Handler functions =================================

var handlers = Handlers{
	users:    make([]User, 0),
	sessions: make([]UserSession, 0),
	mu:       &sync.Mutex{},
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

	handlers.HandleEmpty(w, r)
}

func HandleProfilePicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(r.URL.Path)

	if r.Method == http.MethodPost {
		handlers.HandleEditProfileUserPicture(w, r)
		return
	}

	handlers.HandleEmpty(w, r)
}

func main() {
	
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/users/", HandleUsers)
	http.HandleFunc("/registration/", HandleRegistration)
	http.HandleFunc("/login/", HandleLogin)
	http.HandleFunc("/logout/", HandleLogout)
	http.HandleFunc("/profile/data", HandleProfileData)
	http.HandleFunc("/profile/picture", HandleProfilePicture)

	/*	http.HandleFunc("/cookies/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		if r.Method == http.MethodPost {
			handlers.HandleCookies(w, r)
			return
		}

		handlers.HandleEmpty(w, r)
	})*/

	http.ListenAndServe(":8080", nil)
}
