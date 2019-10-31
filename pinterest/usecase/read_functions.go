package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC *UsecaseStruct) GetUserIdByEmail(email string) (string, error) {
	var str []string
	var params []interface{}
	params = append(params, email)
	var err error
	str, err = USC.PRepository.SelectOneCol(consts.SELECTUserIdByEmail, params)
	if err != nil {
		return "", err
	}
	if len(str) != 1 {
		return "", errors.New("several users")
	}
	return str[0], nil
}

func (USC *UsecaseStruct) GetUserByUsername(username string) (models.User, error) {
	var userSlice []models.User
	var params []interface{}
	params = append(params, username)
	var err error
	userSlice, err = USC.PRepository.SelectFullUser(consts.SELECTUserByUsername, params)
	if err != nil {
		return models.User{}, err
	}
	if len(userSlice) != 1 {
		return models.User{}, errors.New("several users")
	}
	return USC.Sanitizer.SanitizeUser(userSlice[0]), nil
}

func (USC *UsecaseStruct) GetUserByEmail(email string) (models.User, error) {
	var userSlice []models.User
	var params []interface{}
	params = append(params, email)
	var err error
	userSlice, err = USC.PRepository.SelectFullUser(consts.SELECTUserByUsername, params)
	if err != nil {
		return models.User{}, err
	}
	if len(userSlice) != 1 {
		return models.User{}, errors.New("several users")
	}
	return USC.Sanitizer.SanitizeUser(userSlice[0]), nil
}

func (USC *UsecaseStruct) GetAllUsers() ([]models.User, error) {
	var err error

	users, err := USC.PRepository.SelectFullUser(consts.SELECTAllUsers, nil)
	if err != nil {
		return users, err
	}
	for _, user := range users {
		user = USC.Sanitizer.SanitizeUser(user)
	}
	return users, nil
}

func (USC *UsecaseStruct) GetPin(pinID uint64) (models.Pin, error) {
	var err error
	var params []interface{}
	params = append(params, pinID)

	pin, err := USC.PRepository.SelectPin(consts.SELECTPinById, params)
	if err != nil {
		return pin[0], err
	}
	return USC.Sanitizer.SanitizePin(pin[0]), nil
}

func (USC *UsecaseStruct) GetBoard(boardID uint64) (models.Board, error) {
	var err error
	var params []interface{}
	params = append(params, boardID)

	board, err := USC.PRepository.SelectBoard(consts.SELECTBoardById, params)
	if err != nil {
		return board, err
	}
	return USC.Sanitizer.SanitizeBoard(board), nil
}

func (USC *UsecaseStruct) GetPins(boardID uint64) ([]models.Pin, error) {
	var err error
	var params []interface{}
	params = append(params, boardID)

	pins, err := USC.PRepository.SelectPin(consts.SELECTPinsByBoardId, params)
	if err != nil {
		return []models.Pin{}, err
	}
	for _, pin := range pins {
		pin = USC.Sanitizer.SanitizePin(pin)
	}
	return pins, nil
}

func (USC *UsecaseStruct) GetNewPins() ([]models.PinForMainPage, error) {
	var err error
	var params []interface{}
	params = append(params, consts.NumberOfPinsOnPage)
	pins, err := USC.PRepository.SelectIdDirPins(consts.SELECTNewPinsByNumber, params)
	if err != nil {
		return []models.PinForMainPage{}, err
	}
	return pins, nil
}

func (USC *UsecaseStruct) GetMyPins(userId uint64) ([]models.PinForMainPage, error) {
	var err error
	var params []interface{}
	params = append(params, consts.NumberOfPinsOnPage, userId)
	pins, err := USC.PRepository.SelectIdDirPins(consts.SELECTMyPinsByNumber, params)
	if err != nil {
		return []models.PinForMainPage{}, err
	}
	return pins, nil
}

func (USC *UsecaseStruct) GetSubscribePins(userId uint64) ([]models.PinForMainPage, error) {
	var err error
	var params []interface{}
	params = append(params, consts.NumberOfPinsOnPage, userId)
	pins, err := USC.PRepository.SelectIdDirPins(consts.SELECTSubscribePinsByNumber, params)
	if err != nil {
		return []models.PinForMainPage{}, err
	}
	return pins, nil
}
