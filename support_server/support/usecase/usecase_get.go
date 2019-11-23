package usecase

import "github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"

func (use *UseStruct) GetAdminByLogin(login string) (Admin models.Admin, Err error) {

	admin, err := use.PRepository.SelectAdminByLogin(login)
	if err != nil {
		return admin, err
	}


	return admin, nil
}

