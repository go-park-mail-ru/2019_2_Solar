package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
)

func (use *UseStruct) GetAdminByLogin(login string) (Admin models.Admin, Err error) {

	admin, err := use.PRepository.SelectAdminByLogin(login)
	if err != nil {
		return admin, err
	}


	return admin, nil
}

func (use *UseStruct) GetHubListActiveUsers() (Data map[*webSocket.Client]bool, Err error) {
	activeUsers := use.Hub.Clients

	return activeUsers, nil
}

