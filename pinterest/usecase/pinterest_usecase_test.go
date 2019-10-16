package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestPinterestUsecase_CreateNewUser(t *testing.T) {
	//mockPinterestUsecase := new(mocks.MockUsecase)

	mockListUsers := make([]models.User, 0)
	mockListSession := make([]models.UserSession, 0)

	t.Run("success", func(t *testing.T) {
		mockUser := &models.UserReg{
			Username: "Vitaly",
			Email: "something@mail.ru",
			Password: "123QWErty!",

		}

		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		savedUser := uc.CreateNewUser(mockUser)

		assert.NotNil(t, savedUser)

		assert.Equal(t, mockUser.Email, savedUser.Email)
		assert.Equal(t, mockUser.Password, savedUser.Password)
		assert.Equal(t, mockUser.Username, savedUser.Username)
	})
}

func TestPinterestUsecase_CreateNewUserSession(t *testing.T) {

	user := models.User{
		ID: 0,
		Username: "Vitaly",
		Email: "something@mail.ru",
		Password: "123QWErty!",

	}

	mockListUsers := make([]models.User, 0)
	mockListUsers = append(mockListUsers, user)

	mockListSession := make([]models.UserSession, 0)

	t.Run("success", func(t *testing.T) {

		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		cookie, session, err := uc.CreateNewUserSession(user)

		assert.NoError(t, err)
		assert.NotNil(t, cookie)
		assert.NotNil(t, session)

		assert.Equal(t, session.UserID, user.ID)
		assert.Equal(t, session.UserCookie.Value, cookie[0].Value)
		assert.Equal(t, session.UserCookie.Expiration, cookie[0].Expires)
	})
}

func TestPinterestUsecase_DeleteOldUserSession(t *testing.T) {
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

	t.Run("success", func(t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		err := uc.DeleteOldUserSession("QWERTY")

		assert.NoError(t, err)

	})

	t.Run("incorrect session value", func(t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		err := uc.DeleteOldUserSession("INCORRECT")

		assert.Error(t, err)

	})
}

func TestPinterestUsecase_EditEmailIsUnique(t *testing.T) {
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

		isUnique := uc.EditEmailIsUnique("another@mail.ru", user.ID)

		assert.True(t, isUnique)
	})
}

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

func TestPinterestUsecase_EditProfileDataCheck(t *testing.T) {
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
		newProfile := models.EditUserProfile{
			Username: "Alcost",
			Name:     "Alcost",
			Surname:  "Filcost",
			Password: "123ewrEW#",
			Email:    "email@mail.su",
			Age:      "42",
			Status:   "Ok",
			IsActive: "True",
		}

		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		err := uc.EditProfileDataCheck(&newProfile)

		assert.NoError(t, err)

	})
}

func TestPinterestUsecase_GetAllUsers(t *testing.T) {
	user1 := models.User{
		ID: 0,
		Username: "Vitaly",
		Email: "something@mail.ru",
		Password: "123QWErty!",
	}
	user2 := models.User{
		ID: 1,
		Username: "Vova",
		Email: "something2@mail.ru",
		Password: "123QWErty!",
	}
	user3 := models.User{
		ID: 2,
		Username: "Nastya",
		Email: "something2@mail.ru",
		Password: "123QWErty!",
	}
	user4 := models.User{
		ID: 3,
		Username: "Bogdan",
		Email: "something3@mail.ru",
		Password: "123QWErty!",
	}

	mockListUsers := make([]models.User, 0)
	mockListUsers = append(mockListUsers, user1, user2, user3, user4)

	mockListSession := make([]models.UserSession, 0)

	t.Run("success", func (t *testing.T) {

		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})

		users := uc.GetAllUsers()

		assert.NotNil(t, users)
		assert.Equal(t, users[0].Username, user1.Username)

	})
}

func TestPinterestUsecase_GenSessionKey(t *testing.T) {
	mockListUsers := make([]models.User, 0)
	mockListSession := make([]models.UserSession, 0)

	t.Run("success", func (t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})
		sessionKeyLenght := 20

		key := uc.GenSessionKey(sessionKeyLenght)

		assert.NotNil(t, key)
		assert.Equal(t, len(key), sessionKeyLenght)
	})
}

func TestPinterestUsecase_GetUserByID(t *testing.T) {
	mockListUsers := make([]models.User, 0)
	mockListSession := make([]models.UserSession, 0)

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

	mockListUsers = append(mockListUsers, user1, user2, user3, user4)

	t.Run("success", func(t *testing.T) {
		uc := NewPinterestUsecase(mockListUsers, mockListSession, &sync.Mutex{})
		var userID uint64 = 2

		user := uc.GetUserByID(userID)

		assert.NotNil(t, user)
		assert.Equal(t, user, user3)
	})
}

func TestPinterestUsecase_SaveNewProfileUser(t *testing.T) {
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
		newProfile := models.EditUserProfile{
			Name:     "Alcost",
			Surname:  "Filcost",
			Age:      "42",
			Status:   "Ok",
			IsActive: "True",
		}

		uc.SaveNewProfileUser(user1.ID, &newProfile)
		user := uc.GetUserByID(user1.ID)

		assert.NotNil(t, user)
		assert.Equal(t, user.Name, newProfile.Name)
		assert.Equal(t, user.Status, newProfile.Status)
		assert.Equal(t, user.Age, newProfile.Age)
	})
}

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