package functions

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

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
	return "", errors.New("Invalid file name")
}

func SetResponseError(encoder *json.Encoder, msg string, err error) {
	log.Printf("%s: %s", msg, err)
	data := SetJsonData(nil, msg)
	encoder.Encode(data)
}
