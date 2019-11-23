package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
)

func (use *UseStruct) GetUserByID(userID uint64) (User models.User, Err error) {
	user, err := use.PRepository.SelectUserByID(userID)
	if err != nil {
		return user, err
	}


	return user, nil
}

func (use *UseStruct) GetAdminByLogin(login string) (Admin models.Admin, Err error) {

	admin, err := use.PRepository.SelectAdminByLogin(login)
	if err != nil {
		return admin, err
	}


	return admin, nil
}

func (use *UseStruct) GetHubListActiveUsers() ([]uint64, error) {
	var clientSlice []uint64
	for client:=range use.Hub.Clients {
		clientSlice = append(clientSlice, client.UserId)
	}

	return clientSlice, nil
}

