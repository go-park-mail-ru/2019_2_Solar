package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

type RepositoryStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type RepositoryInterface interface {
	Update(executeQuery string, params []interface{}) (int, error)
	Insert(executeQuery string, params []interface{}) (string, error)
	SelectFullUser(executeQuery string, params []interface{}) ([]models.User, error)
	SelectUserCookies(executeQuery string, params []interface{}) ([]models.UserCookie, error)
	SelectOneCol(executeQuery string, params []interface{}) ([]string, error)
	SelectIdUsernameEmailUser(executeQuery string, params []interface{}) ([]models.UserUnique, error)
	DeleteSession(executeQuery string, params []interface{}) error

	SelectCategory(executeQuery string, params []interface{}) ([]string, error)
	InsertBoard(executeQuery string, params []interface{}) (string, error)
	InsertPin(executeQuery string, params []interface{}) (string, error)
}
