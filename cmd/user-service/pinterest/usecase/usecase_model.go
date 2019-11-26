package usecase

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"io"
	"net/http"
	"sync"
)


type AuthorizationService struct {
	PRepository repository.ReposInterface
	Sanitizer   sanitizer.SanitInterface
	Mu          *sync.Mutex
	Host		string
}

func NewAuthorizationService(mu *sync.Mutex, rep repository.ReposInterface,
	san *sanitizer.SanitStruct, port string) *AuthorizationService {
	return &AuthorizationService{
		PRepository: rep,
		Sanitizer: san,
		Mu: mu,
		Host: port,
	}
}

func (auth *AuthorizationService) CheckSession(ctx context.Context, cookie *services.Cookie) (*services.UserSession, error) {

	return &services.UserSession{
		ID:                   1,
		UserID:               1,
		Value:                "123",
		Exp:                  "23:00",

	}, nil
}

func (auth *AuthorizationService) RegUser(ctx context.Context, userReg *services.UserReg) (*services.Cookie, error) {

	return &services.Cookie{
		Key:                  "key",
		Value:                "123",
		Exp:                  "23:00",
	}, nil
}

func (auth *AuthorizationService) LoginUser(ctx context.Context, userLogin *services.UserLogin) (*services.Cookie, error) {

	return &services.Cookie{
		Key:                  "key",
		Value:                "123",
		Exp:                  "23:00",
	}, nil
}

func (auth *AuthorizationService) LogoutUser(ctx context.Context, cookie *services.Cookie) (*services.Nothing, error) {

	return &services.Nothing{
		Dummy:                false,
	}, nil
}

type UseStruct struct {
	PRepository repository.ReposInterface
	Sanitizer   sanitizer.SanitInterface
	Hub         webSocket.HubStruct
	Mu          *sync.Mutex
}

type UseInterface interface {
	SetJSONData(data interface{}, token string, infMsg string) models.OutJSON
	SetResponseError(encoder *json.Encoder, msg string, err error) error

	GetUserByUsername(username string) (models.AnotherUser, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserIDByEmail(email string) (uint64, error)

	GetAllUsers() ([]models.AnotherUser, error)
	ComparePassword(password, salt, loginPassword string) error

	CheckRegDataValidation(newUser *models.UserReg) error
	CheckRegUsernameEmailIsUnique(username, email string) error

	CheckProfileData(newProfileUser *models.EditUserProfile) error
	CheckUsernameEmailIsUnique(newUsername, newEmail, username, email string, userID uint64) error

	SetUserAvatarDir(idUser uint64, fileName string) (int, error)
	SetUser(newUser models.EditUserProfile, user models.User) (int, error)

	AddNewUser(username, email, password string) (uint64, error)
	AddNewUserSession(userID uint64) (http.Cookie, error)


	ExtractFormatFile(fileName string) (string, error)
	RemoveOldUserSession(sessionKey string) error
	CalculateMD5FromFile(fileByte io.Reader) (string, error)
	AddDir(folder string) error
	AddPictureFile(fileName string, fileByte io.Reader) (Err error)

}
