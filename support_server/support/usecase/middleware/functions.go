package useCaseMiddleware

import (
	"errors"
	repositoryMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support/repository/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
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

func (MU *MUseCaseStruct) NewUseCase(rep repositoryMiddleware.MRepositoryInterface) {
	MU.MRepository = rep
}
