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

	GetUserByEmail(email string) (models.User, error)
	GetUserIdByEmail(email string) (string, error)
	GetAllUsers() ([]models.User, error)

	CheckRegData(newUser *models.UserReg) error
	CheckRegUsernameEmailIsUnique(username, email string) error

	CheckProfileData(newProfileUser *models.EditUserProfile) error
	CheckUsernameEmailIsUnique(newUsername, newEmail, username, email string, userId uint64) error

	CheckBoardData(newBoard models.NewBoard) error
	CheckPinData(newPin models.NewPin) error

	SetUserAvatarDir(idUser, fileName string) (int, error)
	SetUser(newUser models.EditUserProfile, user models.User) (int, error)

	AddNewUser(username, email, password string) (string, error)
	AddNewUserSession(userId string) (http.Cookie, error)

	AddBoard(newBoard models.Board) (uint64, error)
	GetBoard(boardID uint64) (models.Board, error)

	AddPin(newPin models.Pin) (uint64, error)
	GetPin(pinID uint64) (models.Pin, error)
	GetPins(boardID uint64) ([]models.Pin, error)
	GetNewPins() ([]models.PinForMainPage, error)
	GetNewSubscribePins(userId uint64) ([]models.PinForMainPage, error)

	AddNotice(newNotice models.Notice) (uint64, error)

	ExtractFormatFile(fileName string) (string, error)
	RemoveOldUserSession(sessionKey string) error
	CalculateMD5FromFile (fileByte io.Reader) (string, error)
	AddDir(folder string) error
	AddPictureFile(fileName string, fileByte io.Reader) (Err error)
}
