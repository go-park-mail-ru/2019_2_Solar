package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/mocks"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
	"time"

	//"time"
)

func TestPinterestUsecase_InsertNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	//ListUsers := make([]models.User, 0)
	//ListSessions := make([]models.UserSession, 0)


	t.Run("success", func(t *testing.T) {
		user := &models.UserReg{
			Username: "Vitaly",
			Email: "something@mail.ru",
			Password: "123QWErty!",

		}
		var params []interface{}
		params = append(params, user.Username, user.Email, user.Password)

		repo.EXPECT().Insert(consts.INSERTRegistration, params).Return("1", nil)

		newUserId, err := us.AddNewUser(user.Username, user.Email, user.Password)

		assert.NotNil(t, newUserId)
		assert.Equal(t, newUserId, "1")
		assert.Nil(t, err)
	})
}
func TestPinterestUsecase_CreateNewUserSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	user := models.User{
		ID: 1,
		Username: "Vitaly",
		Email: "something@mail.ru",
		Password: "123QWErty!",

	}
	t.Run("success", func(t *testing.T) {
		repo.EXPECT().Insert(consts.INSERTSession, gomock.Any()).Return("1", nil)
		cookie, err := us.AddNewUserSession(strconv.Itoa(int(user.ID)))

		assert.NoError(t, err)
		assert.NotNil(t, cookie)
	})
}

func TestPinterestUsecase_DeleteOldUserSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	session := models.UserSession{
		ID: 0,
		UserID: 0,
		UserCookie: models.UserCookie{
			Value: "QWERTY",
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	t.Run("success", func(t *testing.T) {
		sessionKey := session.Value
		var params []interface{}
		params = append(params, sessionKey)

		repo.EXPECT().DeleteSession(consts.DELETESessionByKey, params).Return(nil)

		err := us.RemoveOldUserSession(sessionKey)

		assert.NoError(t, err)
	})

	t.Run("incorrect session value", func(t *testing.T) {
		sessionKey := "blalba"
		var params []interface{}
		params = append(params, sessionKey)

		repo.EXPECT().DeleteSession(consts.DELETESessionByKey, params).Return(errors.New("incorrect key"))

		err := us.RemoveOldUserSession(sessionKey)

		assert.Error(t, err)
	})
}

func TestPinterestUsecase_EditUsernameEmailIsUnique(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	user := models.User{
		ID: 1,
		Username: "Vitaly",
		Email: "something@mail.ru",
		Password: "123QWErty!",

	}

	t.Run("success", func (t *testing.T) {
		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectFullUser(consts.SELECTUserIdUsernameEmailByUsernameOrEmail, params).Return(nil, nil)

		err := us.CheckUsernameEmailIsUnique(newUsername, newEmail, user.Username, user.Email, user.ID)

		assert.NoError(t, err)
	})
}
/*
func TestPinterestUsecase_EditUsernameIsUnique(t *testing.T) {
	user := models.User{
		ID: 0,
		Username: "Vitaly",
		Email: "something@mail.ru",
		Password: "123QWErty!",
	}

	session := models.UserSession{
		ID: 0,
		UserID: 0,
		UserCookie: models.UserCookie{
			Value: "QWERTY",
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	mockListUsers := make([]models.User, 0)
	mockListUsers = append(mockListUsers, user)

	mockListSession := make([]models.UserSession, 0)
	mockListSession = append(mockListSession, session)

	t.Run("success", func (t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		isUnique := uc.EditUsernameIsUnique("Vova", user.ID)

		assert.True(t, isUnique)
	})
}

func TestPinterestUsecase_RegEmailIsUnique(t *testing.T) {
	user := models.User{
		ID: 0,
		Username: "Vitaly",
		Email: "something@mail.ru",
		Password: "123QWErty!",
	}

	session := models.UserSession{
		ID: 0,
		UserID: 0,
		UserCookie: models.UserCookie{
			Value: "QWERTY",
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	mockListUsers := make([]models.User, 0)
	mockListUsers = append(mockListUsers, user)

	mockListSession := make([]models.UserSession, 0)
	mockListSession = append(mockListSession, session)

	t.Run("success", func (t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		isUnique := uc.RegEmailIsUnique("another@mail.ru")

		assert.True(t, isUnique)
	})
}

func TestPinterestUsecase_RegUsernameIsUnique(t *testing.T) {
	user := models.User{
		ID: 0,
		Username: "Vitaly",
		Email: "something@mail.ru",
		Password: "123QWErty!",
	}

	session := models.UserSession{
		ID: 0,
		UserID: 0,
		UserCookie: models.UserCookie{
			Value: "QWERTY",
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	mockListUsers := make([]models.User, 0)
	mockListUsers = append(mockListUsers, user)

	mockListSession := make([]models.UserSession, 0)
	mockListSession = append(mockListSession, session)

	t.Run("success", func (t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		isUnique := uc.RegUsernameIsUnique("Vova")

		assert.True(t, isUnique)
	})
}
*/

func TestPinterestUsecase_EditProfileDataValidationCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	t.Run("success", func (t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "123ewrEW#",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.NoError(t, err)
	})
}

func TestPinterestUsecase_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	expectedUsers := make([]models.User, 0)

	user1 := models.User{
		ID:       0,
		Username: "Vitaly",
		Email:    "something@mail.ru",
		Password: "123QWErty!",
	}
	user2 := models.User{
		ID:       1,
		Username: "Vova",
		Email:    "something2@mail.ru",
		Password: "123QWErty!",
	}
	user3 := models.User{
		ID:       2,
		Username: "Nastya",
		Email:    "something2@mail.ru",
		Password: "123QWErty!",
	}
	user4 := models.User{
		ID:       3,
		Username: "Bogdan",
		Email:    "something3@mail.ru",
		Password: "123QWErty!",
	}

	expectedUsers = append(expectedUsers, user1, user2, user3, user4)

	t.Run("success", func (t *testing.T) {
		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectFullUser(consts.SELECTAllUsers, nil).Return(expectedUsers, nil)

		users, err := us.GetAllUsers()

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, users, expectedUsers)
	})
}

func TestPinterestUsecase_GenSessionKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	t.Run("success", func (t *testing.T) {
		sessionKeyLenght := 20

		key, err := GenSessionKey(sessionKeyLenght)

		assert.NoError(t, err)
		assert.NotNil(t, key)
		assert.Equal(t, len(key), sessionKeyLenght)
	})
}

func TestPinterestUsecase_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	expectedUser := models.User{
		ID:       2,
		Username: "Nastya",
		Email:    "something2@mail.ru",
		Password: "123QWErty!",
	}
	expectedUsers := []models.User{expectedUser}

	t.Run("success", func(t *testing.T) {
		var email string = "something2@mail.ru"
		var params []interface{}
		params = append(params, email)

		repo.EXPECT().SelectFullUser(consts.SELECTUserByEmail, params).Return(expectedUsers, nil)

		user, err := us.GetUserByEmail(email)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, expectedUser)
	})
}

func TestPinterestUsecase_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepositoryInterface(ctrl)

	var mutex sync.Mutex

	us := UsecaseStruct{}
	us.NewUseCase(&mutex, repo)

	t.Run("success", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "123ewrEW#",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}
		user := models.User{
			ID:       0,
			Username: "Vitaly",
			Email:    "something@mail.ru",
			Password: "123QWErty!",
		}

		var params []interface{}
		params = append(params, newProfile.Username, newProfile.Email, newProfile.Password)

		repo.EXPECT().Insert(consts.INSERTRegistration, params).Return("0", nil)

		id, err := us.SetUser(newProfile, user)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, strconv.Itoa(int(user.ID)))
	})
}

/*
func TestPinterestUsecase_SearchUserByEmail(t *testing.T) {
	mockListUsers := make([]models.User, 0)
	mockListSession := make([]models.UserSession, 0)

	user1 := models.User{
		ID:       0,
		Username: "Vitaly",
		Email:    "something@mail.ru",
		Password: "123QWErty!",
	}

	mockListUsers = append(mockListUsers, user1)
	t.Run("success", func(t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})
		loginUser := models.UserLogin{
			Email:    "something@mail.ru",
			Password: "123QWErty!",
		}

		user := uc.SearchUserByEmail(&loginUser)

		assert.NotNil(t, user)

	})
}
*/