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

type DataJSON struct {
	UserJSON  interface{} `json:"user,omitempty"`
	UsersJSON interface{} `json:"users,omitempty"`
	InfoJSON  interface{} `json:"info,omitempty"`
}

type OutJSON struct {
	BodyJSON interface{} `json:"body"`
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
	value, err := rand.Int(rand.Reader, big.NewInt(80)) //Могут повториться. Исправить
	if err != nil {
		return nil, err
	}
	sessionValue := int(value.Int64())
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
	return errors.New("session has not found")
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

func SetJsonData(data interface{}, infMsg string) OutJSON {

	user, ok := data.(User)
	if ok {
		outJSON := OutJSON{
			BodyJSON: DataJSON{
				UserJSON: user,
				InfoJSON: infMsg,
			},
		}
		return outJSON
	}
	if users, ok := data.([]User); ok {

		outJSON := OutJSON{
			BodyJSON: DataJSON{
				UsersJSON: users,
				InfoJSON:  infMsg,
			},
		}
		return outJSON
	}
	outJSON := OutJSON{
		BodyJSON: DataJSON{
			InfoJSON: infMsg,
		},
	}
	return outJSON
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

func ExtractFormatFile(FileName string) (string, error) {
	for i := 0; i < len(FileName); i++ {
		if string(FileName[i]) == "." {
			return FileName[i:], nil
		}
	}
	return "", errors.New("invalid file name")
}

func SetResponseError(encoder *json.Encoder, msg string, err error) {
	log.Printf("%s: %s", msg, err)
	data := SetJsonData(nil, msg)
	encoder.Encode(data)
}

func (h *Handlers) HandleEmpty(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	data := SetJsonData(nil, "Empty handler has been done")
	encoder.Encode(data)
	log.Printf("Empty handler has been done")
	return
}

func (h *Handlers) HandleRegUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	newUserReg := new(UserReg)
	err := decoder.Decode(newUserReg)
	if err != nil {
		SetResponseError(encoder, "incorrect json", err)
		return
	}

	defer h.mu.Unlock()

	h.mu.Lock()
	if !EmailIsUnique(h, newUserReg.Email) {
		SetResponseError(encoder, "not unique Email", errors.New("not unique Email"))
		return
	}

	newUser := CreateNewUser(h, *newUserReg)
	h.users = append(h.users, newUser)
	cookie, err := CreateNewUserSession(h, newUser)
	if err != nil {
		SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		SetResponseError(encoder, "error while generating sessionValue", err)
		return
	}
	http.SetCookie(w, &correctCookie)

	data := SetJsonData(newUser, "OK")
	err = encoder.Encode(data)
	if err != nil {
		SetResponseError(encoder, "bad user struct", err)
		return
	}
}

func (h *Handlers) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	h.mu.Lock()
	data := SetJsonData(h.users, "OK")
	h.mu.Unlock()
	err := encoder.Encode(data)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "error while marshalling JSON")
		encoder.Encode(data)
		return
	}
	return
}

func (h *Handlers) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	infMsg := ""

	newUserLogin := new(UserLogin)
	err := decoder.Decode(newUserLogin)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "incorrect json")
		encoder.Encode(data)
		return
	}

	defer h.mu.Unlock()
	h.mu.Lock()
	value := SearchUserByEmail(h.users, newUserLogin)
	user, ok := value.(User)
	if !ok {
		log.Printf("email was not found")
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "incorrect combination of Email and Password")
		encoder.Encode(data)
		return
	}
	if user.Password != newUserLogin.Password {
		log.Printf("incorrect password")
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "incorrect combination of Email and Password")
		encoder.Encode(data)
		return
	}
	//Если пришли валидные куки, значит новую сессию не создаём
	idUser, err := SearchIdUserByCookie(r, h)
	fmt.Println(idUser)
	if err == nil {
		infMsg = "successfully log in yet"
		data := SetJsonData(user, infMsg)
		encoder.Encode(data)
		return
	}
	cookie, err := CreateNewUserSession(h, user)
	if err != nil {
		log.Printf("error while generating sessionValue: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "error while generating sessionValue")
		encoder.Encode(data)
		return
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		log.Printf("error while generating sessionValue: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "error while generating sessionValue")
		encoder.Encode(data)
	}
	http.SetCookie(w, &correctCookie)

	data := SetJsonData(user, infMsg)

	err = encoder.Encode(data)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "bad user struct")
		encoder.Encode(data)
		return
	}
	return
}

func (h *Handlers) HandleGetProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	h.mu.Unlock()
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		data := SetJsonData(nil, "invalid cookie or user")
		encoder.Encode(data)
		return
	}
	infMsg := ""
	h.mu.Lock()
	data := SetJsonData(h.users[GetUserIndexByID(h, idUser)], infMsg)
	h.mu.Unlock()
	err = encoder.Encode(data)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "bad user struct")
		encoder.Encode(data)
		return
	}
	return
}

func (h *Handlers) HandleGetProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	h.mu.Unlock()
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		data := SetJsonData(nil, "invalid cookie or user")
		encoder.Encode(data)
		return
	}
	filename := strconv.FormatUint(idUser, 10) + "_picture" + ".jpg"

	openFile, err := os.Open(filename)
	defer openFile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		w.WriteHeader(http.StatusNotFound)
		data := SetJsonData(nil, "file not found")
		encoder.Encode(data)
		return
	}
	//File is found, create and send the correct headers
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	openFile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := openFile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)
	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	openFile.Seek(0, 0)
	io.Copy(w, openFile) //'Copy' the file to the client
	return
}

func (h *Handlers) HandleEditProfileUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	newProfileUser := new(EditUserProfile)
	err := decoder.Decode(newProfileUser)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "incorrect json")
		encoder.Encode(data)
		return
	}

	defer h.mu.Unlock()
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	if err != nil {
		log.Printf("Invalid cookie: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		data := SetJsonData(nil, "invalid cookie or user")
		encoder.Encode(data)
		return
	}
	if !EmailIsUnique(h, newProfileUser.Email) {
		log.Printf("not unique Email")
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "not unique Email")
		encoder.Encode(data)
		return
	}
	if !UsernameIsUnique(h, newProfileUser.Username) {
		log.Printf("not unique Username")
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "not unique Username")
		encoder.Encode(data)
		return
	}
	SaveNewProfileUser(&h.users[GetUserIndexByID(h, idUser)], newProfileUser)

	data := SetJsonData(nil, "data successfully saved")
	encoder.Encode(data)
	return
}

func (h *Handlers) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)

	session, err := SearchCookieSession(r)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "Cookies have not found")
		encoder.Encode(data)
		return
	}
	h.mu.Lock()
	err = DeleteOldUserSession(h, session.Value)
	h.mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "Session has not found")
		encoder.Encode(data)
		return
	}
	session.Path = "/"
	session.Expires = time.Now().AddDate(0, 0, -999)
	http.SetCookie(w, session)
	data := SetJsonData(nil, "Session has been successfully deleted")
	encoder.Encode(data)
	return
}

func (h *Handlers) HandleEditProfileUserPicture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	r.ParseMultipartForm(5 * 1024 * 1025)
	h.mu.Lock()
	idUser, err := SearchIdUserByCookie(r, h)
	h.mu.Unlock()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		data := SetJsonData(nil, "user not found or not valid cookies")
		encoder.Encode(data)
		return
	}
	file, header, err := r.FormFile("profilePicture")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := SetJsonData(nil, "Cannot read profile picture")
		encoder.Encode(data)
		return
	}

	defer file.Close()
	formatFile, err := ExtractFormatFile(header.Filename)
	if err != nil {
		data := SetJsonData(nil, err.Error())
		encoder.Encode(data)
		return
	}
	newFile, err := os.Create(strconv.FormatUint(idUser, 10) + "_picture" + formatFile)
	defer newFile.Close()
	io.Copy(newFile, file)
	data := SetJsonData(nil, "profile picture has been successfully saved")
	encoder.Encode(data)
	return
}

// ================================= Handler functions =================================

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

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

func main() {

	http.Handle("/", CORSMiddleware(http.HandlerFunc(HandleRoot)))
	http.Handle("/users/", CORSMiddleware(http.HandlerFunc(HandleUsers)))
	http.Handle("/registration/", CORSMiddleware(http.HandlerFunc(HandleRegistration)))
	http.Handle("/login/", CORSMiddleware(http.HandlerFunc(HandleLogin)))
	http.Handle("/logout/", CORSMiddleware(http.HandlerFunc(HandleLogout)))
	http.Handle("/profile/data", CORSMiddleware(http.HandlerFunc(HandleProfileData)))
	http.Handle("/profile/picture", CORSMiddleware(http.HandlerFunc(HandleProfilePicture)))

	http.ListenAndServe(":8080", nil)
}
