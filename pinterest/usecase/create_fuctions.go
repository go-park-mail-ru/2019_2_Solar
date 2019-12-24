package usecase

import (
	"crypto/md5"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/validation"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (USC *UseStruct) NewUseCase(mu *sync.Mutex, rep repository.ReposInterface,
	san *sanitizer.SanitStruct, hub webSocket.HubStruct) {
	USC.Mu = mu
	USC.PRepository = rep
	USC.Sanitizer = san
	USC.Hub = hub
	go USC.Hub.Run()
}

func (USC UseStruct) AddNewUserSession(userID uint64) (http.Cookie, error) {
	sessionKeyValue, err := functions.GenSessionKey(12)
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


func (USC *UseStruct) AddNewUser(username, email, password string) (uint64, error) {
	salt, err := functions.GenSessionKey(10)
	if err != nil {
		return 0, err
	}
	hashPassword := functions.HashPassword(password, salt)
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

func (USC *UseStruct) AddBoard(Board models.Board) (uint64, error) {
	lastID, err := USC.PRepository.InsertBoard(Board.OwnerID, Board.Title, Board.Description, Board.Category, Board.CreatedTime)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (USC *UseStruct) AddPin(Pin models.Pin) (uint64, error) {
	lastID, err := USC.PRepository.InsertPin(Pin)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (USC *UseStruct) AddNotice(notice models.Notice) (uint64, error) {
	lastID, err := USC.PRepository.InsertNotice(notice)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (USC *UseStruct) AddComment(pinID, userID uint64, newComment models.NewComment) error {
	_, err := USC.PRepository.InsertComment(pinID, newComment.Text, userID, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (USC *UseStruct) AddSubscribe(userID uint64, followeeName string) error {
	var params []interface{}
	params = append(params, userID, followeeName)
	_, err := USC.PRepository.InsertSubscribe(userID, followeeName)
	if err != nil {
		return err
	}
	return nil
}

func (USC *UseStruct) AddTags(description string, pinID uint64) error {
	tags := validation.FindTags.FindAllString(description, -1)
	for i := 0; i < len(tags); i++ {
		//strings.Re
		tags[i] = strings.TrimPrefix(tags[i], "#")
	}
	uniqueTags, err := USC.PRepository.SelectAllTags()
	if err != nil {
		return err
	}

	alredyExitstflag := false
	for _, tag := range tags {
		for _, uniqueTag := range uniqueTags {
			if uniqueTag == tag {
				alredyExitstflag = true
			}
		}
		if alredyExitstflag != true {
			if err := USC.PRepository.InsertTag(tag); err != nil {
				return err
			}
		}
		alredyExitstflag = false
	}

	for _, tag := range tags {
		if err := USC.PRepository.InsertPinAndTag(pinID, tag); err != nil {
			return err
		}
	}

	return nil
}

func (USC *UseStruct) SetPin(pin models.EditPin, userId uint64) error {
	rowsAffected, err := USC.PRepository.UpdatePin(pin, userId)
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New("Ошибка при редактировании пина")
	}
	return nil
}

func (USC *UseStruct) AddFeedBack(newFeedBack models.NewFeedBack) error {
	err := USC.PRepository.InsertFeedBack(newFeedBack)
	if err != nil {
		return err
	}
	return nil
}
