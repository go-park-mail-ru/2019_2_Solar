package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC UsecaseStruct) ReadUserIdByEmail(email string) (string, error) {
	var str []string
	var params []interface{}
	params = append(params, email)
	var err error = errors.New("several users")
	str, err = USC.PRepository.DBReadDataString(consts.ReadUserIdByEmailSQLQuery, params)
	if err != nil || len(str) != 1 {
		return "", err
	}
	return str[0], nil
}

func (USC UsecaseStruct) ReadUserStructByEmail(email string) (models.User, error) {
	var userSlice []models.User
	var params []interface{}
	params = append(params, email)
	var err error = errors.New("several users")
	userSlice, err = USC.PRepository.DBReadDataUser(consts.ReadUserByEmailSQLQuery, params)
	if err != nil || len(userSlice) != 1 {
		return models.User{}, err
	}
	return userSlice[0], nil
}
