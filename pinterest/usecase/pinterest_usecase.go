package usecase

import (
)


import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"log"
	"net/http"
)

func SaveUserPictureDir(userID uint64, fileName string) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	p.Users[userID].AvatarDir = fileName
}

func DeleteOldUserSession(value string) error {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for i, session := range p.Sessions {
		if session.Value == value {
			p.Sessions = append(p.Sessions[:i], p.Sessions[i+1:]...)
			return nil
		}
	}
	return errors.New("session has not found")
}

func SearchCookie(r *http.Request) (*http.Cookie, error) {
	key, err := r.Cookie("session_key")
	return key, err
}

func EditEmailIsUnique(email string, idUser uint64) bool {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, user := range p.Users {
		if user.Email == email {
			if user.ID != idUser {
				return false
			}
		}
	}
	return true
}

func EditUsernameIsUnique(username string, idUser uint64) bool {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, user := range p.Users {
		if user.Username == username {
			if user.ID != idUser {
				return false
			}
		}
	}
	return true
}

func SearchUserByEmail(newUserLogin *models.UserLogin) interface{} {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, user := range p.Users {
		if user.Email == newUserLogin.Email {
			return user
		}
	}
	return ""
}

func GetUserIndexByID(id uint64) int {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for index, user := range p.Users {
		if user.ID == id {
			return index
		}
	}
	return -1
}

func GetAllUsers() []models.User {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	return p.Users
}

func GetUserByID(id uint64) models.User {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	return p.Users[id]
}

func SetJsonData(data interface{}, infMsg string) models.OutJSON {
	user, ok := data.(models.User)
	if ok {
		outJSON := models.OutJSON{
			BodyJSON: models.DataJSON{
				UserJSON: user,
				InfoJSON: infMsg,
			},
		}
		return outJSON
	}
	if users, ok := data.([]models.User); ok {

		outJSON := models.OutJSON{
			BodyJSON: models.DataJSON{
				UsersJSON: users,
				InfoJSON:  infMsg,
			},
		}
		return outJSON
	}
	outJSON := models.OutJSON{
		BodyJSON: models.DataJSON{
			InfoJSON: infMsg,
		},
	}
	return outJSON
}

func SearchIdUserByCookie(r *http.Request) (uint64, error) {
	p.Mu.Lock()
	defer p.Mu.Unlock()


	sessionKey, err := p.SearchCookie(r)
	if err == http.ErrNoCookie {
		return 0, errors.New("cookies not found")
	}

	for _, oneSession := range p.Sessions {
		if oneSession.Value == sessionKey.Value {
			return oneSession.UserID, nil
		}
	}
	return 0, errors.New("idUser not found")
}

func SaveNewProfileUser(userID uint64, newUser *models.EditUserProfile) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	user := p.Users[userID]

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

	p.Users[userID] = user
}

func ExtractFormatFile(FileName string) (string, error) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

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

func GenSessionKey(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = functions.SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & consts.LetterIdxMask); idx < len(consts.LetterBytes) {
			result[i] = consts.LetterBytes[idx]
			i++
		}
	}

	return string(result)
}

func RegDataCheck(newUser *models.UserReg) error {
	if err := functions.EmailCheck(newUser.Email); err != nil {
		return err
	}
	if err := functions.UsernameCheck(newUser.Username); err != nil {
		return err
	}
	if err := functions.PasswordCheck(newUser.Password); err != nil {
		return err
	}
	return nil
}

func EditProfileDataCheck(newProfileUser *models.EditUserProfile) error {
	if newProfileUser.Email != "" {
		if err := functions.EmailCheck(newProfileUser.Email); err != nil {
			//SetResponseError(encoder, "incorrect Email", err)
			return err
		}
	}
	if newProfileUser.Username != "" {
		if err := functions.UsernameCheck(newProfileUser.Username); err != nil {
			//SetResponseError(encoder, "incorrect Username", err)
			return err
		}
	}
	if newProfileUser.Password != "" {
		if err := functions.PasswordCheck(newProfileUser.Password); err != nil {
			//SetResponseError(encoder, "incorrect Password", err)
			return err
		}
	}
	if newProfileUser.Name != "" {
		if err := functions.NameCheck(newProfileUser.Name); err != nil {
			//SetResponseError(encoder, "incorrect Name", err)
			return err
		}
	}
	if newProfileUser.Surname != "" {
		if err := functions.SurnameCheck(newProfileUser.Surname); err != nil {
			//SetResponseError(encoder, "incorrect Surname", err)
			return err
		}
	}
	if newProfileUser.Status != "" {
		if err := functions.StatusCheck(newProfileUser.Status); err != nil {
			//SetResponseError(encoder, "incorrect Status", err)
			return err
		}
	}
	if newProfileUser.Age != "" {
		if err := functions.AgeCheck(newProfileUser.Age); err != nil {
			//SetResponseError(encoder, "incorrect Status", err)
			return err
		}
	}
	return nil
}
