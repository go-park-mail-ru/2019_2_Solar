package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC UsecaseStruct) ReadUserIdByEmail(email string) (string, error) {
	var str repository.StringSlice
	var params []interface{}
	params = append(params, email)
	var err error = errors.New("several users")
	err = USC.PRepository.UniversalRead(consts.ReadUserIdByEmailSQLQuery, &str, params)
	if err != nil || len(str) != 1 {
		return "", err
	}
	return str[0], nil
}

func (USC UsecaseStruct) ReadUserStructByEmail(email string) (models.User, error) {
	var userSlice repository.UsersSlice
	var params []interface{}
	params = append(params, email)
	var err error = errors.New("several users")
	err = USC.PRepository.UniversalRead(consts.ReadUserByEmailSQLQuery, &userSlice, params)
	if err != nil || len(userSlice) != 1 {
		return models.User{}, err
	}
	return userSlice[0], nil
}
