package usecase

import (
	"crypto/rand"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"net/http"
	"sync"
	"time"
)

func (USC *UsecaseStruct) NewUseCase(mu *sync.Mutex, IRepository repository.RepositoryInterface) {
	USC.Mu = mu
	USC.PRepository = IRepository
}

func (USC UsecaseStruct) CreateNewUserSession(userId string) (http.Cookie, error) {
	sessionKeyValue, err := GenSessionKey(12)
	if err != nil {
		return http.Cookie{}, err
	}
	cookieSessionKey := new(http.Cookie)
	cookieSessionKey.Name = "session_key"
	cookieSessionKey.Value = sessionKeyValue
	cookieSessionKey.Path = "/"
	cookieSessionKey.Expires = time.Now().Add(1 * time.Hour)
	var params []interface{}
	params = append(params, userId)
	params = append(params, cookieSessionKey.Value)
	params = append(params, cookieSessionKey.Expires)
	_, err = USC.PRepository.WriteData(consts.InsertSessionQuery, params)
	if err != nil {
		return *cookieSessionKey, err
	}
	return *cookieSessionKey, nil
}

func GenSessionKey(length int) (string, error) {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			var err error = nil
			randomBytes, err = SecureRandomBytes(bufferSize)
			if err != nil {
				return "", err
			}
		}
		if idx := int(randomBytes[j%length] & consts.LetterIdxMask); idx < len(consts.LetterBytes) {
			result[i] = consts.LetterBytes[idx]
			i++
		}
	}

	return string(result), nil
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) ([]byte, error) {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return []byte(""), err
	}
	return randomBytes, nil
}

func (USC UsecaseStruct) InsertNewUser(username, email, password string) (string, error) {
	var params []interface{}
	params = append(params, username, email, password)
	lastId, err := USC.PRepository.WriteData(consts.InsertRegistrationQuery, params)
	if err != nil {
		return "", err
	}
	return lastId, nil
}

func (USC UsecaseStruct)UpdateUser(user models.User, userId uint64) (string, error) {
	var params []interface{}
	params = append(params, user.Username)
	params = append(params, user.Email)
	params = append(params, user.Password)
	lastId, err := USC.PRepository.WriteData(consts.InsertRegistrationQuery, params)
	if err != nil {
		return "", err
	}
	return lastId, nil
}
