package delivery

import (
	"context"
	user_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/user-service/service_model"
	//"github.com/go-park-mail-ru/2019_2_Solar/cmd/authorization-service/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
)

type HandlersStruct struct {
	PUsecase usecase.UseInterface
}

type UserService struct {
	Usecase 	usecase.UseInterface
	Host		string
}

func NewUserService(use usecase.UseInterface,
	port string) *UserService {
	return &UserService{
		Usecase: use,
		Host: port,
	}
}

func (us *UserService) GetUsers(ctx context.Context, in *user_service.Nothing) (*user_service.Users, error) {

	users, err := us.Usecase.GetAllUsers()
	if err != nil {
		return &user_service.Users{}, err
	}

	var usersMessage []*user_service.AnotherUser

	for _, element := range users {
		usersMessage = append(usersMessage, &user_service.AnotherUser{
			ID:                   element.ID,
			Username:             element.Username,
			Name:                 element.Name,
			Surname:              element.Surname,
			Email:                element.Email,
			Age:                  uint64(element.Age),
			Status:               element.Status,
			AvatarDir:            element.AvatarDir,
			IsActive:             element.IsActive,
			IsFollowee:           element.IsFollowee,
		})
	}

	
	return &user_service.Users{
		Users:                usersMessage,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

func (us *UserService) GetUserByUsername(ctx context.Context, in *user_service.Username) (*user_service.UserAndPins, error) {



	userProfile, err := us.Usecase.GetUserByUsername(in.Username)
	if err != nil {
		return &user_service.UserAndPins{}, nil
	}

	pins, err := us.Usecase.GetPinsByUsername(int(userProfile.ID))
	if err != nil {
		return &user_service.UserAndPins{}, nil
	}

	userMessage :=  user_service.AnotherUser{
		ID:                   userProfile.ID,
		Username:             userProfile.Username,
		Name:                 userProfile.Name,
		Surname:              userProfile.Surname,
		Email:                userProfile.Email,
		Age:                  uint64(userProfile.Age),
		Status:               userProfile.Status,
		AvatarDir:            userProfile.AvatarDir,
		IsActive:             userProfile.IsActive,
		IsFollowee:           userProfile.IsFollowee,
	}

	var pinsMessage []*user_service.PinDisplay

	for _, element := range pins {
		pinsMessage = append(pinsMessage, &user_service.PinDisplay{
			ID:                   element.ID,
			PinDir:               element.PinDir,
			Title:                element.Title,
		})
	}



	return &user_service.UserAndPins{
		User:                 &userMessage,
		Pins:                 pinsMessage,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

func (us *UserService) CreateSubscribe(ctx context.Context, in *user_service.UserIDAndFolloweeUsername) (*user_service.Nothing, error) {


	if err := us.Usecase.AddSubscribe(in.UserID, in.FolloweeUsername); err != nil {
		return &user_service.Nothing{}, err
	}

	return &user_service.Nothing{}, nil
}

func (us *UserService) DeleteSubscribe(ctx context.Context, in *user_service.UserIDAndFolloweeUsername) (*user_service.Nothing, error) {

	if err := us.Usecase.RemoveSubscribe(in.UserID, in.FolloweeUsername); err != nil {
		return &user_service.Nothing{}, err
	}

	return &user_service.Nothing{}, nil
}