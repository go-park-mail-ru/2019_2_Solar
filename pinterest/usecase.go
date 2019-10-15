package pinterest

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"net/http"
)

type Usecase interface {
	CreateNewUser(newUser *models.UserReg) models.User
	CreateNewUserSession(newUserSession models.User) ([]http.Cookie, models.UserSession, error)
	SaveNewProfileUser(userID uint64, newUser *models.EditUserProfile)
	SaveUserPictureDir(userID uint64, fileName string)

	DeleteOldUserSession(string) error

	SearchCookie(*http.Request) (*http.Cookie, error)
	SearchUserByEmail(newUserLogin *models.UserLogin) interface{}
	SearchIdUserByCookie(r *http.Request) (uint64, error)

	GetUserIndexByID(id uint64) int
	GetUserByID(id uint64) models.User
	GetAllUsers() []models.User

	ExtractFormatFile(FileName string) (string, error)

	SetJsonData(data interface{}, infMsg string) models.OutJSON
	SetResponseError(encoder *json.Encoder, msg string, err error)

	RegEmailIsUnique(string) bool
	RegUsernameIsUnique(username string) bool
	EditEmailIsUnique(email string, idUser uint64) bool
	EditUsernameIsUnique(username string, idUser uint64) bool

	RegDataCheck(newUser *models.UserReg) error
	EditProfileDataCheck(newProfileUser *models.EditUserProfile) error
}
