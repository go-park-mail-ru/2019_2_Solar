package useCaseMiddleware

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	repositoryMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/repository/middleware"
)

func (MU *MUseCaseStruct) GetUserByCookieValue(cookieValue string) (models.User, error) {
	user, err := MU.MRepository.SelectUsersByCookieValue(cookieValue)
	if err != nil {
		return models.User{}, err
	}
	if len(user) == 0 {
		return models.User{}, errors.New("cookie not found")
	}
	if len(user) > 1 {
		return models.User{}, errors.New("several same cookies")
	}
	return user[0], nil
}

func (MU *MUseCaseStruct) GetAdminByCookieValue(cookieValue string) (models.Admin, error) {
	admin, err := MU.MRepository.SelectAdminByCookieValue(cookieValue)
	if err != nil {
		return models.Admin{}, err
	}
	if len(admin) == 0 {
		return models.Admin{}, errors.New("cookie not found")
	}
	if len(admin) > 1 {
		return models.Admin{}, errors.New("several same cookies")
	}
	return admin[0], nil
}

func (MU *MUseCaseStruct) GetSessionsByCookieValue(cookieValue string) (models.UserSession, error) {
	userSession, err := MU.MRepository.SelectSessionsByCookieValue(cookieValue)
	if err != nil {
		return models.UserSession{}, err
	}
	if len(userSession) == 0 {
		return models.UserSession{}, errors.New("cookie not found")
	}
	if len(userSession) > 1 {
		return models.UserSession{}, errors.New("several same cookies")
	}
	return userSession[0], nil
}

func (MU *MUseCaseStruct) GetAdminSessionsByCookieValue(cookieValue string) (models.AdminSession, error) {
	adminSession, err := MU.MRepository.SelectAdminSessionsByCookieValue(cookieValue)
	if err != nil {
		return models.AdminSession{}, err
	}
	if len(adminSession) == 0 {
		return models.AdminSession{}, errors.New("cookie not found")
	}
	if len(adminSession) > 1 {
		return models.AdminSession{}, errors.New("several same cookies")
	}
	return adminSession[0], nil
}

func (MU *MUseCaseStruct) NewUseCase(rep repositoryMiddleware.MRepositoryInterface) {
	MU.MRepository = rep
}
