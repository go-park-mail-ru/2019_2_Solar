package pinterest

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"net/http"
)

type Usecase interface {
	CreateNewUser([]models.User, models.UserReg) models.User
	CreateNewUserSession([]models.UserSession, models.User) ([]http.Cookie, models.UserSession, error)
	SaveNewProfileUser(user *models.User, newUser *models.EditUserProfile)

	DeleteOldUserSession(*([]models.UserSession), string) error

	SearchCookie(*http.Request) (*http.Cookie, error)
	SearchUserByEmail(users []models.User, newUserLogin *models.UserLogin) interface{}
	SearchIdUserByCookie(r *http.Request, sessions []models.UserSession) (uint64, error)
	GetUserIndexByID(users []models.User, id uint64) int

	ExtractFormatFile(FileName string) (string, error)

	SetJsonData(data interface{}, infMsg string) models.OutJSON
	SetResponseError(encoder *json.Encoder, msg string, err error)

	RegEmailIsUnique([]models.User, string) bool
	RegUsernameIsUnique(users []models.User, username string) bool
	EditEmailIsUnique(users []models.User, email string, idUser uint64) bool
	EditUsernameIsUnique(users []models.User, username string, idUser uint64) bool

	RegDataCheck(newUser *models.UserReg) error
	EditProfileDataCheck(newProfileUser *models.EditUserProfile) error
}
