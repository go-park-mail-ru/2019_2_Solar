package delivery

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/mocks"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlers_HandleListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUsecase(ctrl)

	users := []models.User{
		{
			Username: "Name1",
			Email:    "email1",
			Password: "pass1",
		},
		{
			Username: "Name2",
			Email:	"email1",
			Password: "pass1",
		},
	}

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			UsersJSON:  []models.User{
				{
					Username: "Name1",
					Email:    "email1",
					Password: "pass1",
				},
				{
					Username: "Name2",
					Email:	"email1",
					Password: "pass1",
				},
			},
			InfoJSON:  "OK",
		},
	}

	usecase.EXPECT().GetAllUsers().Return(users)
	usecase.EXPECT().SetJsonData(users, "OK").Return(outJson)

	handler := Handlers{usecase}
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/")

	err := handler.HandleListUsers(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"body":{`+
		`"users":[`+
		`{"username":"Name1","name":"","surname":"","email":"email1","age":"","status":"","isactive":""},`+
		`{"username":"Name2","name":"","surname":"","email":"email1","age":"","status":"","isactive":""}],`+
		`"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlers_HandleRegUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUsecase(ctrl)

	newUserReg := models.UserReg{
		Email:    "new@mail.ru",
		Password: "12345QweRTY!",
		Username: "Nova",
	}

	newUser := models.User{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "12345QweRTY!",
	}
	newUserSession := models.UserSession{
		ID:         200,
		UserID:     100,
		UserCookie: models.UserCookie{
			Value: "AAA",
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			UserJSON:  models.User{
				Email:    "new@mail.ru",
				Password: "12345QweRTY!",
				Username: "Nova",
			},
			InfoJSON:  "OK",
		},
	}

	cookies := []http.Cookie{
		{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
		},
	}

	usecase.EXPECT().RegDataCheck(&newUserReg).Return(nil)
	usecase.EXPECT().RegEmailIsUnique(newUserReg.Email).Return(true)
	usecase.EXPECT().RegUsernameIsUnique(newUserReg.Username).Return(true)
	usecase.EXPECT().CreateNewUser(&newUserReg).Return(newUser)
	usecase.EXPECT().CreateNewUserSession(newUser).Return(cookies, newUserSession, nil)
	usecase.EXPECT().SetJsonData(newUser, "OK").Return(outJson)

	handler := Handlers{usecase}
	e := echo.New()

	bodyReader := strings.NewReader(`{"email": "new@mail.ru", "password": "12345QweRTY!", "username": "Nova"}`)

	req := httptest.NewRequest(http.MethodGet, "/registration/", bodyReader)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/registration/")

	err := handler.HandleRegUser(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"body":{`+
		`"user":`+
		`{"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":"","status":"","isactive":""},`+
		`"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlers_HandleLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUsecase(ctrl)

	newUserLogin := models.UserLogin{
		Email:    "new@mail.ru",
		Password: "12345QweRTY!",
	}

	user := models.User{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "12345QweRTY!",
	}
	newUserSession := models.UserSession{
		ID:         200,
		UserID:     100,
		UserCookie: models.UserCookie{
			Value: "AAA",
			Expiration: time.Now().Add(1 * time.Hour),
		},
	}

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			UserJSON:  models.User{
				Email:    "new@mail.ru",
				Password: "12345QweRTY!",
				Username: "Nova",
			},
			InfoJSON:  "OK",
		},
	}

	cookies := []http.Cookie{
		{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
		},
	}

	handler := Handlers{usecase}
	e := echo.New()

	bodyReader := strings.NewReader(`{"email": "new@mail.ru", "password": "12345QweRTY!"}`)

	req := httptest.NewRequest(http.MethodPost, "/login/", bodyReader)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/login/")

	usecase.EXPECT().SearchUserByEmail(&newUserLogin).Return(user)
	usecase.EXPECT().SearchIdUserByCookie(c.Request()).Return(user.ID, errors.New("idUser not found"))
	usecase.EXPECT().CreateNewUserSession(user).Return(cookies, newUserSession, nil)
	usecase.EXPECT().SetJsonData(user, "OK").Return(outJson)

	err := handler.HandleLoginUser(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"body":{`+
		`"user":`+
		`{"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":"","status":"","isactive":""},`+
		`"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlers_HandleEditProfileUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUsecase(ctrl)

	newProfileUser := models.EditUserProfile{
		Name: "Andrey",
		Surname: "dmitrievich",
	}

	user := models.User{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "new@mail.ru",
	}

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			InfoJSON:  "data successfully saved",
		},
	}

	cookies := []http.Cookie{
		{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
		},
	}

	handler := Handlers{usecase}
	e := echo.New()

	bodyReader := strings.NewReader(`{"name": "Andrey", "surname": "dmitrievich"}`)

	req := httptest.NewRequest(http.MethodPost, "/profile/data", bodyReader)
	req.AddCookie(&cookies[0])
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/profile/data")

	usecase.EXPECT().EditProfileDataCheck(&newProfileUser).Return(nil)
	usecase.EXPECT().SearchIdUserByCookie(c.Request()).Return(user.ID, nil)
	usecase.EXPECT().EditUsernameIsUnique(newProfileUser.Username, user.ID).Return(true)
	usecase.EXPECT().EditEmailIsUnique(newProfileUser.Email, user.ID).Return(true)
	usecase.EXPECT().SaveNewProfileUser(user.ID, &newProfileUser).Return()
	usecase.EXPECT().SetJsonData(nil, "data successfully saved").Return(outJson)

	err := handler.HandleEditProfileUserData(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"body":{"info":"data successfully saved"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlers_HandleGetProfileUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUsecase(ctrl)

	user := models.User{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "new@mail.ru",
		Name: "Andrey",
		Surname: "dmitrievich",
	}

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			UserJSON: user,
			InfoJSON:  "OK",
		},
	}

	cookies := []http.Cookie{
		{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
		},
	}

	handler := Handlers{usecase}
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/profile/data", nil)
	req.AddCookie(&cookies[0])
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/profile/data")

	usecase.EXPECT().SearchIdUserByCookie(c.Request()).Return(user.ID, nil)
	usecase.EXPECT().GetUserByID(user.ID).Return(user)
	usecase.EXPECT().SetJsonData(user, "OK").Return(outJson)

	err := handler.HandleGetProfileUserData(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"body":{`+
		`"user":{"username":"Nova","name":"Andrey","surname":"dmitrievich","email":"new@mail.ru","age":"","status":"","isactive":""},`+
		`"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlers_HandleLogoutUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUsecase(ctrl)

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			InfoJSON:  "Session has been successfully deleted",
		},
	}

	cookie := http.Cookie{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
		}

	handler := Handlers{usecase}
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/profile/data", nil)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/profile/data")

	usecase.EXPECT().SearchCookie(c.Request()).Return(&cookie, nil)
	usecase.EXPECT().DeleteOldUserSession(cookie.Value).Return(nil)
	usecase.EXPECT().SetJsonData(nil, "Session has been successfully deleted").Return(outJson)

	err := handler.HandleLogoutUser(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"body":{"info":"Session has been successfully deleted"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}


func TestHandlers_HandleEditProfileUserPicture(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUsecase(ctrl)

	user := models.User{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "new@mail.ru",
		Name: "Andrey",
		Surname: "dmitrievich",
	}

	cookies := []http.Cookie{
		{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
		},
	}

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			InfoJSON:  "Cannot read profile picture",
		},
	}

	handler := Handlers{usecase}
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/profile/picture", nil)
	req.AddCookie(&cookies[0])
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/profile/picture")

	usecase.EXPECT().SearchIdUserByCookie(c.Request()).Return(user.ID, nil)
	usecase.EXPECT().SetJsonData(nil, "Cannot read profile picture").Return(outJson)

	err := handler.HandleEditProfileUserPicture(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"body":{"info":"Cannot read profile picture"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}
