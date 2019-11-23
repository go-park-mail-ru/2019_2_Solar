package useCaseMiddleware

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	repositoryMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/repository/middleware"
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
