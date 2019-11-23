package repository

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
)

func (rep *ReposStruct) SelectUserByID(userID uint64) (User models.User, Err error) {
	user := models.DBUser{}
	err := rep.DataBase.QueryRow(consts.SELECTUserByID, userID).Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Age,
		&user.Status, &user.AvatarDir, &user.IsActive, &user.Salt, &user.CreatedTime)
	if err != nil {
		return models.User{}, err
	}

	goodUser := models.User{
		ID:          user.ID,
		Username:    user.Username,
		Name:        user.Name.String,
		Surname:     user.Surname.String,
		Password:    user.Password,
		Email:       user.Email,
		Age:         uint(user.Age.Int32),
		Status:      user.Status.String,
		AvatarDir:   user.AvatarDir.String,
		IsActive:    user.IsActive,
		Salt:        user.Salt,
		CreatedTime: user.CreatedTime,
	}

	return goodUser, nil
}

func (rep *ReposStruct) SelectAdminByLogin(login string) (Admin models.Admin, Err error) {
	admin := models.Admin{}
	err := rep.DataBase.QueryRow(consts.SELECTAdminByLogin, login).Scan(&admin.ID, &admin.Login, &admin.Password)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

