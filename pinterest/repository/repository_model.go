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
	WriteData(executeQuery string, params []interface{}) (string, error)
	ReadUser(executeQuery string, params []interface{}) ([]models.User, error)
	ReadUserCookies(executeQuery string, params []interface{}) ([]models.UserCookie, error)
	ReadOneCol(executeQuery string, params []interface{}) ([]string, error)

	DeleteSession(executeQuery string, params []interface{}) (error)
}
