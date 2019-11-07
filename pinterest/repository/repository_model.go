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
	Update(executeQuery string, params []interface{}) (int, error)
	Insert(executeQuery string, params []interface{}) (string, error)

	SelectFullUser(executeQuery string, params []interface{}) ([]models.User, error)
	//SelectUserCookies(executeQuery string, params []interface{}) ([]models.UserCookie, error)
	SelectOneCol(executeQuery string, params []interface{}) ([]string, error)
	SelectIDUsernameEmailUser(executeQuery string, params []interface{}) ([]models.UserUnique, error)
	//DeleteSession(executeQuery string, params []interface{}) error
	DeleteSubscribe(executeQuery string, params []interface{}) error

	SelectCategory(executeQuery string, params []interface{}) ([]string, error)

	SelectBoard(executeQuery string, params []interface{}) (models.Board, error)

	SelectPin(executeQuery string, params []interface{}) ([]models.Pin, error)
	SelectIDDirPins(executeQuery string, params []interface{}) (Pins []models.PinForMainPage, Err error)
	SelectComments(executeQuery string, params []interface{}) (Comments []models.CommentForSend, Err error)
	SelectPinsByTag(executeQuery string, params []interface{}) (Pins []models.PinForSearchResult, Err error)

	SelectSessions(executeQuery string, params []interface{}) (Sessions []models.UserSession, Err error)

	SelectBoards(executeQuery string, params []interface{}) (Boards []models.Board, Err error)
	//-----------------------------------------------------
	SelectUsersByCookieValue(cookieValue string) (Users []models.User, Err error)
	SelectUsersByEmail(email string) (Users []models.User, Err error)
	SelectCookiesByCookieValue(cookieValue string) (Cookies []models.UserCookie, Err error)
	InsertUser(username, email, salt string, hashPassword []byte, createdTime time.Time) (uint64, error)
	InsertSession(userId uint64, cookieValue string, cookieExpires time.Time) (uint64, error)
	DeleteSessionByKey(cookieValue string) error
	SelectCategoryByName(categoryName string) (categories []string, Err error)
	//DeleteSubscribeByName()
	InsertBoard(ownerID uint64, title, description, category string, createdTime time.Time) (uint64, error)
	SelectBoardsByID(boardId uint64) (Boards []models.Board, Err error)
	SelectBoardsByOwnerId(ownerId uint64) (Boards []models.Board, Err error)
	SelectPinsDisplayByBoardId(boardID uint64) (Pins []models.PinDisplay, Err error)

}
