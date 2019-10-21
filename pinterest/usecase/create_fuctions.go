package usecase

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
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

func (USC *UsecaseStruct) NewUseCase(mu *sync.Mutex, IRepository repository.RepositoryInterface) {
	USC.Mu = mu
	USC.PRepository = IRepository
}

func (USC UsecaseStruct) AddNewUserSession(userId string) (http.Cookie, error) {
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
	_, err = USC.PRepository.Insert(consts.INSERTSession, params)
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

func (USC UsecaseStruct) AddNewUser(username, email, password string) (string, error) {
	var params []interface{}
	params = append(params, username, email, password)
	lastId, err := USC.PRepository.Insert(consts.INSERTRegistration, params)
	if err != nil {
		return "", err
	}
	return lastId, nil
}

func (USC *UsecaseStruct) SetUser(newUser models.EditUserProfile, user models.User) (int, error) {
	var params []interface{}
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
		user.Age =uint(age)
	}
	params = append(params, user.Username, user.Name, user.Surname, user.Password, user.Email, user.Age, user.Status, user.ID)
	editUsers, err := USC.PRepository.Update(consts.UPDATEUserByID, params)
	if err != nil {
		return 0, err
	}
	return editUsers, nil
}

func (USC *UsecaseStruct) SetUserAvatarDir(idUser, fileName string) (int, error) {
	var params []interface{}
	params = append(params, fileName, idUser)
	editUsers, err := USC.PRepository.Update(consts.UPDATEUserAvatarDirByID, params)
	if err != nil {
		return 0, err
	}
	return editUsers, nil
}

func (USC *UsecaseStruct) CalculateMD5FromFile(fileByte io.Reader) (string, error) {
	hasher := md5.New()
	if _, err := io.Copy(hasher, fileByte); err != nil {
		return "", err
	}
	fileHash := string(hasher.Sum(nil))
	fileHash = fmt.Sprintf("%x", fileHash)
	return fileHash, nil
}

func (USC *UsecaseStruct) AddDir(folder string) error {
	if err := os.MkdirAll(folder, 0777); err != nil {
		return err
	}
	return nil
}
func (USC *UsecaseStruct) AddPictureFile(fileName string, fileByte io.Reader) (Err error) {
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
