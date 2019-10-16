package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
)

func (USC *UsecaseStruct) RegDataValidationCheck(newUser *models.UserReg) error {
	if err := USC.EmailCheck(newUser.Email); err != nil {
		return err
	}
	if err := USC.UsernameCheck(newUser.Username); err != nil {
		return err
	}
	if err := USC.PasswordCheck(newUser.Password); err != nil {
		return err
	}
	return nil
}

func (USC *UsecaseStruct) UsernameCheck(username string) error {
	if len(username) >= 1 && len(username) <= 30 && validation.UsernameIsCorrect.MatchString(username) {
		return nil
	}
	return errors.New("Incorrect username")
}

func (USC *UsecaseStruct) EmailCheck(email string) error {
	if validation.EmailIsCorrect.MatchString(email) {
		return nil
	}
	return errors.New("Incorrect email")
}

//Упростить
func (USC *UsecaseStruct) PasswordCheck(password string) error {
	if len(password) >= 8 && len(password) <= 30 && validation.PasswordIsCorrect.MatchString(password) {
		if validation.PasswordHasAperCaseChar.MatchString(password) {
			if validation.PasswordHasDownCaseChar.MatchString(password) {
				if validation.PasswordHasSpecChar.MatchString(password) {
					return nil
				}
				return errors.New("Password has not special symbol")
			}
			return errors.New("Password has not symbol in down case")
		}
		return errors.New("Password has not symbol in upper case")
	}
	return errors.New("Incorrect password")
}

func (USC *UsecaseStruct) NameCheck(name string) error {
	if len(name) >= 1 && len(name) <= 30 && validation.NameIsCorrect.MatchString(name) {
		return nil
	}
	return errors.New("Incorrct name")
}

func (USC *UsecaseStruct) SurnameCheck(surname string) error {
	if len(surname) >= 1 && len(surname) <= 30 && validation.SurnameIsCorrect.MatchString(surname) {
		return nil
	}
	return errors.New("Incorrect surname")
}

func (USC *UsecaseStruct) AgeCheck(age string) error {
	if validation.AgeIsCorrect.MatchString(age) {
		return nil
	}
	return errors.New("Incorrect age")
}

func (USC *UsecaseStruct) StatusCheck(status string) error {
	if len(status) >= 1 && len(status) <= 200 && validation.StatusIsCorrect.MatchString(status) {
		return nil
	}
	return errors.New("Incorrect status")
}

func (USC *UsecaseStruct) RegEmailIsUnique(email string) (bool, error) {
	var str repository.StringSlice
	var params []interface{}
	params = append(params, email)
	err := USC.PRepository.UniversalRead(consts.FindEmailSQLQuery, &str, params)
	if err != nil || len(str) > 1 {
		return false, err
	}
	return true, nil
}

func (USC *UsecaseStruct) RegUsernameIsUnique(username string) (bool, error) {
	var str repository.StringSlice
	var params []interface{}
	params = append(params, username)
	err := USC.PRepository.UniversalRead(consts.FindUsernameSQLQuery, &str, params)
	if err != nil || len(str) > 1 {
		return false, err
	}
	return true, nil
}
