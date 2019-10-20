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
	DBWriteData(executeQuery string, params []interface{}) (string, error)
	DBReadDataUser(executeQuery string, params []interface{}) ([]models.User, error)
	DBReadDataUserCookies(executeQuery string, params []interface{}) ([]models.UserCookie, error)
	DBReadDataString(executeQuery string, params []interface{}) ([]string, error)

	DELETE_SESSION(executeQuery string, params []interface{}) (error)
}
