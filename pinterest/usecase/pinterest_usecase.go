package usecase

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"log"
	"net/http"
	"sync"
	"time"
)

type PinterestUsecase struct {
	Users    []models.User
	Sessions []models.UserSession
	Mu       *sync.Mutex
}

func NewPinterestUsecase(users []models.User, sessions []models.UserSession,
	mu *sync.Mutex) pinterest.Usecase {
	return &PinterestUsecase{
		Users:    users,
		Sessions: sessions,
		Mu:       mu,
	}
}

func (p *PinterestUsecase) CreateNewUser(newUserReg *models.UserReg) models.User {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	var id uint64 = 0
	if len(p.Users) > 0 {
		id = p.Users[len(p.Users)-1].ID + 1
	}

	newUser := models.User{
		ID:       id,
		Name:     "",
		Password: newUserReg.Password,
		Email:    newUserReg.Email,
		Username: newUserReg.Username,
	}

	p.Users = append(p.Users, newUser)

	return newUser
}

func (p *PinterestUsecase) CreateNewUserSession(user models.User) ([]http.Cookie, models.UserSession, error) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	cookies := []http.Cookie{}

	var sessionValue uint64 = 0
	if len(p.Sessions) > 0 {
		sessionValue = p.Sessions[len(p.Sessions)-1].ID + 1
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

	p.Sessions = append(p.Sessions, newUserSession)

	return cookies, newUserSession, nil
}

func (p *PinterestUsecase) SaveUserPictureDir(userID uint64, fileName string) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	p.Users[userID].AvatarDir = fileName
}

func (p *PinterestUsecase) DeleteOldUserSession(value string) error {
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

func (p *PinterestUsecase) SearchCookie(r *http.Request) (*http.Cookie, error) {
	key, err := r.Cookie("session_key")
	return key, err
}

func (p *PinterestUsecase) RegEmailIsUnique(email string) bool {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, user := range p.Users {
		if user.Email == email {
			return false
		}
	}
	return true
}

func (p *PinterestUsecase) RegUsernameIsUnique(username string) bool {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, user := range p.Users {
		if user.Username == username {
			return false
		}
	}
	return true
}

func (p *PinterestUsecase) EditEmailIsUnique(email string, idUser uint64) bool {
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

func (p *PinterestUsecase) EditUsernameIsUnique(username string, idUser uint64) bool {
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

func (p *PinterestUsecase) SearchUserByEmail(newUserLogin *models.UserLogin) interface{} {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, user := range p.Users {
		if user.Email == newUserLogin.Email {
			return user
		}
	}
	return ""
}

func (p *PinterestUsecase) GetUserIndexByID(id uint64) int {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for index, user := range p.Users {
		if user.ID == id {
			return index
		}
	}
	return -1
}

func (p *PinterestUsecase) GetAllUsers() []models.User {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	return p.Users
}

func (p *PinterestUsecase) GetUserByID(id uint64) models.User {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	return p.Users[id]
}

func (p *PinterestUsecase) SetJsonData(data interface{}, infMsg string) models.OutJSON {
	p.Mu.Lock()
	defer p.Mu.Unlock()

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

func (p *PinterestUsecase) SearchIdUserByCookie(r *http.Request) (uint64, error) {
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

func (p *PinterestUsecase) SaveNewProfileUser(userID uint64, newUser *models.EditUserProfile) {
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

func (p *PinterestUsecase) ExtractFormatFile(FileName string) (string, error) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for i := 0; i < len(FileName); i++ {
		if string(FileName[i]) == "." {
			return FileName[i:], nil
		}
	}
	return "", errors.New("invalid file name")
}

func (p *PinterestUsecase) SetResponseError(encoder *json.Encoder, msg string, err error) {
	log.Printf("%s: %s", msg, err)
	data := p.SetJsonData(nil, msg)
	encoder.Encode(data)
}

func (p *PinterestUsecase) GenSessionKey(length int) string {
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

func (p *PinterestUsecase) RegDataCheck(newUser *models.UserReg) error {
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

func (p *PinterestUsecase) EditProfileDataCheck(newProfileUser *models.EditUserProfile) error {
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
