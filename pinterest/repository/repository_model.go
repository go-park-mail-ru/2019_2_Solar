package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

type ReposStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type ReposInterface interface {
	Update(executeQuery string, params []interface{}) (int, error)
	Insert(executeQuery string, params []interface{}) (string, error)

	SelectFullUser(executeQuery string, params []interface{}) ([]models.User, error)
	SelectUserCookies(executeQuery string, params []interface{}) ([]models.UserCookie, error)
	SelectOneCol(executeQuery string, params []interface{}) ([]string, error)
	SelectIDUsernameEmailUser(executeQuery string, params []interface{}) ([]models.UserUnique, error)
	DeleteSession(executeQuery string, params []interface{}) error
	DeleteSubscribe(executeQuery string, params []interface{}) error

	SelectCategory(executeQuery string, params []interface{}) ([]string, error)

	SelectBoard(executeQuery string, params []interface{}) (models.Board, error)

	SelectPin(executeQuery string, params []interface{}) ([]models.Pin, error)
	SelectIDDirPins(executeQuery string, params []interface{}) (Pins []models.PinForMainPage, Err error)
	SelectComments(executeQuery string, params []interface{}) (Comments []models.CommentForSend, Err error)
	SelectPinsByTag(executeQuery string, params []interface{}) (Pins []models.PinForSearchResult, Err error)

	SelectSessions(executeQuery string, params []interface{}) (Sessions []models.UserSession, Err error)

	SelectBoards(executeQuery string, params []interface{}) (Boards []models.Board, Err error)
}
