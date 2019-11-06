package delivery

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/mocks"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHandlers_HandleListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUseInterface(ctrl)

	users := []models.AnotherUser{
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

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/login/")
	c.Set("token", "")

	usecase.EXPECT().GetAllUsers().Return(users, nil)
	usecase.EXPECT().SetJSONData(users, gomock.Any(), "OK").Return(outJson)

	err := handlers.HandleListUsers(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"users":[{"id":0,"username":"Name1","name":"","surname":"","email":"email1","age":0,"status":"","avatar_dir":"","is_active":false},{"id":0,"username":"Name2","name":"","surname":"","email":"email1","age":0,"status":"","avatar_dir":"","is_active":false}],"info":"OK"}}`


	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))
	fmt.Println(string(expectedJSON))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlers_HandleRegUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUseInterface(ctrl)

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

	cookie := http.Cookie{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
	}

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"email": "new@mail.ru", "password": "12345QweRTY!", "username": "Nova"}`)

	req := httptest.NewRequest(http.MethodGet, "/registration", bodyReader)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/registration")
	c.Set("token", "")

	usecase.EXPECT().CheckRegData(&newUserReg).Return(nil)
	usecase.EXPECT().CheckRegUsernameEmailIsUnique(newUserReg.Username, newUserReg.Email).Return(nil)
	usecase.EXPECT().AddNewUser(newUserReg.Username, newUserReg.Email, newUserReg.Password).Return("100", nil)
	usecase.EXPECT().AddNewUserSession(strconv.Itoa(int(newUser.ID))).Return(cookie, nil)
	usecase.EXPECT().SetJSONData(&newUserReg, "", "OK").Return(outJson)

	err := handlers.HandleRegUser(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"user":{"id":0,"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false},"info":"OK"}}`

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

	usecase := mocks.NewMockUseInterface(ctrl)

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

	cookie := http.Cookie{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
	}

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"email": "new@mail.ru", "password": "12345QweRTY!"}`)

	req := httptest.NewRequest(http.MethodPost, "/login", bodyReader)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/login")
	c.Set("token", "")

	usecase.EXPECT().GetUserByEmail(newUserLogin.Email).Return(user, nil)
	usecase.EXPECT().AddNewUserSession(strconv.Itoa(int(user.ID))).Return(cookie, nil)
	usecase.EXPECT().SetJSONData(user, "", "OK").Return(outJson)

	err := handlers.HandleLoginUser(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"user":{"id":0,"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false},"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlers_HandleGetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUseInterface(ctrl)


	user := models.AnotherUser{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "12345QweRTY!",
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

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"email": "new@mail.ru", "password": "12345QweRTY!"}`)

	req := httptest.NewRequest(http.MethodGet, "/users/Nova", bodyReader)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/Nova")
	//c.Set("email", "new@mail.ru")
	c.SetParamNames("username")
	c.SetParamValues("Nova")
	c.Set("token", "")

	usecase.EXPECT().GetUserByUsername("Nova").Return(user, nil)
	usecase.EXPECT().SetJSONData(user, "", "OK").Return(outJson)

	err := handlers.HandleGetUserByUsername(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"user":{"id":0,"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false},"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

/*
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
	usecase.EXPECT().SetJSONData(nil, "data successfully saved").Return(outJson)

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
*/

func TestHandlers_HandleGetProfileUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUseInterface(ctrl)

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

	cookie := http.Cookie{
			Name:    "session_key",
			Value:   "AAA",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Hour),
	}

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	req := httptest.NewRequest(http.MethodGet, "/profile/data", nil)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/profile/data")
	c.Set("User", user)
	c.Set("Cookie", cookie)
	c.Set("token", "")


	usecase.EXPECT().SetJSONData(user, "", "OK").Return(outJson)

	err := handlers.HandleGetProfileUserData(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"user":{"id":100,"username":"Nova","name":"Andrey","surname":"dmitrievich","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false},"info":"OK"}}`

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

	usecase := mocks.NewMockUseInterface(ctrl)

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

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/logout")
	c.Set("token", "")

	usecase.EXPECT().RemoveOldUserSession(cookie.Value).Return(nil)
	usecase.EXPECT().SetJSONData(nil, "","Session has been successfully deleted").Return(outJson)

	err := handlers.HandleLogoutUser(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"info":"Session has been successfully deleted"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

