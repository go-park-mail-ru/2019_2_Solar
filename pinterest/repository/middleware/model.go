package repositoryMiddleware

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

type MRepositoryStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type MRepositoryInterface interface {
	DataBaseInit() error
	SelectUsersByCookieValue(cookieValue string) (Users []models.User, Err error)
	//SelectCookiesByCookieValue(cookieValue string) (Cookies []models.UserCookie, Err error)
	SelectSessionsByCookieValue(cookieValue string) (Sessions []models.UserSession, Err error)
}
