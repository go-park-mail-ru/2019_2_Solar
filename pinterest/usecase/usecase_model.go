package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"io"
	"net/http"
	"sync"
)

type UsecaseStruct struct {
	PRepository repository.RepositoryInterface
	Mu          *sync.Mutex
}

type UsecaseInterface interface {
	SetJsonData(data interface{}, infMsg string) models.OutJSON
	SetResponseError(encoder *json.Encoder, msg string, err error)

	ReadUserStructByEmail(email string) (models.User, error)
	ReadUserIdByEmail(email string) (string, error)
	GetAllUsers() ([]models.User, error)

	RegDataValidationCheck(newUser *models.UserReg) error
	RegUsernameEmailIsUnique(username, email string) error

	EditProfileDataValidationCheck(newProfileUser *models.EditUserProfile) error
	EditUsernameEmailIsUnique(newUsername, newEmail, username, email string, userId uint64) error

	SetUserAvatarDir(idUser, fileName string) (int, error)
	SetUser(newUser models.EditUserProfile, user models.User) (int, error)
	InsertNewUser(username, email, password string) (string, error)
	CreateNewUserSession(userId string) (http.Cookie, error)


	ExtractFormatFile(fileName string) (string, error)
	DeleteOldUserSession(sessionKey string) error
	CalculateMD5FromFile (fileByte io.Reader) (string, error)
	CreateDir(folder string) error
	CreatePictureFile(fileName string, fileByte io.Reader) (Err error)
}
