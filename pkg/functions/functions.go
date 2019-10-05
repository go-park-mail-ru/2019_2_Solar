package functions

import (
	"2019_2_Solar/pkg/structs"
	"crypto/rand"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

type Handlers struct {
	Users    []structs.User
	Sessions []structs.UserSession
	Mu       *sync.Mutex
}

func CreateNewUser(users []structs.User, newUserReg structs.UserReg) structs.User {
	var id uint64 = 0
	if len(users) > 0 {
		id = users[len(users)-1].ID + 1
	}

	newUser := structs.User{
		ID:       id,
		Name:     "",
		Password: newUserReg.Password,
		Email:    newUserReg.Email,
		Username: newUserReg.Username,
	}
	return newUser
}

func CreateNewUserSession(sessions []structs.UserSession, user structs.User) ([]http.Cookie, structs.UserSession, error) {
	cookies := []http.Cookie{}

	var sessionValue uint64 = 0
	if len(sessions) > 0 {
		sessionValue = sessions[len(sessions)-1].ID + 1
	}

	sessionKey := GenSessionKey(12)

	cookieSessionKey := http.Cookie{
		Name:    "session_key",
		Value:   sessionKey,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	cookies = append(cookies, cookieSessionKey)

	newUserSession := structs.UserSession{
		ID:     sessionValue,
		UserID: user.ID,
		UserCookie: structs.UserCookie{
			Value:      sessionKey,
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	return cookies, newUserSession, nil
}

func DeleteOldUserSession(sessions *([]structs.UserSession), value string) error {
	for i, session := range *sessions {
		if session.Value == value {
			*sessions = append((*sessions)[:i], (*sessions)[i+1:]...)
			return nil
		}
	}
	return errors.New("session has not found")
}

func SearchCookie(r *http.Request) (*http.Cookie, error) {
	key, err := r.Cookie("session_key")
	return key, err
}

func RegEmailIsUnique(users []structs.User, email string) bool {
	for _, user := range users {
		if user.Email == email {
			return false
		}
	}
	return true
}

func RegUsernameIsUnique(users []structs.User, username string) bool {
	for _, user := range users {
		if user.Username == username {
			return false
		}
	}
	return true
}

func EditEmailIsUnique(users []structs.User, email string, idUser uint64) bool {
	for _, user := range users {
		if user.Email == email {
			if user.ID != idUser {
				return false
			}
		}
	}
	return true
}

func EditUsernameIsUnique(users []structs.User, username string, idUser uint64) bool {
	for _, user := range users {
		if user.Username == username {
			if user.ID != idUser {
				return false
			}
		}
	}
	return true
}

func SearchUserByEmail(users []structs.User, newUserLogin *structs.UserLogin) interface{} {
	for _, user := range users {
		if user.Email == newUserLogin.Email {
			return user
		}
	}
	return ""
}

func GetUserIndexByID(users []structs.User, id uint64) int {
	for index, user := range users {
		if user.ID == id {
			return index
		}
	}
	return -1
}

func SetJsonData(data interface{}, infMsg string) structs.OutJSON {
	user, ok := data.(structs.User)
	if ok {
		outJSON := structs.OutJSON{
			BodyJSON: structs.DataJSON{
				UserJSON: user,
				InfoJSON: infMsg,
			},
		}
		return outJSON
	}
	if users, ok := data.([]structs.User); ok {

		outJSON := structs.OutJSON{
			BodyJSON: structs.DataJSON{
				UsersJSON: users,
				InfoJSON:  infMsg,
			},
		}
		return outJSON
	}
	outJSON := structs.OutJSON{
		BodyJSON: structs.DataJSON{
			InfoJSON: infMsg,
		},
	}
	return outJSON
}

func SearchIdUserByCookie(r *http.Request, sessions []structs.UserSession) (uint64, error) {
	sessionKey, err := SearchCookie(r)
	if err == http.ErrNoCookie {
		return 0, errors.New("cookies not found")
	}

	for _, oneSession := range sessions {
		if oneSession.Value == sessionKey.Value {
			return oneSession.UserID, nil
		}
	}
	return 0, errors.New("idUser not found")
}

func SaveNewProfileUser(user *structs.User, newUser *structs.EditUserProfile) {
	user.Age = newUser.Age
	user.Status = newUser.Status
	user.Name = newUser.Name
	user.Surname = newUser.Surname

	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
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

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 62 possibilities
	letterIdxBits = 6                                                                // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                             // All 1-bits, as many as letterIdxBits
)

func GenSessionKey(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

func UsernameCheck(username string) error {
	if len(username) >= 1 && len(username) <= 30 && structs.UsernameIsCorrect.MatchString(username) {
		return nil
	}
	return errors.New("Incorrect username")
}

func EmailCheck(email string) error {
	if structs.EmailIsCorrect.MatchString(email) {
		return nil
	}
	return errors.New("Incorrect email")
}

func PasswordCheck(password string) error {
	if len(password) >= 8 && len(password) <= 30 && structs.PasswordIsCorrect.MatchString(password) {
		if structs.PasswordHasAperCaseChar.MatchString(password) {
			if structs.PasswordHasDownCaseChar.MatchString(password) {
				if structs.PasswordHasSpecChar.MatchString(password) {
					return nil
				}
				return errors.New("Password has not special symbol")
			}
			return errors.New("Password has not symbol in down case")
		}
		return errors.New("Password has not symbol in upper case")
	}
	return errors.New("Incorrect password")
}

func NameCheck(name string) error {
	if len(name) >= 1 && len(name) <= 30 && structs.NameIsCorrect.MatchString(name) {
		return nil
	}
	return errors.New("Incorrct name")
}

func SurnameCheck(surname string) error {
	if len(surname) >= 1 && len(surname) <= 30 && structs.SurnameIsCorrect.MatchString(surname) {
		return nil
	}
	return errors.New("Incorrect surname")
}

func AgeCheck(age string) error {
	if structs.AgeIsCorrect.MatchString(age) {
		return nil
	}
	return errors.New("Incorrect age")
}

func StatusCheck(status string) error {
	if len(status) >= 1 && len(status) <= 200 && structs.StatusIsCorrect.MatchString(status) {
		return nil
	}
	return errors.New("Incorrect status")
}

func EditProfileDataCheck(newProfileUser *structs.EditUserProfile) error {
	if newProfileUser.Email != "" {
		if err := EmailCheck(newProfileUser.Email); err != nil {
			//SetResponseError(encoder, "incorrect Email", err)
			return err
		}
	}
	if newProfileUser.Username != "" {
		if err := UsernameCheck(newProfileUser.Username); err != nil {
			//SetResponseError(encoder, "incorrect Username", err)
			return err
		}
	}
	if newProfileUser.Password != "" {
		if err := PasswordCheck(newProfileUser.Password); err != nil {
			//SetResponseError(encoder, "incorrect Password", err)
			return err
		}
	}
	if newProfileUser.Name != "" {
		if err := NameCheck(newProfileUser.Name); err != nil {
			//SetResponseError(encoder, "incorrect Name", err)
			return err
		}
	}
	if newProfileUser.Surname != "" {
		if err := SurnameCheck(newProfileUser.Surname); err != nil {
			//SetResponseError(encoder, "incorrect Surname", err)
			return err
		}
	}
	if newProfileUser.Status != "" {
		if err := StatusCheck(newProfileUser.Status); err != nil {
			//SetResponseError(encoder, "incorrect Status", err)
			return err
		}
	}
	if newProfileUser.Age != "" {
		if err := AgeCheck(newProfileUser.Age); err != nil {
			//SetResponseError(encoder, "incorrect Status", err)
			return err
		}
	}
	return nil
}
