package functions

import (
	"crypto/rand"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
	"log"
	"net/http"
	"time"
)

func CreateNewUserSession(userId string) (http.Cookie, error) {

	sessionKeyValue := GenSessionKey(12)
	cookieSessionKey := http.Cookie{
		Name:    "session_key",
		Value:   sessionKeyValue,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	DBWorker := repository.DataBaseWorker{}
	DBWorker.NewDataBaseWorker()
	err := DBWorker.WriteData(repository.CombineInsertSessionQuery(userId, cookieSessionKey.Value, cookieSessionKey.Expires.String()))
	if err != nil {
		return cookieSessionKey, err
	}

	return cookieSessionKey, nil
}

func GenSessionKey(length int) string {
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
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

