package usecase

import (
	"crypto/rand"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/consts"
	"net/http"
	"time"
)

func (use *UseStruct) AddNewAdminSession(adminID uint64) (Cookie http.Cookie, Err error) {
	sessionKeyValue, err := GenSessionKey(12)
	id := uint64( adminID)
	if err != nil {
		return http.Cookie{}, err
	}
	cookieSessionKey := new(http.Cookie)
	cookieSessionKey.Name = "admin_session_key"
	cookieSessionKey.Value = sessionKeyValue
	cookieSessionKey.Path = "/"
	cookieSessionKey.Expires = time.Now().Add(365 * 24 * time.Hour)
	_, err = use.PRepository.InsertAdminSession(id, cookieSessionKey.Value, cookieSessionKey.Expires)
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
			var err error
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
