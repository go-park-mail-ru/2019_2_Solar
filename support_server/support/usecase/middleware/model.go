package useCaseMiddleware

import (
	repositoryMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support/repository/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

type MUseCaseStruct struct {
	MRepository repositoryMiddleware.MRepositoryInterface
}

type MUseCaseInterface interface {
	NewUseCase(rep repositoryMiddleware.MRepositoryInterface)
	GetUserByCookieValue(cookieValue string) (models.User, error)
	//GetCookieByCookieValue(cookieValue string) (models.UserCookie, error)
	GetSessionsByCookieValue(cookieValue string) (models.UserSession, error)
}
