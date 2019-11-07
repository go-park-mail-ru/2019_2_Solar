package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (USC *UseStruct) GetUserIDByEmail(email string) (string, error) {
	var str []string
	var params []interface{}
	params = append(params, email)
	var err error
	str, err = USC.PRepository.SelectOneCol(consts.SELECTUserIDByEmail, params)
	if err != nil {
		return "", err
	}
	if len(str) != 1 {
		return "", errors.New("several users or no one user")
	}
	return str[0], nil
}

func (USC *UseStruct) GetUserByUsername(username string) (models.AnotherUser, error) {
	var userSlice []models.User
	var params []interface{}
	params = append(params, username)
	var err error
	userSlice, err = USC.PRepository.SelectFullUser(consts.SELECTUserByUsername, params)
	if err != nil {
		return models.AnotherUser{}, err
	}
	if len(userSlice) != 1 {
		return models.AnotherUser{}, errors.New("several users or no one user")
	}
	USC.Sanitizer.SanitUser(&userSlice[0])
	anotherUser := models.AnotherUser{
		ID:        userSlice[0].ID,
		Username:  userSlice[0].Username,
		Name:      userSlice[0].Name,
		Surname:   userSlice[0].Surname,
		Password:  userSlice[0].Password,
		Email:     userSlice[0].Email,
		Age:       userSlice[0].Age,
		Status:    userSlice[0].Status,
		AvatarDir: userSlice[0].AvatarDir,
		IsActive:  userSlice[0].IsActive,
	}
	return anotherUser, nil
}

func (USC *UseStruct) GetUserByEmail(email string) (models.User, error) {
	userSlice, err := USC.PRepository.SelectUsersByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	if len(userSlice) == 0 {
		return models.User{}, errors.New("user not found")
	}
	if len(userSlice) > 1 {
		return models.User{}, errors.New("several same users")
	}
	USC.Sanitizer.SanitUser(&userSlice[0])
	return userSlice[0], nil
}

func (USC *UseStruct) GetAllUsers() ([]models.AnotherUser, error) {
	users, err := USC.PRepository.SelectAllUsers()
	if err != nil {
		return []models.AnotherUser{}, err
	}
	anotherUsers := []models.AnotherUser{}
	for _, user := range users {
		USC.Sanitizer.SanitUser(&user)
		anotherUser := models.AnotherUser{
			ID:        user.ID,
			Username:  user.Username,
			Name:      user.Name,
			Surname:   user.Surname,
			Password:  user.Password,
			Email:     user.Email,
			Age:       user.Age,
			Status:    user.Status,
			AvatarDir: user.AvatarDir,
			IsActive:  user.IsActive,
		}
		anotherUsers = append(anotherUsers, anotherUser)
	}
	return anotherUsers, nil
}

func (USC *UseStruct) GetPin(pinID uint64) (models.Pin, error) {
	pin, err := USC.PRepository.SelectPinsById(pinID)
	if err != nil {
		return pin[0], err
	}
	USC.Sanitizer.SanitPin(&pin[0])
	return pin[0], nil
}

func (USC *UseStruct) GetBoard(boardID uint64) (models.Board, error) {
	board, err := USC.PRepository.SelectBoardsByID(boardID)
	if err != nil {
		return models.Board{}, err
	}
	if len(board) == 0 {
		return models.Board{}, errors.New("board not found")
	}
	if len(board) > 1 {
		return models.Board{}, errors.New("several same boards")
	}
	USC.Sanitizer.SanitBoard(&board[0])
	return board[0], nil
}

func (USC *UseStruct) GetMyBoards(UserID uint64) ([]models.Board, error) {
	boards, err := USC.PRepository.SelectBoardsByOwnerId(UserID)
	if err != nil {
		return boards, err
	}
	for _, board := range boards {
		USC.Sanitizer.SanitBoard(&board)
	}
	return boards, nil
}

func (USC *UseStruct) GetPinsDisplay(boardID uint64) ([]models.PinDisplay, error) {
	pins, err := USC.PRepository.SelectPinsDisplayByBoardId(boardID)
	if err != nil {
		return []models.PinDisplay{}, err
	}
	for _, pin := range pins {
		USC.Sanitizer.SanitPinDisplay(&pin)
	}
	return pins, nil
}

func (USC *UseStruct) GetNewPins() ([]models.PinDisplay, error) {
	pins, err := USC.PRepository.SelectNewPinsDisplayByNumber(0, consts.NumberOfPinsOnPage)
	if err != nil {
		return []models.PinDisplay{}, err
	}
	return pins, nil
}

func (USC *UseStruct) GetMyPins(userID uint64) ([]models.PinForMainPage, error) {
	var err error
	var params []interface{}
	params = append(params, consts.NumberOfPinsOnPage, userID)
	pins, err := USC.PRepository.SelectIDDirPins(consts.SELECTMyPinsByNumber, params)
	if err != nil {
		return []models.PinForMainPage{}, err
	}
	return pins, nil
}

func (USC *UseStruct) GetSubscribePins(userID uint64) ([]models.PinForMainPage, error) {
	var err error
	var params []interface{}
	params = append(params, consts.NumberOfPinsOnPage, userID)
	pins, err := USC.PRepository.SelectIDDirPins(consts.SELECTSubscribePinsByNumber, params)
	if err != nil {
		return []models.PinForMainPage{}, err
	}
	return pins, nil
}

func (USC *UseStruct) GetComments(pinID uint64) ([]models.CommentDisplay, error) {
	var err error
	var params []interface{}
	params = append(params, pinID)
	comments, err := USC.PRepository.SelectCommentsByPinId(pinID)
	if err != nil {
		return []models.CommentDisplay{}, err
	}
	for _, comment := range comments {
		USC.Sanitizer.SanitComment(&comment)
	}
	return comments, nil
}
