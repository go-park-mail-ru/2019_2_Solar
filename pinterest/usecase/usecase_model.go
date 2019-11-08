package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"sync"
)


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

	CheckBoardData(newBoard models.NewBoard) error
	CheckPinData(newPin models.NewPin) error

	SetUserAvatarDir(idUser uint64, fileName string) (int, error)
	SetUser(newUser models.EditUserProfile, user models.User) (int, error)

	AddNewUser(username, email, password string) (uint64, error)
	AddNewUserSession(userID uint64) (http.Cookie, error)

	AddBoard(newBoard models.Board) (uint64, error)
	GetBoard(boardID uint64) (models.Board, error)
	GetMyBoards(UserID uint64) ([]models.Board, error)

	AddPin(newPin models.Pin) (uint64, error)
	GetPin(pinID uint64) (models.FullPin, error)
	//GetPins(boardID uint64) ([]models.Pin, error)
	GetPinsDisplay(boardID uint64) ([]models.PinDisplay, error)
	GetNewPins() ([]models.PinDisplay, error)
	GetMyPins(userID uint64) ([]models.PinDisplay, error)
	GetSubscribePins(userID uint64) ([]models.PinDisplay, error)

	AddComment(pinID, userID uint64, newComment models.NewComment) error
	GetComments(pinID uint64) ([]models.CommentDisplay, error)

	AddNotice(newNotice models.Notice) (uint64, error)
	GetMyNotices(userID uint64) ([]models.Notice, error)

	AddSubscribe(userID uint64, followeeName string) error
	RemoveSubscribe(userID uint64, followeeName string) error

	ExtractFormatFile(fileName string) (string, error)
	RemoveOldUserSession(sessionKey string) error
	CalculateMD5FromFile(fileByte io.Reader) (string, error)
	AddDir(folder string) error
	AddPictureFile(fileName string, fileByte io.Reader) (Err error)

	ReturnHub() *webSocket.HubStruct
	
	SearchPinsByTag(tag string) ([]models.PinDisplay, error)

	CreateClient(conn *websocket.Conn, userId uint64)


	GetMySubscribeByUsername(userId uint64, username string) (bool, error)

	AddTags(description string, pinID uint64) error



}
