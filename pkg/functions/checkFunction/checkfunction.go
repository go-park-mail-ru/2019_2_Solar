package checkFunction

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
)

func RegDataCheck(newUser *models.UserReg) error {
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

func RegEmailIsUnique(email string) (bool, error) {
	DBWorker := repository.DataBaseWorker{}
	DBWorker.NewDataBaseWorker()
	var str repository.StringSlice
	err := DBWorker.UniversalRead(consts.FindEmailSQLQuery+"'"+email+"'", &str)
	if err != nil || len(str) > 1 {
		return false, err
	}
	return true, nil
}

func RegUsernameIsUnique(username string) (bool, error) {
	DBWorker := repository.DataBaseWorker{}
	DBWorker.NewDataBaseWorker()
	var str repository.StringSlice
	err := DBWorker.UniversalRead(consts.FindUsernameSQLQuery+"'"+username+"'", &str)
	if err != nil || len(str) > 1 {
		return false, err
	}
	return true, nil
}
