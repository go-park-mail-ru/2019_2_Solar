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

	expectedJSON := `{"csrf_token":"","body":{"users":[{"username":"Name1","name":"","surname":"","email":"email1","age":0,"status":"","avatar_dir":"","is_active":false,"created_time":"0001-01-01T00:00:00Z"},{"username":"Name2","name":"","surname":"","email":"email1","age":0,"status":"","avatar_dir":"","is_active":false,"created_time":"0001-01-01T00:00:00Z"}],"info":"OK"}}`

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

	expectedJSON := `{"csrf_token":"","body":{"user":{"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false,"created_time":"0001-01-01T00:00:00Z"},"info":"OK"}}`

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
	usecase.EXPECT().ComparePassword(user.Password, gomock.Any(), newUserLogin.Password)
	usecase.EXPECT().AddNewUserSession(strconv.Itoa(int(user.ID))).Return(cookie, nil)
	usecase.EXPECT().SetJSONData(user, "", "OK").Return(outJson)

	err := handlers.HandleLoginUser(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"user":{"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false,"created_time":"0001-01-01T00:00:00Z"},"info":"OK"}}`

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

	expectedJSON := `{"csrf_token":"","body":{"user":{"username":"Nova","name":"","surname":"","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false,"created_time":"0001-01-01T00:00:00Z"},"info":"OK"}}`

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

	expectedJSON := `{"csrf_token":"","body":{"user":{"username":"Nova","name":"Andrey","surname":"dmitrievich","email":"new@mail.ru","age":0,"status":"","avatar_dir":"","is_active":false,"created_time":"0001-01-01T00:00:00Z"},"info":"OK"}}`

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

func TestHandlersStruct_HandleCreateBoard(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	newBoard := models.NewBoard{
		Title:       "title",
		Description: "desc",
		Category:    "cars",
	}

	//board := models.Board{
	//	OwnerID:     user.ID,
	//	Title:       newBoard.Title,
	//	Description: newBoard.Description,
	//	Category:    newBoard.Category,
	//	CreatedTime: time.Now(),
	//}

	boardID := uint64(1)

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodPost, "/board", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/board")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().CheckBoardData(newBoard).Return(nil)
	usecase.EXPECT().AddBoard(gomock.Any()).Return(boardID, nil)

	err := handlers.HandleCreateBoard(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	//expectedJSON := `{"csrf_token":"","body":{"board":{"id":1,"owner_id":100,"title":"title","description":"desc","category":"cars","created_time":"2019-11-06T12:33:01.248573142+03:00","is_deleted":false},"info":"data successfully saved"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn == "" {
		t.Errorf("Body is nil: %s", err)
	}
}

func TestHandlersStruct_HandleGetBoard(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	newBoard := models.NewBoard{
		Title:       "title",
		Description: "desc",
		Category:    "cars",
	}

	board := models.Board{
		OwnerID:     user.ID,
		Title:       newBoard.Title,
		Description: newBoard.Description,
		Category:    newBoard.Category,
	}

	pins := []models.Pin{
		{
			OwnerID:     14,
			AuthorID:    14,
			BoardID:     1,
			PinDir:      "/die/",
			Title:       "SomeTitle",
			Description: "SomeDesc",
		},
		{
			OwnerID:     1,
			AuthorID:    1,
			BoardID:     1,
			PinDir:      "/die/",
			Title:       "SomeTitle",
			Description: "SomeDesc",
		},
	}

	boardID := uint64(1)

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodGet, "/board/1", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/board/1")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().GetBoard(boardID).Return(board, nil)
	usecase.EXPECT().GetPins(boardID).Return(pins, nil)

	err := handlers.HandleGetBoard(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	//expectedJSON := `{"csrf_token":"","body":{"board":{"id":0,"owner_id":100,"title":"title","description":"desc","category":"cars","created_time":"0001-01-01T00:00:00Z","is_deleted":false},"pins":[{"id":0,"owner_id":14,"author_id":14,"board_id":1,"pin_dir":"/die/","title":"SomeTitle","description":"SomeDesc","created_time":"0001-01-01T00:00:00Z","is_deleted":false},{"id":0,"owner_id":1,"author_id":1,"board_id":1,"pin_dir":"/die/","title":"SomeTitle","description":"SomeDesc","created_time":"0001-01-01T00:00:00Z","is_deleted":false}],"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn == "" {
		t.Errorf("body is nil: %s", err)
	}
}

/*
func TestHandlersStruct_HandleUpgradeWebSocket(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodGet, "/chat", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/chat")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	ws, err := webSocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)

	usecase.EXPECT().CreateClient(ws, user.ID)

	err = handlers.HandleUpgradeWebSocket(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}
}
*/

func TestHandlersStruct_HandleCreateNotice(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	//newNotice := models.NewNotice{
	//	Message: "Hello",
	//}
	//
	//notice := models.Notice{
	//	ID:          0,
	//	UserID:      0,
	//	ReceiverID:  0,
	//	Message:     "",
	//	CreatedTime: time.Time{},
	//	IsRead:      false,
	//}

	lastID := uint64(1)

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodPost, "/notice/1", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/notice/1")
	c.SetParamNames("receiver_id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().AddNotice(gomock.Any()).Return(lastID, nil)

	err := handlers.HandleCreateNotice(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	//expectedJSON := `{"csrf_token":"","body":{"board":{"id":0,"owner_id":100,"title":"title","description":"desc","category":"cars","created_time":"0001-01-01T00:00:00Z","is_deleted":false},"pins":[{"id":0,"owner_id":14,"author_id":14,"board_id":1,"pin_dir":"/die/","title":"SomeTitle","description":"SomeDesc","created_time":"0001-01-01T00:00:00Z","is_deleted":false},{"id":0,"owner_id":1,"author_id":1,"board_id":1,"pin_dir":"/die/","title":"SomeTitle","description":"SomeDesc","created_time":"0001-01-01T00:00:00Z","is_deleted":false}],"info":"OK"}}`

	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn == "" {
		t.Errorf("body is nil: %s", err)
	}
}

func TestHandlersStruct_HandleGetPin(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	pin := models.Pin{
		OwnerID:     1,
		AuthorID:    1,
		BoardID:     1,
		PinDir:      "/die/",
		Title:       "SomeTitle",
		Description: "SomeDesc",
	}

	expComments := []models.CommentForSend{
		{
			Text:        "blablalb",
			//CreatedTime: time.Now(),
			Author:      "Iam",
		},
		{
			Text:        "blablalb2",
			//CreatedTime: time.Now(),
			Author:      "Iam2",
		},
	}

	//newNotice := models.NewNotice{
	//	Message: "Hello",
	//}
	//
	//notice := models.Notice{
	//	ID:          0,
	//	UserID:      0,
	//	ReceiverID:  0,
	//	Message:     "",
	//	CreatedTime: time.Time{},
	//	IsRead:      false,
	//}

	pinID := "1"


	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	//bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodGet, "/pin/1", nil)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/pin/1")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().GetPin(pinID).Return(pin, nil)
	usecase.EXPECT().GetComments(pinID).Return(expComments, nil)

	err := handlers.HandleGetPin(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":[{"id":0,"owner_id":1,"author_id":1,"board_id":1,"pin_dir":"/die/","title":"SomeTitle","description":"SomeDesc","created_time":"0001-01-01T00:00:00Z","is_deleted":false},[{"text":"blablalb","created_time":"0001-01-01T00:00:00Z","author_username":"Iam"},{"text":"blablalb2","created_time":"0001-01-01T00:00:00Z","author_username":"Iam2"}]]}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleGetNewPins(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	pins := []models.PinForMainPage{
		{
			ID: 1,
			PinDir:      "/die/1",
		},
		{
			ID: 2,
			PinDir:      "/die/2",
		},
	}

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	//bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodGet, "/pin/list/new", nil)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/pin/list/new")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().GetNewPins().Return(pins, nil)

	err := handlers.HandleGetNewPins(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":[{"id":1,"pin_dir":"/die/1","is_deleted":false},{"id":2,"pin_dir":"/die/2","is_deleted":false}]}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleGetMyPins(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	pins := []models.PinForMainPage{
		{
			ID: 1,
			PinDir:      "/die/1",
		},
		{
			ID: 2,
			PinDir:      "/die/2",
		},
	}

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	//bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodGet, "/pin/list/my", nil)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/pin/list/my")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().GetMyPins(user.ID).Return(pins, nil)

	err := handlers.HandleGetMyPins(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":[{"id":1,"pin_dir":"/die/1","is_deleted":false},{"id":2,"pin_dir":"/die/2","is_deleted":false}]}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleGetSubscribePins(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	pins := []models.PinForMainPage{
		{
			ID: 1,
			PinDir:      "/die/1",
		},
		{
			ID: 2,
			PinDir:      "/die/2",
		},
	}

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	//bodyReader := strings.NewReader(`{"title": "title", "description": "desc", "category": "cars"}`)

	req := httptest.NewRequest(http.MethodGet, "/pin/list/my", nil)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/pin/list/my")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().GetSubscribePins(user.ID).Return(pins, nil)

	err := handlers.HandleGetSubscribePins(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":[{"id":1,"pin_dir":"/die/1","is_deleted":false},{"id":2,"pin_dir":"/die/2","is_deleted":false}]}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleCreateComment(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	pinID := "1"

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"text": "blabla"}`)

	req := httptest.NewRequest(http.MethodPost, "/pin/1/comment", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/pin/1/comment")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().AddComment(pinID, user.ID, gomock.Any()).Return(nil)

	err := handlers.HandleCreateComment(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":"data successfully saved"}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleEditProfileUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUseInterface(ctrl)

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			InfoJSON:  "data successfully saved",
		},
	}

	user := models.User{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "new@mail.ru",
		Name: "Andrey",
		Surname: "dmitrievich",
	}

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	newUserProfile := models.EditUserProfile{
		Username: "Nova",
		Name:     "Nova",
		Surname:  "NHJFewl",
	}

	editStrings := 1

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"username":"Nova", "name":"Nova", "surname":"NHJFewl"}`)

	req := httptest.NewRequest(http.MethodPost, "/profile/data", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/profile/data")
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().CheckProfileData(gomock.Any()).Return(nil)
	usecase.EXPECT().CheckUsernameEmailIsUnique(newUserProfile.Username, newUserProfile.Email,
		user.Username, user.Email, user.ID).Return(nil)
	usecase.EXPECT().SetUser(newUserProfile, user).Return(editStrings, nil)
	usecase.EXPECT().SetJSONData(nil, "","data successfully saved").Return(outJson)

	err := handlers.HandleEditProfileUserData(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"info":"data successfully saved"}}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandlerFindPinByTag(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	pins := []models.PinForSearchResult{
		{
			ID: 1,
			Title: "CoolPin",
			PinDir: "/die/1",
		},
		{
			ID: 2,
			Title: "BadPin",
			PinDir:      "/die/2",
		},
	}

	tagName := "car"

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"username":"Nova", "name":"Nova", "surname":"NHJFewl"}`)

	req := httptest.NewRequest(http.MethodGet, "/find/pins/by/tag/car", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/find/pins/by/tag/car")
	c.SetParamNames("tag")
	c.SetParamValues("car")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().SearchPinsByTag(tagName).Return(pins, nil)

	err := handlers. HandlerFindPinByTag(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":[{"id":1,"pin_dir":"/die/1","title":"CoolPin"},{"id":2,"pin_dir":"/die/2","title":"BadPin"}]}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleCreateSubscribe(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	followeeName := "Dmitry"

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"username":"Nova", "name":"Nova", "surname":"NHJFewl"}`)

	req := httptest.NewRequest(http.MethodPost, "/subscribe/Dmitry", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/subscribe/Dmitry")
	c.SetParamNames("username")
	c.SetParamValues("Dmitry")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().AddSubscribe(strconv.FormatUint(user.ID, 10), followeeName).Return(nil)

	err := handlers.HandleCreateSubscribe(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":"data successfully saved"}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleDeleteSubscribe(t *testing.T) {
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

	cookie := http.Cookie{
		Name:    "session_key",
		Value:   "AAA",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	followeeName := "Dmitry"

	e := echo.New()
	handlers := HandlersStruct{}
	handlers.NewHandlers(e, usecase)

	bodyReader := strings.NewReader(`{"username":"Nova", "name":"Nova", "surname":"NHJFewl"}`)

	req := httptest.NewRequest(http.MethodDelete, "/subscribe/Dmitry", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/subscribe/Dmitry")
	c.SetParamNames("username")
	c.SetParamValues("Dmitry")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().RemoveSubscribe(strconv.FormatUint(user.ID, 10), followeeName).Return(nil)

	err := handlers.HandleDeleteSubscribe(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"Body":"data successfully deleted"}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandlersStruct_HandleEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mocks.NewMockUseInterface(ctrl)

	outJson := models.OutJSON{
		BodyJSON: models.DataJSON{
			InfoJSON:  "Empty handler has been done",
		},
	}

	user := models.User{
		ID:        100,
		Username:  "Nova",
		Password:  "12345QweRTY!",
		Email:     "new@mail.ru",
		Name: "Andrey",
		Surname: "dmitrievich",
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

	bodyReader := strings.NewReader(`{"username":"Nova", "name":"Nova", "surname":"NHJFewl"}`)

	req := httptest.NewRequest(http.MethodDelete, "/subscribe/Dmitry", bodyReader)
	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/subscribe/Dmitry")
	c.SetParamNames("username")
	c.SetParamValues("Dmitry")
	c.Set("User", user)
	c.Set("token", "")

	usecase.EXPECT().SetJSONData(nil, "", "Empty handler has been done").Return(outJson)

	err := handlers.HandleEmpty(c)

	if err != nil {
		t.Errorf("err is not nil: %s", err)
	}

	expectedJSON := `{"csrf_token":"","body":{"info":"Empty handler has been done"}}`
	bytes, _ := ioutil.ReadAll(rec.Body)
	bodyJSOn := strings.Trim(string(bytes), "\n")

	fmt.Println(string(bodyJSOn))

	if bodyJSOn != expectedJSON {
		t.Errorf("Test failed")
	}
}