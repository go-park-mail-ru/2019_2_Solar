package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
)

func (USC *UsecaseStruct) RegDataValidationCheck(newUser *models.UserReg) error {
	if err := EmailCheck(newUser.Email); err != nil {
		return err
	}
	if err := UsernameCheck(newUser.Username); err != nil {
		return err
	}
	if err := PasswordCheck(newUser.Password); err != nil {
		return err
	}
	return nil
}

func UsernameCheck(username string) error {
	if len(username) >= 1 && len(username) <= 30 && validation.UsernameIsCorrect.MatchString(username) {
		return nil
	}
	return errors.New("Incorrect username")
}

func EmailCheck(email string) error {
	if validation.EmailIsCorrect.MatchString(email) {
		return nil
	}
	return errors.New("Incorrect email")
}

//Упростить
func PasswordCheck(password string) error {
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

func NameCheck(name string) error {
	if len(name) >= 1 && len(name) <= 30 && validation.NameIsCorrect.MatchString(name) {
		return nil
	}
	return errors.New("Incorrct name")
}

func SurnameCheck(surname string) error {
	if len(surname) >= 1 && len(surname) <= 30 && validation.SurnameIsCorrect.MatchString(surname) {
		return nil
	}
	return errors.New("Incorrect surname")
}

func AgeCheck(age string) error {
	if validation.AgeIsCorrect.MatchString(age) {
		return nil
	}
	return errors.New("Incorrect age")
}

func StatusCheck(status string) error {
	if len(status) >= 1 && len(status) <= 200 && validation.StatusIsCorrect.MatchString(status) {
		return nil
	}
	return errors.New("incorrect status")
}

func (USC *UsecaseStruct) RegEmailIsUnique(email string) (bool, error) {
	var str []string
	var params []interface{}
	params = append(params, email)
	str, err := USC.PRepository.DBReadDataString(consts.ReadUserIdByEmailSQLQuery, params)
	if err != nil || len(str) > 1 {
		return false, err
	}
	return true, nil
}

func (USC *UsecaseStruct) RegUsernameIsUnique(username string) (bool, error) {
	var str []string
	var params []interface{}
	params = append(params, username)
	str, err := USC.PRepository.DBReadDataString(consts.SelectUserUsernameByUsername, params)
	if err != nil || len(str) > 1 {
		return false, err
	}
	return true, nil
}

func (USC *UsecaseStruct) EditProfileDataValidationCheck(newProfileUser *models.EditUserProfile) error {
	if newProfileUser.Email != "" {
		if err := EmailCheck(newProfileUser.Email); err != nil {
			return err
		}
	}
	if newProfileUser.Username != "" {
		if err := UsernameCheck(newProfileUser.Username); err != nil {
			return err
		}
	}
	if newProfileUser.Password != "" {
		if err := PasswordCheck(newProfileUser.Password); err != nil {
			return err
		}
	}
	if newProfileUser.Name != "" {
		if err := NameCheck(newProfileUser.Name); err != nil {
			return err
		}
	}
	if newProfileUser.Surname != "" {
		if err := SurnameCheck(newProfileUser.Surname); err != nil {
			return err
		}
	}
	if newProfileUser.Status != "" {
		if err := StatusCheck(newProfileUser.Status); err != nil {
			return err
		}
	}
	if newProfileUser.Age != "" {
		if err := AgeCheck(newProfileUser.Age); err != nil {
			return err
		}
	}
	return nil
}

func (USC *UsecaseStruct) EditUsernameEmailIsUnique(newUsername, newEmail, username, email string, userId uint64) (bool, error) {
	if newUsername == username && newEmail == email {
		return true, nil
	}
	var userSlice []models.User
	var params []interface{}
	params = append(params, newUsername)
	params = append(params, newEmail)
	userSlice, err := USC.PRepository.DBReadDataUser(consts.ReadUserIdByUsernameEmailSQLQuery, params)
	if err != nil {
		return false, err
	}
	for _, user := range userSlice {
		if user.ID == userId {
			continue
		}
		if user.Username == newUsername {
			return false, errors.New("username is not unique")
		}
		if user.Email == newEmail {
			return false, errors.New("email is not unique")
		}
	}
	return true, nil
}
