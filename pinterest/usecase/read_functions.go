package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC UsecaseStruct) GetUserIdByEmail(email string) (string, error) {
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

func (USC UsecaseStruct) GetUserByEmail(email string) (models.User, error) {
	var userSlice []models.User
	var params []interface{}
	params = append(params, email)
	var err error
	userSlice, err = USC.PRepository.SelectFullUser(consts.SELECTUserByEmail, params)
	if err != nil {
		return models.User{}, err
	}
	if len(userSlice) != 1 {
		return models.User{}, errors.New("several users")
	}
	return userSlice[0], nil
}

func (USC *UsecaseStruct) GetAllUsers() ([]models.User, error) {
	var err error

	users, err := USC.PRepository.SelectFullUser(consts.SELECTAllUsers, nil)
	if err != nil {
		return users, err
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
	return pin[0], nil
}

func (USC *UsecaseStruct) GetBoard(boardID uint64) (models.Board, error) {
	var err error
	var params []interface{}
	params = append(params, boardID)

	board, err := USC.PRepository.SelectBoard(consts.SELECTBoardById, params)
	if err != nil {
		return board, err
	}
	return board, nil
}

func (USC *UsecaseStruct) GetPins(boardID uint64) ([]models.Pin, error) {
	var err error
	var params []interface{}
	params = append(params, boardID)

	pins, err := USC.PRepository.SelectPin(consts.SELECTPinsByBoardId, params)
	if err != nil {
		return []models.Pin{}, err
	}
	return pins, nil
}
func (USC *UsecaseStruct) GetNewPins() ([]models.Pin, error) {
	var err error
	var params []interface{}
	params = append(params, consts.NumberOfPinsOnPage)
	pins, err := USC.PRepository.SelectPinsSortByDate(consts.SELECTNewPinsByNumber, params)
	if err != nil {
		return []models.Pin{}, err
	}
	return pins, nil
}
