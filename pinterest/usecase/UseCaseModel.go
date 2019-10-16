package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"sync"
)

type UsecaseStruct struct {
	PRepository repository.RepositoryInterface
	Mu *sync.Mutex
}

type UsecaseInterface interface {
	RegDataValidationCheck(newUser *models.UserReg) error
	UsernameCheck(username string) error
	EmailCheck(email string) error
	PasswordCheck(password string) error
	NameCheck(name string) error
	SurnameCheck(surname string) error
	AgeCheck(age string) error
	StatusCheck(status string) error
	RegEmailIsUnique(email string) (bool, error)
	RegUsernameIsUnique(username string) (bool, error)
	InsertNewUser(username, email, password string)
}
