package functions

import (
	"crypto/rand"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
	"log"
)

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

func UsernameCheck(username string) error {
	if len(username) >= 1 && len(username) <= 30 && validation.UsernameIsCorrect.MatchString(username) {
		return nil
	}
	return errors.New("Incorrect username")
}

func EmailCheck(email string) error {
	if validation.EmailIsCorrect.MatchString(email) {
		return nil
	}
	return errors.New("Incorrect email")
}

func PasswordCheck(password string) error {
	if len(password) >= 8 && len(password) <= 30 && validation.PasswordIsCorrect.MatchString(password) {
		if validation.PasswordHasAperCaseChar.MatchString(password) {
			if validation.PasswordHasDownCaseChar.MatchString(password) {
				if validation.PasswordHasSpecChar.MatchString(password) {
					return nil
				}
				return errors.New("Password has not special symbol")
			}
			return errors.New("Password has not symbol in down case")
		}
		return errors.New("Password has not symbol in upper case")
	}
	return errors.New("Incorrect password")
}

func NameCheck(name string) error {
	if len(name) >= 1 && len(name) <= 30 && validation.NameIsCorrect.MatchString(name) {
		return nil
	}
	return errors.New("Incorrct name")
}

func SurnameCheck(surname string) error {
	if len(surname) >= 1 && len(surname) <= 30 && validation.SurnameIsCorrect.MatchString(surname) {
		return nil
	}
	return errors.New("Incorrect surname")
}

func AgeCheck(age string) error {
	if validation.AgeIsCorrect.MatchString(age) {
		return nil
	}
	return errors.New("Incorrect age")
}

func StatusCheck(status string) error {
	if len(status) >= 1 && len(status) <= 200 && validation.StatusIsCorrect.MatchString(status) {
		return nil
	}
	return errors.New("Incorrect status")
}
