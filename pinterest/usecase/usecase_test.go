package usecase

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/mocks"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	//"time"
)

func TestInsertNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		user := &models.UserReg{
			Username: "Vitaly",
			Email:    "something@mail.ru",
			Password: "123QWErty!",
		}
		//salt := "500"
		var params []interface{}
		params = append(params, user.Username, user.Email, user.Password)

		repo.EXPECT().InsertUser(user.Username, user.Email, gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(1), nil)

		newUserId, err := us.AddNewUser(user.Username, user.Email, user.Password)

		assert.NotNil(t, newUserId)
		assert.Equal(t, newUserId, uint64(1))
		assert.Nil(t, err)
	})
}

func TestCreateNewUserSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	user := models.User{
		ID:       1,
		Username: "Vitaly",
		Email:    "something@mail.ru",
		Password: "123QWErty!",
	}
	//cookieValue := "FF"
	t.Run("success", func(t *testing.T) {
		repo.EXPECT().InsertSession(user.ID, gomock.Any(), gomock.Any()).Return(uint64(1), nil)

		cookie, err := us.AddNewUserSession(user.ID)

		assert.NoError(t, err)
		assert.NotNil(t, cookie)
	})
}

func TestDeleteOldUserSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	session := models.UserSession{
		ID:     0,
		UserID: 0,
		UserCookie: models.UserCookie{
			Value:      "QWERTY",
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	t.Run("success", func(t *testing.T) {
		sessionKey := session.Value
		var params []interface{}
		params = append(params, sessionKey)

		repo.EXPECT().DeleteSessionByKey(sessionKey).Return(nil)

		err := us.RemoveOldUserSession(sessionKey)

		assert.NoError(t, err)
	})

	t.Run("incorrect session value", func(t *testing.T) {
		sessionKey := "blalba"
		var params []interface{}
		params = append(params, sessionKey)

		repo.EXPECT().DeleteSessionByKey(sessionKey).Return(errors.New("incorrect key"))

		err := us.RemoveOldUserSession(sessionKey)

		assert.Error(t, err)
	})
}

func TestEditUsernameEmailIsUnique(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	user := models.User{
		ID:       1,
		Username: "Vitaly",
		Email:    "something@mail.ru",
		Password: "123QWErty!",
	}

	t.Run("success", func(t *testing.T) {
		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectIDUsernameEmailUser(newUsername, newEmail).Return(nil, nil)

		err := us.CheckUsernameEmailIsUnique(newUsername, newEmail, user.Username, user.Email, user.ID)

		assert.NoError(t, err)
	})
	t.Run("not unique username", func(t *testing.T) {
		users := []models.UserUnique{
			{
				ID:       100,
				Email:    "gnkngvsltkbs",
				Username: "UniqueUsername",
			},
		}
		
		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectIDUsernameEmailUser(newUsername, newEmail).Return(users, nil)

		err := us.CheckUsernameEmailIsUnique(newUsername, newEmail, user.Username, user.Email, user.ID)

		assert.Error(t, err)
	})
	t.Run("not unique email", func(t *testing.T) {
		users := []models.UserUnique{
			{
				ID:       100,
				Email:    "UniqueEmail",
				Username: "Toooosername",
			},
		}

		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectIDUsernameEmailUser(newUsername, newEmail).Return(users, nil)

		err := us.CheckUsernameEmailIsUnique(newUsername, newEmail, user.Username, user.Email, user.ID)

		assert.Error(t, err)
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

func TestEditProfileDataValidationCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

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

		err := us.CheckProfileData(&newProfile)

		assert.NoError(t, err)
	})
	t.Run("long pass", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "12nreonr324RBIVEURNvrvIERFNRUBUI5RBGfiwfbwbufewgegw",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("no char", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "1234362865765353635634",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("no number", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "vinbsonvslktbnstrs",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("incorrect name", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "@#%#$%@",
			Surname:  "Filcost",
			Password: "vinbs23421352NBfgtrs",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("incorrect surname", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "?????",
			Password: "vinbs23421352NBfgtrs",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("incorrect username", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "ЛОПИУЦАТ",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "vinbs23421352NBfgtrs",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("incorrect email", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "vinbs23421352NBfgtrs",
			Email:    "@mail.su",
			Age:      "42",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("incorrect age", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "vinbs23421352NBfgtrs",
			Email:    "email@mail.su",
			Age:      "-999",
			Status:   "Ok",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
	t.Run("incorrect status", func(t *testing.T) {
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "vinbs23421352NBfgtrs",
			Email:    "email@mail.su",
			Age:      "42",
			Status: "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh",
		}

		err := us.CheckProfileData(&newProfile)

		assert.Error(t, err)
	})
}

//
//func TestGetAllUsers(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	repo := mocks.NewMockReposInterface(ctrl)
//
//	var mutex sync.Mutex
//	san := sanitizer.SanitStruct{}
//	san.NewSanitizer()
//	hub := webSocket.HubStruct{}
//	hub.NewHub()
//
//	us := UseStruct{}
//	us.NewUseCase(&mutex, repo, &san, hub)
//
//	expectedUsers := make([]models.User, 0)
//
//	user1 := models.User{
//		ID:       0,
//		Username: "Vitaly",
//		Email:    "something@mail.ru",
//		Password: "123QWErty!",
//	}
//	user2 := models.User{
//		ID:       1,
//		Username: "Vova",
//		Email:    "something2@mail.ru",
//		Password: "123QWErty!",
//	}
//	user3 := models.User{
//		ID:       2,
//		Username: "Nastya",
//		Email:    "something2@mail.ru",
//		Password: "123QWErty!",
//	}
//	user4 := models.User{
//		ID:       3,
//		Username: "Bogdan",
//		Email:    "something3@mail.ru",
//		Password: "123QWErty!",
//	}
//
//	expectedUsers = append(expectedUsers, user1, user2, user3, user4)
//
//	t.Run("success", func(t *testing.T) {
//		newUsername := "UniqueUsername"
//		newEmail := "UniqueEmail"
//		var params []interface{}
//		params = append(params, newUsername, newEmail)
//
//		repo.EXPECT().SelectAllUsers().Return(expectedUsers, nil)
//
//		users, err := us.GetAllUsers()
//
//		assert.NoError(t, err)
//		assert.NotNil(t, users)
//		assert.Equal(t, users[0].Email, expectedUsers[0].Email)
//		assert.Equal(t, users[1].Email, expectedUsers[1].Email)
//		assert.Equal(t, users[2].Email, expectedUsers[2].Email)
//	})
//}

func TestGenSessionKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		sessionKeyLenght := 20

		key, err := GenSessionKey(sessionKeyLenght)

		assert.NoError(t, err)
		assert.NotNil(t, key)
		assert.Equal(t, len(key), sessionKeyLenght)
	})
}

func TestGetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

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

		repo.EXPECT().SelectUsersByEmail(email).Return(expectedUsers, nil)

		user, err := us.GetUserByEmail(email)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, expectedUser)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

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

		actualProfile := models.User{
			ID:          0,
			Username:    "Alcost",
			Name:        "Alcost",
			Surname:     "Filcost",
			Password:    "123ewrEW#",
			Email:       "email@mail.su",
			Age:         uint(42),
			Status:      "Ok",
			IsActive:    false,
			Salt:        "",
			CreatedTime: time.Time{},
		}

		repo.EXPECT().UpdateUser(actualProfile).Return(0, nil)

		id, err := us.SetUser(newProfile, user)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, int(user.ID))
	})
}

func TestAddBoard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		newBoard := models.Board{
			OwnerID:     14,
			Title:       "SomeTitle",
			Description: "SomeDesc",
			Category:    "Cars",
			CreatedTime: time.Now(),
		}
		var params []interface{}
		params = append(params, newBoard.OwnerID, newBoard.Title, newBoard.Description,
			newBoard.Category, newBoard.CreatedTime)

		repo.EXPECT().InsertBoard(newBoard.OwnerID, newBoard.Title, newBoard.Description,
			newBoard.Category, gomock.Any()).Return(uint64(1), nil)

		id, err := us.AddBoard(newBoard)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, uint64(1))
	})
}

func TestAddPin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		newPin := models.Pin{
			OwnerID:     14,
			AuthorID:    14,
			BoardID:     1,
			PinDir:      "/die/",
			Title:       "SomeTitle",
			Description: "SomeDesc",
			CreatedTime: time.Now(),
		}
		var params []interface{}
		params = append(params, newPin.OwnerID, newPin.AuthorID, newPin.BoardID, newPin.Title, newPin.Description,
			newPin.PinDir, newPin.CreatedTime)

		repo.EXPECT().InsertPin(newPin).Return(uint64(1), nil)

		id, err := us.AddPin(newPin)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, uint64(1))
	})
}

func TestSetJSONData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success for user json", func(t *testing.T) {
		user := models.User{
			ID:       0,
			Username: "Vitaly",
			Email:    "something@mail.ru",
			Password: "123QWErty!",
		}
		json := us.SetJSONData(user, "", "OK")
		assert.NotNil(t, json)
	})
	t.Run("success for anotherUser json", func(t *testing.T) {
		user := models.AnotherUser{
			ID:       0,
			Username: "Vitaly",
			Password: "123QWErty!",
		}
		json := us.SetJSONData(user, "", "OK")
		assert.NotNil(t, json)
	})
	t.Run("success for users json", func(t *testing.T) {
		users := []models.User{
			{
				ID:       0,
				Username: "Vitaly",
				Email:    "something@mail.ru",
				Password: "123QWErty!",
			},
			{
				ID:       1,
				Username: "Anton",
				Email:    "Logvinov@mail.ru",
				Password: "123QWErty!",
			},
		}
		json := us.SetJSONData(users, "", "OK")
		assert.NotNil(t, json)
	})
	t.Run("success for AnotherUsers json", func(t *testing.T) {
		users := []models.AnotherUser{
			{
				ID:       0,
				Username: "Vitaly",
				Password: "123QWErty!",
			},
			{
				ID:       1,
				Username: "Anton",
				Password: "123QWErty!",
			},
		}
		json := us.SetJSONData(users, "", "OK")
		assert.NotNil(t, json)
	})
	t.Run("success for info json", func(t *testing.T) {
		json := us.SetJSONData(nil, "", "OK")
		assert.NotNil(t, json)
	})
}

func TestAddNotice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		newNotice := models.Notice{
			UserID:      1,
			ReceiverID:  1,
			Message:     "Hello",
			CreatedTime: time.Time{},
		}
		var params []interface{}
		params = append(params, newNotice.UserID, newNotice.ReceiverID, newNotice.Message,
			newNotice.CreatedTime)

		repo.EXPECT().InsertNotice(newNotice).Return(uint64(1), nil)

		lastID, err := us.AddNotice(newNotice)

		assert.NoError(t, err)
		assert.NotNil(t, lastID)
		assert.Equal(t, lastID, uint64(1))
	})
}

func TestAddConment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		newComment := models.NewComment{
			Text: "Fuu",
		}
		var params []interface{}
		params = append(params, "1", newComment.Text, uint64(1), time.Now())

		repo.EXPECT().InsertComment(uint64(1), newComment.Text, uint64(1), gomock.Any()).Return(uint64(1), nil)

		err := us.AddComment(1, 1, newComment)

		assert.NoError(t, err)
	})
}

func TestExtractFormatFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		fileName := "file.txt"

		format, err := us.ExtractFormatFile(fileName)

		assert.NoError(t, err)
		assert.NotNil(t, format)
		assert.Equal(t, ".txt", format)
	})
	t.Run("failed", func(t *testing.T) {
		fileName := "file,txt"

		_, err := us.ExtractFormatFile(fileName)

		assert.Error(t, err)
	})
}

func TestSearchPinsByTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		tagName := "car"

		expectedPin := []models.PinDisplay{
			{
				ID:     1,
				PinDir: "/dir/",
				Title:  "TheBest",
			},
		}

		repo.EXPECT().SelectPinsByTag(tagName).Return(expectedPin, nil)

		pinst, err := us.SearchPinsByTag(tagName)

		assert.NoError(t, err)
		assert.NotNil(t, pinst)
		assert.Equal(t, pinst[0], expectedPin[0])
	})
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		selectedUsers := []models.User{
			{
				ID:       0,
				Username: "Vitaly",
				Email:    "something@mail.ru",
				Password: "123QWErty!",
			},
			{
				ID:       1,
				Username: "Anton",
				Email:    "Logvinov@mail.ru",
				Password: "123QWErty!",
			},
		}

		expectedUsers := []models.AnotherUser{
			{
				ID:       0,
				Username: "Vitaly",
				Email:    "something@mail.ru",
				Password: "123QWErty!",
			},
			{
				ID:       1,
				Username: "Anton",
				Email:    "Logvinov@mail.ru",
				Password: "123QWErty!",
			},
		}

		repo.EXPECT().SelectAllUsers().Return(selectedUsers, nil)

		users, err := us.GetAllUsers()

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, users, expectedUsers)
	})
}

func TestGetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		username := "Vitaly"
		selectedUsers := []models.User{
			{
				ID:       0,
				Username: "Vitaly",
				Email:    "something@mail.ru",
				Password: "123QWErty!",
			},
		}

		expectedUsers := []models.AnotherUser{
			{
				ID:       0,
				Username: "Vitaly",
				Email:    "something@mail.ru",
				Password: "123QWErty!",
			},
		}
		var params []interface{}
		params = append(params, username)

		repo.EXPECT().SelectUsersByUsername(username).Return(selectedUsers, nil)

		user, err := us.GetUserByUsername(username)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUsers[0], user)
	})
}

func TestGetUserIDByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		users := []models.User{
			{
				ID:       uint64(1),
				Email:    "something@mail.ru",
				Username: "userq",
				Password: "123",
			},
		}
		email := "something@mail.ru"

		repo.EXPECT().SelectUsersByEmail(email).Return(users, nil)

		id, err := us.GetUserIDByEmail(email)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, uint64(1))
	})
}

func TestGetPin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		pinID := uint64(1)
		selectedPins := []models.FullPin{
			{
				BoardID:     1,
				PinDir:      "/die/",
				Title:       "SomeTitle",
				Description: "SomeDesc",
				CreatedTime: time.Now(),
			},
		}

		repo.EXPECT().SelectPinsById(pinID).Return(selectedPins, nil)

		pin, err := us.GetPin(pinID)

		assert.NoError(t, err)
		assert.NotNil(t, pin)
		assert.Equal(t, pin, selectedPins[0])
	})
}

func TestGetBoard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		boardID := uint64(1)
		selectedBoard := []models.Board{
			{
				ID:          1,
				OwnerID:     14,
				Title:       "SomeTitle",
				Description: "SomeDesc",
				Category:    "cars",
				CreatedTime: time.Now(),
			},
		}

		repo.EXPECT().SelectBoardsByID(boardID).Return(selectedBoard, nil)

		board, err := us.GetBoard(boardID)

		assert.NoError(t, err)
		assert.NotNil(t, board)
		assert.Equal(t, board, selectedBoard[0])
	})
}

func TestGetPins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		boardID := uint64(1)
		expPins := []models.PinDisplay{
			{
				PinDir: "/die/",
				Title:  "SomeTitle",
			},
			{

				PinDir: "/die/1",
				Title:  "SomeTitle",
			},
		}
		var params []interface{}
		params = append(params, boardID)

		repo.EXPECT().SelectPinsDisplayByBoardId(boardID).Return(expPins, nil)

		pins, err := us.GetPinsDisplay(boardID)

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestGetNewPins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		expPins := []models.PinDisplay{
			{
				ID:     1,
				PinDir: "/dir/1",
			},
			{
				ID:     2,
				PinDir: "/dir/2",
			},
		}
		var params []interface{}
		params = append(params, consts.NumberOfPinsOnPage)

		repo.EXPECT().SelectNewPinsDisplayByNumber(0, 10).Return(expPins, nil)

		pins, err := us.GetNewPins()

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestGetMyPins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		userID := uint64(1)
		expPins := []models.PinDisplay{
			{
				ID:     1,
				PinDir: "/dir/1",
			},
			{
				ID:     2,
				PinDir: "/dir/2",
			},
		}
		var params []interface{}
		params = append(params, consts.NumberOfPinsOnPage, userID)

		repo.EXPECT().SelectMyPinsDisplayByNumber(userID, 10).Return(expPins, nil)

		pins, err := us.GetMyPins(userID)

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestGetSubscribePins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		userID := uint64(1)
		expPins := []models.PinDisplay{
			{
				ID:     1,
				PinDir: "/dir/1",
			},
			{
				ID:     2,
				PinDir: "/dir/2",
			},
		}
		var params []interface{}
		params = append(params, consts.NumberOfPinsOnPage, userID)

		repo.EXPECT().SelectSubscribePinsDisplayByNumber(userID, 0, 10).Return(expPins, nil)

		pins, err := us.GetSubscribePins(userID)

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestGetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		pinID := uint64(1)
		expComments := []models.CommentDisplay{
			{
				Text:        "blablalb",
				CreatedTime: time.Now(),
				Author:      "Iam",
			},
			{
				Text:        "blablalb2",
				CreatedTime: time.Now(),
				Author:      "Iam2",
			},
		}
		var params []interface{}
		params = append(params, pinID)

		repo.EXPECT().SelectCommentsByPinId(pinID).Return(expComments, nil)

		comments, err := us.GetComments(pinID)

		assert.NoError(t, err)
		assert.NotNil(t, comments)
		assert.Equal(t, comments, expComments)
	})
}

func TestCheckRegDataValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		newUser := models.UserReg{
			Email:    "Vitalian42@mail.ru",
			Password: "12341ttbceybEMFwef",
			Username: "Vinil",
		}

		err := us.CheckRegDataValidation(&newUser)

		assert.NoError(t, err)
	})
	t.Run("bad email", func(t *testing.T) {
		newUser := models.UserReg{
			Email:    ".ru",
			Password: "12341ttbceybEMFwef",
			Username: "Vinil",
		}

		err := us.CheckRegDataValidation(&newUser)

		assert.Error(t, err)
	})
	t.Run("bad password", func(t *testing.T) {
		newUser := models.UserReg{
			Email:    "Vitalian42@mail.ru",
			Password: "1234",
			Username: "Vinil",
		}

		err := us.CheckRegDataValidation(&newUser)

		assert.Error(t, err)
	})
	t.Run("bad username", func(t *testing.T) {
		newUser := models.UserReg{
			Email:    "Vitalian42@mail.ru",
			Password: "12341ttbceybEMFwef",
			Username: "!",
		}

		err := us.CheckRegDataValidation(&newUser)

		assert.Error(t, err)
	})
}

func TestCheckBoardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	t.Run("success", func(t *testing.T) {
		newBoard := models.NewBoard{
			Title:       "CoolTitle",
			Description: "Desc",
			Category:    "cars",
		}

		repo.EXPECT().SelectCategoryByName(newBoard.Category).Return([]string{"cars"}, nil)

		err := us.CheckBoardData(newBoard)

		assert.NoError(t, err)
	})
	t.Run("bad title", func(t *testing.T) {
		newBoard := models.NewBoard{
			Title:       "",
			Description: "Desc",
			Category:    "cars",
		}

		//repo.EXPECT().SelectCategoryByName(newBoard.Category).Return([]string{"cars"}, nil)

		err := us.CheckBoardData(newBoard)

		assert.Error(t, err)
	})
	t.Run("bad description", func(t *testing.T) {
		newBoard := models.NewBoard{
			Title: "Cooltitle",
			Description: "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
				"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh",
			Category: "cars",
		}

		//repo.EXPECT().SelectCategoryByName(newBoard.Category).Return([]string{"cars"}, nil)

		err := us.CheckBoardData(newBoard)

		assert.Error(t, err)
	})
	t.Run("bad category", func(t *testing.T) {
		newBoard := models.NewBoard{
			Title:       "Cooltitle",
			Description: "desc",
			Category:    "ca",
		}

		repo.EXPECT().SelectCategoryByName(newBoard.Category).Return([]string{}, nil)

		err := us.CheckBoardData(newBoard)

		assert.Error(t, err)
	})
	t.Run("success", func(t *testing.T) {
		loginPass := "2nfrjkvnderfeNFKJb"
		salt := "500"
		password := HashPassword(loginPass, salt)

		err := us.ComparePassword(string(password), salt, loginPass)

		assert.NoError(t, err)
	})
	t.Run("bad compare", func(t *testing.T) {
		loginPass := "2nfrjkvnderfeNFKJb"
		salt := "500"
		password := HashPassword(loginPass, salt)

		err := us.ComparePassword(string(password)+"2", salt, loginPass)

		assert.Error(t, err)
	})
}


func TestCheckRegUsernameEmailIsUnique(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	us.NewUseCase(&mutex, repo, &san, hub)

	//user := models.User{
	//	ID:       1,
	//	Username: "Vitaly",
	//	Email:    "something@mail.ru",
	//	Password: "123QWErty!",
	//}

	t.Run("success", func(t *testing.T) {
		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectIDUsernameEmailUser(newUsername, newEmail).Return(nil, nil)

		err := us.CheckRegUsernameEmailIsUnique(newUsername, newEmail)

		assert.NoError(t, err)
	})
	t.Run("not unique username", func(t *testing.T) {
		users := []models.UserUnique{
			{
				ID:       100,
				Email:    "gnkngvsltkbs",
				Username: "UniqueUsername",
			},
		}

		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectIDUsernameEmailUser(newUsername, newEmail).Return(users, nil)

		err := us.CheckRegUsernameEmailIsUnique(newUsername, newEmail)

		assert.Error(t, err)
	})
	t.Run("not unique email", func(t *testing.T) {
		users := []models.UserUnique{
			{
				ID:       100,
				Email:    "UniqueEmail",
				Username: "Toooosername",
			},
		}

		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectIDUsernameEmailUser(newUsername, newEmail).Return(users, nil)

		err := us.CheckRegUsernameEmailIsUnique(newUsername, newEmail)

		assert.Error(t, err)
	})
}