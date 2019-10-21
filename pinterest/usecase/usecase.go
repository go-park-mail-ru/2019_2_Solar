package usecase

/*import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"net/http"
)

func SaveUserPictureDir(userID uint64, fileName string) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	p.Users[userID].AvatarDir = fileName
}

func DeleteOldUserSession(value string) error {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for i, session := range p.Sessions {
		if session.Value == value {
			p.Sessions = append(p.Sessions[:i], p.Sessions[i+1:]...)
			return nil
		}
	}
	return errors.New("session has not found")
}

func SearchCookie(r *http.Request) (*http.Cookie, error) {
	key, err := r.Cookie("session_key")
	return key, err
}


func SearchUserByEmail(newUserLogin *models.UserLogin) interface{} {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, user := range p.Users {
		if user.Email == newUserLogin.Email {
			return user
		}
	}
	return ""
}

func GetUserIndexByID(id uint64) int {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for index, user := range p.Users {
		if user.ID == id {
			return index
		}
	}
	return -1
}

func GetAllUsers() []models.User {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	return p.Users
}

func GetUserByID(id uint64) models.User {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	return p.Users[id]
}


func SearchIdUserByCookie(r *http.Request) (uint64, error) {
	p.Mu.Lock()
	defer p.Mu.Unlock()


	sessionKey, err := p.SearchCookie(r)
	if err == http.ErrNoCookie {
		return 0, errors.New("cookies not found")
	}

	for _, oneSession := range p.Sessions {
		if oneSession.Value == sessionKey.Value {
			return oneSession.UserID, nil
		}
	}
	return 0, errors.New("idUser not found")
}

func SaveNewProfileUser(userID uint64, newUser *models.EditUserProfile) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	user := p.Users[userID]

	user.Age = newUser.Age
	user.Status = newUser.Status
	user.Name = newUser.Name
	user.Surname = newUser.Surname

	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
	}

	p.Users[userID] = user
}

func ExtractFormatFile(FileName string) (string, error) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for i := 0; i < len(FileName); i++ {
		if string(FileName[i]) == "." {
			return FileName[i:], nil
		}
	}
	return "", errors.New("invalid file name")
}



func EditProfileDataCheck(newProfileUser *models.EditUserProfile) error {
	if newProfileUser.Email != "" {
		if err := functions.EmailCheck(newProfileUser.Email); err != nil {
			return err
		}
	}
	if newProfileUser.Username != "" {
		if err := functions.UsernameCheck(newProfileUser.Username); err != nil {
			return err
		}
	}
	if newProfileUser.Password != "" {
		if err := functions.PasswordCheck(newProfileUser.Password); err != nil {
			return err
		}
	}
	if newProfileUser.Name != "" {
		if err := functions.NameCheck(newProfileUser.Name); err != nil {
			return err
		}
	}
	if newProfileUser.Surname != "" {
		if err := functions.SurnameCheck(newProfileUser.Surname); err != nil {
			return err
		}
	}
	if newProfileUser.Status != "" {
		if err := functions.StatusCheck(newProfileUser.Status); err != nil {
			return err
		}
	}
	if newProfileUser.Age != "" {
		if err := functions.AgeCheck(newProfileUser.Age); err != nil {
			return err
		}
	}
	return nil
}*/
