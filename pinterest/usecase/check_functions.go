package usecase

import (
	"bytes"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
)

func (USC *UseStruct) CheckRegDataValidation(newUser *models.UserReg) error {
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
	if len(username) >= 3 && len(username) <= 30 && validation.UsernameIsCorrect.MatchString(username) {
		return nil
	}
	return errors.New("Некорректное имя пользователя")
}

func EmailCheck(email string) error {
	if validation.EmailIsCorrect.MatchString(email) {
		return nil
	}
	return errors.New("Некорректный e-mail")
}

func PasswordCheck(password string) error {
	if len(password) < 6 {
		return errors.New("Слишком короткий пароль")
	}
	if len(password) > 30 {
		return errors.New("Слишком длинный пароль")
	}

	return nil
}

func NameCheck(name string) error {
	if len(name) >= 1 && len(name) <= 30 && validation.NameIsCorrect.MatchString(name) {
		return nil
	}
	return errors.New("Некорректное имя")
}

func SurnameCheck(surname string) error {
	if len(surname) >= 1 && len(surname) <= 30 && validation.SurnameIsCorrect.MatchString(surname) {
		return nil
	}
	return errors.New("Некорректная фамилия")
}

func AgeCheck(age string) error {
	if validation.AgeIsCorrect.MatchString(age) {
		return nil
	}
	return errors.New("Некорректный возраст")
}

func StatusCheck(status string) error {
	if len(status) >= 1 && len(status) <= 200 && validation.StatusIsCorrect.MatchString(status) {
		return nil
	}
	return errors.New("Некорректный статус")
}

func CheckBoardTitle(title string) error {
	if validation.BoardTitle.MatchString(title) {
		return nil
	}
	return errors.New("Некорректный заголовок")
}

func CheckBoardDescription(description string) error {
	if validation.BoardDescription.MatchString(description) {
		return nil
	}
	return errors.New("Некорректное описание")
}

func CheckPinTitle(title string) error {
	if validation.PinTitle.MatchString(title) {
		return nil
	}
	return errors.New("Некорректный заголовок")
}

func CheckPinDescription(description string) error {
	if validation.PinDescription.MatchString(description) {
		return nil
	}
	return errors.New("Некорректное описание")
}

func (USC *UseStruct) CheckBoardCategory(category string) error {
	var params []interface{}
	params = append(params, category)
	categories, err := USC.PRepository.SelectCategoryByName(category)
	if err != nil {
		return err
	}
	if len(categories) != 1 {
		return errors.New("Некорректное категория")
	}
	return nil
}

func (USC *UseStruct) CheckRegUsernameEmailIsUnique(username, email string) error {
	var userSlice []models.UserUnique
	var params []interface{}
	params = append(params, username, email)
	userSlice, err := USC.PRepository.SelectIDUsernameEmailUser(username, email)
	if err != nil {
		return err
	}
	for _, user := range userSlice {
		if user.Username == username {
			return errors.New("Никнейм уже занят")
		}
		if user.Email == email {
			return errors.New("Аккаунт с таким email уже существует")
		}
	}
	return nil
}

func (USC *UseStruct) CheckBoardData(newBoard models.NewBoard) error {
	if err := CheckBoardTitle(newBoard.Title); err != nil {
		return err
	}
	if err := CheckBoardDescription(newBoard.Description); err != nil {
		return err
	}
	if err := USC.CheckBoardCategory(newBoard.Category); err != nil {
		return err
	}
	return nil
}

func (USC *UseStruct) CheckPinData(newPin models.NewPin) error {
	if err := CheckPinTitle(newPin.Title); err != nil {
		return err
	}
	if err := CheckPinDescription(newPin.Description); err != nil {
		return err
	}
	return nil
}

func (USC *UseStruct) CheckProfileData(newProfileUser *models.EditUserProfile) error {
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

func (USC *UseStruct) CheckUsernameEmailIsUnique(newUsername, newEmail, username, email string, userID uint64) error {
	if newUsername == username && newEmail == email {
		return nil
	}
	var userSlice []models.UserUnique
	var params []interface{}
	params = append(params, newUsername, newEmail)
	userSlice, err := USC.PRepository.SelectIDUsernameEmailUser(newUsername, newEmail)
	if err != nil {
		return err
	}
	for _, user := range userSlice {
		if user.ID == userID {
			continue
		}
		if user.Username == newUsername {
			return errors.New("Никнейм уже занят")
		}
		if user.Email == newEmail {
			return errors.New("Аккаунт с таким email уже существует")
		}
	}
	return nil
}

func (USC *UseStruct) ComparePassword(password, salt, loginPassword string) error {
	if bytes.Equal([]byte(password), HashPassword(loginPassword, salt)) {
		return nil
	}
	return errors.New("Неверный пароль")
}
