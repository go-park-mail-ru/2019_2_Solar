package usecase

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func (USC *UseStruct) NewUseCase(mu *sync.Mutex, rep repository.ReposInterface,
	san *sanitizer.SanitStruct) {
	USC.Mu = mu
	USC.PRepository = rep
	USC.Sanitizer = san



}

func (USC UseStruct) AddNewUserSession(userID uint64) (http.Cookie, error) {
	sessionKeyValue, err := GenSessionKey(12)
	if err != nil {
		return http.Cookie{}, err
	}
	cookieSessionKey := new(http.Cookie)
	cookieSessionKey.Name = "session_key"
	cookieSessionKey.Value = sessionKeyValue
	cookieSessionKey.Path = "/"
	cookieSessionKey.Expires = time.Now().Add(365 * 24 * time.Hour)
	_, err = USC.PRepository.InsertSession(userID, cookieSessionKey.Value, cookieSessionKey.Expires)
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

func (USC *UseStruct) AddNewUser(username, email, password string) (uint64, error) {
	salt, err := GenSessionKey(10)
	if err != nil {
		return 0, err
	}
	hashPassword := HashPassword(password, salt)
	lastID, err := USC.PRepository.InsertUser(username, email, salt, hashPassword, time.Now())
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (USC *UseStruct) SetUser(newUser models.EditUserProfile, user models.User) (int, error) {
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Name != "" {
		user.Name = newUser.Name
	}
	if newUser.Surname != "" {
		user.Surname = newUser.Surname
	}
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Age != "" {
		age, err := strconv.Atoi(newUser.Age)
		if err != nil {
			return 0, err
		}
		user.Age = uint(age)
	}
	if newUser.Status != "" {
		user.Status = newUser.Status
	}
	editUsers, err := USC.PRepository.UpdateUser(user)
	if err != nil {
		return 0, err
	}
	return editUsers, nil
}

func (USC *UseStruct) SetUserAvatarDir(idUser uint64, fileName string) (int, error) {
	editUsers, err := USC.PRepository.UpdateUserAvatar(fileName, idUser)
	if err != nil {
		return 0, err
	}
	return editUsers, nil
}

func (USC *UseStruct) CalculateMD5FromFile(fileByte io.Reader) (string, error) {
	hasher := md5.New()
	if _, err := io.Copy(hasher, fileByte); err != nil {
		return "", err
	}
	fileHash := string(hasher.Sum(nil))
	fileHash = fmt.Sprintf("%x", fileHash)
	return fileHash, nil
}

func (USC *UseStruct) AddDir(folder string) error {
	if err := os.MkdirAll(folder, 0777); err != nil {
		return err
	}
	return nil
}
func (USC *UseStruct) AddPictureFile(fileName string, fileByte io.Reader) (Err error) {
	newFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := newFile.Close(); err != nil {
			Err = err
		}
	}()
	if _, err = io.Copy(newFile, fileByte); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	return nil
}
