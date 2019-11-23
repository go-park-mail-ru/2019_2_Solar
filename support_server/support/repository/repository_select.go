package repository

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
)

func (rep *ReposStruct) SelectAdminByLogin(login string) (Admin models.Admin, Err error) {
	admin := models.Admin{}
	err := rep.DataBase.QueryRow(consts.SELECTAdminByLogin, login).Scan(&admin.ID, &admin.Login, &admin.Password)
	if err != nil {
		return admin, err
	}

	return admin, nil
}
