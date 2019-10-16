package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"net/http"
	"sync"
	"time"
)

func (USC *UsecaseStruct) NewUseCase(mu *sync.Mutex, IRepository repository.RepositoryInterface) {
	USC.Mu = mu
	USC.PRepository = IRepository
}

func (USC *UsecaseStruct) CreateNewUserSession(userId string) (http.Cookie, error) {

	sessionKeyValue := USC.GenSessionKey(12)
	cookieSessionKey := http.Cookie{
		Name:    "session_key",
		Value:   sessionKeyValue,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	var params []interface{}
	params = append(params, userId)
	params = append(params, cookieSessionKey.Value)
	params = append(params, cookieSessionKey.Expires)
	err := USC.PRepository.WriteData(consts.InsertSessionQuery, params)
	if err != nil {
		return cookieSessionKey, err
	}

	return cookieSessionKey, nil
}

func (USC *UsecaseStruct) GenSessionKey(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & consts.LetterIdxMask); idx < len(consts.LetterBytes) {
			result[i] = consts.LetterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func (USC *UsecaseStruct) SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

func (USC *UsecaseStruct) InsertNewUser(username, email, password string) error {
	var params []interface{}
	params = append(params, username)
	params = append(params, email)
	params = append(params, password)
	err := USC.PRepository.WriteData(consts.InsertRegistrationQuery, params)
	if err != nil {
		return err
	}
	return nil
}
