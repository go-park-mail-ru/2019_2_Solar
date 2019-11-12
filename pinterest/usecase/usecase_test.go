package usecase

import (
	"errors"
	"strconv"
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

func TestPinterestUsecase_InsertNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		user := &models.UserReg{
			Username: "Vitaly",
			Email:    "something@mail.ru",
			Password: "123QWErty!",
		}
		var params []interface{}
		params = append(params, user.Username, user.Email, user.Password)

		repo.EXPECT().Insert(consts.INSERTRegistration, gomock.Any()).Return("1", nil)

		newUserId, err := us.AddNewUser(user.Username, user.Email, user.Password)

		assert.NotNil(t, newUserId)
		assert.Equal(t, newUserId, "1")
		assert.Nil(t, err)
	})
}

func TestPinterestUsecase_CreateNewUserSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	user := models.User{
		ID:       1,
		Username: "Vitaly",
		Email:    "something@mail.ru",
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

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

		repo.EXPECT().SelectIDUsernameEmailUser(consts.SELECTUserIDUsernameEmailByUsernameOrEmail, params).Return(nil, nil)

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

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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
}

func TestPinterestUsecase_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

	t.Run("success", func(t *testing.T) {
		newUsername := "UniqueUsername"
		newEmail := "UniqueEmail"
		var params []interface{}
		params = append(params, newUsername, newEmail)

		repo.EXPECT().SelectFullUser(consts.SELECTAllUsers, nil).Return(expectedUsers, nil)

		users, err := us.GetAllUsers()

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, users[0].Email, expectedUsers[0].Email)
		assert.Equal(t, users[1].Email, expectedUsers[1].Email)
		assert.Equal(t, users[2].Email, expectedUsers[2].Email)
	})
}

func TestPinterestUsecase_GenSessionKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		sessionKeyLenght := 20

		key, err := GenSessionKey(sessionKeyLenght)

		assert.NoError(t, err)
		assert.NotNil(t, key)
		assert.Equal(t, len(key), sessionKeyLenght)
	})
}

func TestPinterestUsecase_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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
		params = append(params, newProfile.Username, newProfile.Name, newProfile.Surname, newProfile.Password, newProfile.Email,
			newProfile.Age, newProfile.Status, user.ID)

		repo.EXPECT().Update(consts.UPDATEUserByID, gomock.Any()).Return(0, nil)

		id, err := us.SetUser(newProfile, user)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, int(user.ID))
	})
}

func TestUseStruct_AddBoard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

		repo.EXPECT().Insert(consts.INSERTBoard, params).Return("1", nil)

		id, err := us.AddBoard(newBoard)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, uint64(1))
	})
}

func TestUseStruct_AddPin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

		repo.EXPECT().Insert(consts.INSERTPin, params).Return("1", nil)

		id, err := us.AddPin(newPin)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, uint64(1))
	})
}

func TestUseStruct_SetJSONData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

func TestUseStruct_AddNotice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

		repo.EXPECT().Insert(consts.INSERTNotice, params).Return("1", nil)

		lastID, err := us.AddNotice(newNotice)

		assert.NoError(t, err)
		assert.NotNil(t, lastID)
		assert.Equal(t, lastID, uint64(1))
	})
}

func TestUseStruct_AddConment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		newComment := models.NewComment{
			Text: "Fuu",
		}
		var params []interface{}
		params = append(params, "1", newComment.Text, uint64(1), time.Now())

		repo.EXPECT().Insert(consts.INSERTComment, gomock.Any()).Return("1", nil)

		err := us.AddComment("1", uint64(1), newComment)

		assert.NoError(t, err)
	})
}

func TestUseStruct_ExtractFormatFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

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

func TestUseStruct_SearchPinsByTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		tagName := "car"

		expectedPin := []models.PinForSearchResult{
			{
				ID:     1,
				PinDir: "/dir/",
				Title:  "TheBest",
			},
		}
		var params []interface{}
		params = append(params, tagName)

		repo.EXPECT().SelectPinsByTag(consts.SELECTPinsByTag, params).Return(expectedPin, nil)

		pinst, err := us.SearchPinsByTag(tagName)

		assert.NoError(t, err)
		assert.NotNil(t, pinst)
		assert.Equal(t, pinst[0], expectedPin[0])
	})
}

func TestUseStruct_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}
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

		repo.EXPECT().SelectFullUser(consts.SELECTAllUsers, nil).Return(selectedUsers, nil)

		users, err := us.GetAllUsers()

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, users, expectedUsers)
	})
}

func TestUseStruct_GetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}
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

		repo.EXPECT().SelectFullUser(consts.SELECTUsersByUsername, params).Return(selectedUsers, nil)

		user, err := us.GetUserByUsername(username)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUsers[0], user)
	})
}

func TestUseStruct_GetUserIDByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		email := "something@mail.ru"

		var params []interface{}
		params = append(params, email)

		repo.EXPECT().SelectOneCol(consts.SELECTUserIDByEmail, params).Return([]string{"1"}, nil)

		id, err := us.GetUserIDByEmail(email)

		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, "1")
	})
}

func TestUseStruct_GetPin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		pinID := "1"
		selectedPins := []models.Pin{
			{
				OwnerID:     14,
				AuthorID:    14,
				BoardID:     1,
				PinDir:      "/die/",
				Title:       "SomeTitle",
				Description: "SomeDesc",
				CreatedTime: time.Now(),
			},
		}

		var params []interface{}
		params = append(params, pinID)

		repo.EXPECT().SelectPin(consts.SELECTPinByID, params).Return(selectedPins, nil)

		pin, err := us.GetPin(pinID)

		assert.NoError(t, err)
		assert.NotNil(t, pin)
		assert.Equal(t, pin, selectedPins[0])
	})
}

func TestUseStruct_GetBoard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		boardID := uint64(1)
		selectedBoard := models.Board{
			ID:          1,
			OwnerID:     14,
			Title:       "SomeTitle",
			Description: "SomeDesc",
			Category:    "cars",
			CreatedTime: time.Now(),
		}

		var params []interface{}
		params = append(params, boardID)

		repo.EXPECT().SelectBoard(consts.SELECTBoardByID, params).Return(selectedBoard, nil)

		board, err := us.GetBoard(boardID)

		assert.NoError(t, err)
		assert.NotNil(t, board)
		assert.Equal(t, board, selectedBoard)
	})
}

func TestUseStruct_GetPins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		boardID := uint64(1)
		expPins := []models.Pin{
			{
				OwnerID:     1,
				AuthorID:    1,
				BoardID:     1,
				PinDir:      "/die/",
				Title:       "SomeTitle",
				Description: "SomeDesc",
				CreatedTime: time.Now(),
			},
			{
				OwnerID:     14,
				AuthorID:    14,
				BoardID:     1,
				PinDir:      "/die/1",
				Title:       "SomeTitle",
				Description: "SomeDesc",
				CreatedTime: time.Now(),
			},
		}
		var params []interface{}
		params = append(params, boardID)

		repo.EXPECT().SelectPin(consts.SELECTPinsByBoardID, params).Return(expPins, nil)

		pins, err := us.GetPins(boardID)

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestUseStruct_GetNewPins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		expPins := []models.PinForMainPage{
			{
				ID:        1,
				PinDir:    "/dir/1",
				IsDeleted: false,
			},
			{
				ID:        2,
				PinDir:    "/dir/2",
				IsDeleted: false,
			},
		}
		var params []interface{}
		params = append(params, consts.NumberOfPinsOnPage)

		repo.EXPECT().SelectIDDirPins(consts.SELECTNewPinsByNumber, params).Return(expPins, nil)

		pins, err := us.GetNewPins()

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestUseStruct_GetMyPins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		userID := uint64(1)
		expPins := []models.PinForMainPage{
			{
				ID:        1,
				PinDir:    "/dir/1",
				IsDeleted: false,
			},
			{
				ID:        2,
				PinDir:    "/dir/2",
				IsDeleted: false,
			},
		}
		var params []interface{}
		params = append(params, consts.NumberOfPinsOnPage, userID)

		repo.EXPECT().SelectIDDirPins(consts.SELECTMyPinsByNumber, params).Return(expPins, nil)

		pins, err := us.GetMyPins(userID)

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestUseStruct_GetSubscribePins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		userID := uint64(1)
		expPins := []models.PinForMainPage{
			{
				ID:        1,
				PinDir:    "/dir/1",
				IsDeleted: false,
			},
			{
				ID:        2,
				PinDir:    "/dir/2",
				IsDeleted: false,
			},
		}
		var params []interface{}
		params = append(params, consts.NumberOfPinsOnPage, userID)

		repo.EXPECT().SelectIDDirPins(consts.SELECTSubscribePinsDisplayByNumber, params).Return(expPins, nil)

		pins, err := us.GetSubscribePins(userID)

		assert.NoError(t, err)
		assert.NotNil(t, pins)
		assert.Equal(t, pins, expPins)
	})
}

func TestUseStruct_GetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockReposInterface(ctrl)

	var mutex sync.Mutex
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	us := UseStruct{}
	if err := us.NewUseCase(&mutex, repo, &san, hub); err != nil {
		assert.NoError(t, err)
	}

	t.Run("success", func(t *testing.T) {
		pinID := "1"
		expComments := []models.CommentForSend{
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

		repo.EXPECT().SelectComments(consts.SELECTComments, params).Return(expComments, nil)

		comments, err := us.GetComments(pinID)

		assert.NoError(t, err)
		assert.NotNil(t, comments)
		assert.Equal(t, comments, expComments)
	})
}
