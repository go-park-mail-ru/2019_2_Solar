package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"time"
)

type ReposStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type ReposInterface interface {
	SelectUsersByCookieValue(cookieValue string) (Users []models.User, Err error)
	SelectUsersByEmail(email string) (Users []models.User, Err error)
	SelectCookiesByCookieValue(cookieValue string) (Cookies []models.UserCookie, Err error)
	InsertUser(username, email, salt string, hashPassword []byte, createdTime time.Time) (uint64, error)
	InsertSession(userId uint64, cookieValue string, cookieExpires time.Time) (uint64, error)
	DeleteSessionByKey(cookieValue string) error
	SelectCategoryByName(categoryName string) (categories []string, Err error)
	InsertBoard(ownerID uint64, title, description, category string, createdTime time.Time) (uint64, error)
	SelectBoardsByID(boardId uint64) (Boards []models.Board, Err error)
	SelectBoardsByOwnerId(ownerId uint64) (Boards []models.Board, Err error)
	SelectPinsDisplayByBoardId(boardID uint64) (Pins []models.PinDisplay, Err error)
	SelectAllUsers() (Users []models.User, Err error)
	InsertNotice(notice models.Notice) (uint64, error)
	InsertPin(pin models.Pin) (uint64, error)
	SelectPinsById(pinId uint64) (Pins []models.FullPin, Err error)
	SelectCommentsByPinId(pinId uint64) (Comments []models.CommentDisplay, Err error)
	SelectNewPinsDisplayByNumber(first, last int) (Pins []models.PinDisplay, Err error)
	SelectMyPinsDisplayByNumber(userId uint64, number int) (Pins []models.PinDisplay, Err error)
	SelectSubscribePinsDisplayByNumber(userId uint64, first, last int) (Pins []models.PinDisplay, Err error)
	InsertComment(pinID uint64, commentText string, userID uint64, createdTime time.Time) (uint64, error)
	SelectIDUsernameEmailUser(username, email string) (Users []models.UserUnique, Err error)
	UpdateUser(user models.User) (int, error)
	UpdateUserAvatar(fileName string, idUser uint64) (int, error)
	SelectPinsByTag(tag string) (Pins []models.PinDisplay, Err error)
	SelectUsersByUsername(username string) (Users []models.User, Err error)
	InsertSubscribe(userID uint64, followeeName string) (uint64, error)
	DeleteSubscribeByName(userID uint64, followeeName string) error
	InsertChatMessage(message models.NewChatMessage, createdTime time.Time) (uint64, error)
	SelectSessionsByCookieValue(cookieValue string) (Sessions []models.UserSession, Err error)

	SelectNoticesByUserID(userId uint64) (Notices []models.Notice, Err error)

	SelectMySubscribeByUsername(userId uint64, username string) (Subscribes []models.Subscribe, Err error)

	SelectAllTags() (Tag []string, Err error)

	InsertTag(Tag string) (Err error)

	InsertPinAndTag (PinID uint64, TagName string) (Err error)

}
