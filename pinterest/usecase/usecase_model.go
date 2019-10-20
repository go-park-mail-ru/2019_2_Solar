package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
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

	RegDataValidationCheck(newUser *models.UserReg) error
	RegEmailIsUnique(email string) (bool, error)
	RegUsernameIsUnique(username string) (bool, error)

	EditProfileDataValidationCheck(newProfileUser *models.EditUserProfile) error
	EditUsernameEmailIsUnique(newUsername, newEmail, username, email string, userId uint64) (bool, error)

	UpdateUser(user models.User, userId uint64) (string, error)
	InsertNewUser(username, email, password string) (string, error)
	CreateNewUserSession(userId string) (http.Cookie, error)
}
