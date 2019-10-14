package usecase

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"log"
	"net/http"
	"time"
)

type PinterestUseCase struct {
}

func (p *PinterestUseCase) CreateNewUser(users []models.User, newUserReg models.UserReg) models.User {
	var id uint64 = 0
	if len(users) > 0 {
		id = users[len(users)-1].ID + 1
	}

	newUser := models.User{
		ID:       id,
		Name:     "",
		Password: newUserReg.Password,
		Email:    newUserReg.Email,
		Username: newUserReg.Username,
	}
	return newUser
}

func (p *PinterestUseCase) CreateNewUserSession(sessions []models.UserSession, user models.User) ([]http.Cookie, models.UserSession, error) {
	cookies := []http.Cookie{}

	var sessionValue uint64 = 0
	if len(sessions) > 0 {
		sessionValue = sessions[len(sessions)-1].ID + 1
	}

	sessionKey := p.GenSessionKey(12)

	cookieSessionKey := http.Cookie{
		Name:    "session_key",
		Value:   sessionKey,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	cookies = append(cookies, cookieSessionKey)

	newUserSession := models.UserSession{
		ID:     sessionValue,
		UserID: user.ID,
		UserCookie: models.UserCookie{
			Value:      sessionKey,
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	return cookies, newUserSession, nil
}

func (p *PinterestUseCase) DeleteOldUserSession(sessions *([]models.UserSession), value string) error {
	for i, session := range *sessions {
		if session.Value == value {
			*sessions = append((*sessions)[:i], (*sessions)[i+1:]...)
			return nil
		}
	}
	return errors.New("session has not found")
}

func (p *PinterestUseCase) SearchCookie(r *http.Request) (*http.Cookie, error) {
	key, err := r.Cookie("session_key")
	return key, err
}

func (p *PinterestUseCase) RegEmailIsUnique(users []models.User, email string) bool {
	for _, user := range users {
		if user.Email == email {
			return false
		}
	}
	return true
}

func (p *PinterestUseCase) RegUsernameIsUnique(users []models.User, username string) bool {
	for _, user := range users {
		if user.Username == username {
			return false
		}
	}
	return true
}

func (p *PinterestUseCase) EditEmailIsUnique(users []models.User, email string, idUser uint64) bool {
	for _, user := range users {
		if user.Email == email {
			if user.ID != idUser {
				return false
			}
		}
	}
	return true
}

func (p *PinterestUseCase) EditUsernameIsUnique(users []models.User, username string, idUser uint64) bool {
	for _, user := range users {
		if user.Username == username {
			if user.ID != idUser {
				return false
			}
		}
	}
	return true
}

func (p *PinterestUseCase) SearchUserByEmail(users []models.User, newUserLogin *models.UserLogin) interface{} {
	for _, user := range users {
		if user.Email == newUserLogin.Email {
			return user
		}
	}
	return ""
}

func (p *PinterestUseCase) GetUserIndexByID(users []models.User, id uint64) int {
	for index, user := range users {
		if user.ID == id {
			return index
		}
	}
	return -1
}

func (p *PinterestUseCase) SetJsonData(data interface{}, infMsg string) models.OutJSON {
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

func (p *PinterestUseCase) SearchIdUserByCookie(r *http.Request, sessions []models.UserSession) (uint64, error) {
	sessionKey, err := p.SearchCookie(r)
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

func (p *PinterestUseCase) SaveNewProfileUser(user *models.User, newUser *models.EditUserProfile) {
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

func (p *PinterestUseCase) ExtractFormatFile(FileName string) (string, error) {
	for i := 0; i < len(FileName); i++ {
		if string(FileName[i]) == "." {
			return FileName[i:], nil
		}
	}
	return "", errors.New("invalid file name")
}

func (p *PinterestUseCase) SetResponseError(encoder *json.Encoder, msg string, err error) {
	log.Printf("%s: %s", msg, err)
	data := p.SetJsonData(nil, msg)
	encoder.Encode(data)
}

func (p *PinterestUseCase) GenSessionKey(length int) string {
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

func (p *PinterestUseCase) RegDataCheck(newUser *models.UserReg) error {
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

func (p *PinterestUseCase) EditProfileDataCheck(newProfileUser *models.EditUserProfile) error {
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
