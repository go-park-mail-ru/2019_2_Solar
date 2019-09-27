package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
	//"github.com/stretchr/testify/assert"
)

func (h *Handlers) HandleForTest(w http.ResponseWriter, r *http.Request) {

}

func TestCreateNewUser1(t *testing.T) {

	hTest := Handlers{
		users:    make([]User, 0),
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	newUserReg := UserReg{
		Email:    "vitaly@gmail.com",
		Password: "1234",
		Username: "Vitaly",
	}

	newUserOK := User{
		ID:       0,
		Name:     "",
		Password: "1234",
		Email:    "vitaly@gmail.com",
		Username: "Vitaly",
	}

	newUser := CreateNewUser(&hTest, newUserReg)

	if newUser != newUserOK {
		t.Errorf("Test failed")
	}
}

func TestCreateNewUser2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Vitaly",
				Email:    "vitaly@gmail.com",
				Password: "1234",
				Username: "Vitaly",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	newUserReg := UserReg{
		Email:    "Ivan@gmail.com",
		Password: "424242",
		Username: "AmigoMail",
	}

	newUserOK := User{
		ID:       1,
		Name:     "",
		Password: "424242",
		Email:    "Ivan@gmail.com",
		Username: "AmigoMail",
	}

	newUser := CreateNewUser(&hTest, newUserReg)

	if newUser != newUserOK {
		t.Errorf("Test failed")
	}
}

func TestCreateNewUserSession1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d5",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	sessionsCountOK := 1

	cookie, err := CreateNewUserSession(&hTest, hTest.users[len(hTest.users)-1])

	if err != nil {
		t.Errorf("Test failed")
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		t.Errorf("Test failed")
	}
	if correctCookie.Name != "session_id" {
		t.Errorf("Test failed")
	}
	if len(hTest.sessions) < sessionsCountOK {
		t.Errorf("Test failed")
	}
	if hTest.sessions[len(hTest.sessions)-1].UserID != hTest.users[len(hTest.users)-1].ID {
		t.Errorf("Test failed")
	}
}

func TestCreateNewUserSession2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	sessionsCountOK := len(hTest.sessions) + 1

	cookie, err := CreateNewUserSession(&hTest, hTest.users[len(hTest.users)-1])

	if err != nil {
		t.Errorf("Test failed")
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		t.Errorf("Test failed")
	}
	if correctCookie.Name != "session_id" {
		t.Errorf("Test failed")
	}
	if len(hTest.sessions) < sessionsCountOK {
		t.Errorf("Test failed")
	}
	if hTest.sessions[len(hTest.sessions)-1].UserID != hTest.users[len(hTest.users)-1].ID {
		t.Errorf("Test failed")
	}
}

func TestCreateNewUserSession3(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d8",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 12,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	sessionsCountOK := len(hTest.sessions) + 1

	cookie, err := CreateNewUserSession(&hTest, hTest.users[len(hTest.users)-1])

	if err != nil {
		t.Errorf("Test failed")
	}
	correctCookie, ok := cookie.(http.Cookie)
	if !ok {
		t.Errorf("Test failed")
	}
	if correctCookie.Name != "session_id" {
		t.Errorf("Test failed")
	}
	if len(hTest.sessions) < sessionsCountOK {
		t.Errorf("Test failed")
	}
	if hTest.sessions[len(hTest.sessions)-1].UserID != hTest.users[len(hTest.users)-1].ID {
		t.Errorf("Test failed")
	}
}

func TestDeleteOldUserSession1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d8",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 12,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	sessionsCoutOK := len(hTest.sessions) - 1

	value := "5h7x"

	err := DeleteOldUserSession(&hTest, value)

	if err != nil {
		t.Errorf("Test failed")
	}
	if len(hTest.sessions) != sessionsCoutOK {
		t.Errorf("Test failed")
	}
}

func TestDeleteOldUserSession2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d8",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 12,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 5,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 16,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	sessionsCoutOK := len(hTest.sessions) - 2

	cookieValue1 := "6h7x"
	cookieValue2 := "5h7x"

	err := DeleteOldUserSession(&hTest, cookieValue1)
	if err != nil {
		t.Errorf("Test failed")
	}
	err = DeleteOldUserSession(&hTest, cookieValue2)
	if err != nil {
		t.Errorf("Test failed")
	}

	if len(hTest.sessions) != sessionsCoutOK {
		t.Errorf("Test failed")
	}
	if hTest.sessions[len(hTest.sessions)-1].UserID != 16 {
		t.Errorf("Test failed")
	}
}

func TestSearchCookieSession1(t *testing.T) {

	r := httptest.NewRequest("GET", "/", nil)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "6h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)

	session, err := SearchCookieSession(r)

	if err != nil {
		t.Errorf("Test failed")
	}
	if session.Value != "6h7x" {
		t.Errorf("Test failed")
	}
}

func TestSearchCookieSession2(t *testing.T) {

	r := httptest.NewRequest("GET", "/", nil)
	cookie := http.Cookie{
		Name:    "ABC", //incorrect name
		Value:   "6h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)

	_, err := SearchCookieSession(r)

	if err == nil {
		t.Errorf("Test failed")
	}
}

func TestEmailIsUnique1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	newUserReg := UserReg{
		Email:    "unique@mul.com",
		Password: "1001",
		Username: "jonny",
	}

	ok := EmailIsUnique(&hTest, newUserReg.Username)
	if !ok {
		t.Errorf("Test failed")
	}

}

func TestEmailIsUnique2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	newUserReg := UserReg{
		Email:    "ABC45@mail.su", // not unique
		Password: "1001",
		Username: "jonny",
	}

	ok := EmailIsUnique(&hTest, newUserReg.Email)
	if ok {
		t.Errorf("Test failed")
	}

}

func TestUserNameIsUnique1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	newUserReg := UserReg{
		Email:    "QUE45@mail.su",
		Password: "1001",
		Username: "jonny",
	}

	ok := UsernameIsUnique(&hTest, newUserReg.Username)
	if !ok {
		t.Errorf("Test failed")
	}
}

func TestUserNameIsUnique2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	newUserReg := UserReg{
		Email:    "QUE45@mail.su",
		Password: "1001",
		Username: "12d8", // not unique
	}

	ok := UsernameIsUnique(&hTest, newUserReg.Username)
	if ok {
		t.Errorf("Test failed")
	}
}

func TestSearchUserByEmail1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	newUserLogin := UserLogin{
		Email:    "ABC45@mail.su",
		Password: "abcd",
	}

	value := SearchUserByEmail(hTest.users, &newUserLogin)
	user, ok := value.(User)
	if !ok {
		t.Errorf("Test failed")
	}
	if user.Name != "Bob" {
		t.Errorf("Test failed")
	}
}

func TestExtractForamatFile1(t *testing.T) {
	fileName := "xxx.img"
	format, err := ExtractFormatFile(fileName)
	if err != nil || format != ".img" {
		t.Errorf("Test failed")
	}
}

func TestExtractForamatFile2(t *testing.T) {
	fileName := "xxximg"
	_, err := ExtractFormatFile(fileName)
	if err == nil {
		t.Errorf("Test failed")
	}
}

func TestUserIndexByID1(t *testing.T) {
	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	if index := GetUserIndexByID(&hTest, 12); index != 1 {
		t.Errorf("Test failed")
	}
}

func TestUserIndexByID2(t *testing.T) {
	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	if index := GetUserIndexByID(&hTest, 100); index != -1 {
		t.Errorf("Test failed")
	}
}

func TestHandleEmpty1(t *testing.T) {

	hTest := Handlers{
		users:    make([]User, 0),
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	r := httptest.NewRequest("PUT", "/", nil)
	w := httptest.NewRecorder()

	expectedJSON := "{}"

	hTest.HandleEmpty(w, r)

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")

	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEmpty2(t *testing.T) {

	hTest := Handlers{
		users:    make([]User, 0),
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	expectedJSON := "{}"

	hTest.HandleEmpty(w, r)

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")

	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleRegUser1(t *testing.T) {

	hTest := Handlers{
		users:    make([]User, 0),
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"email": "vitalian42@mail.ru", "password": "1234", "username": "Vitalian42"}`)

	r := httptest.NewRequest("POST", "/registration/", bodyReader)
	w := httptest.NewRecorder()

	hTest.HandleRegUser(w, r)

	if hTest.users[len(hTest.users)-1].Email != "vitalian42@mail.ru" {
		t.Errorf("Test failed")
	}
}

func TestHandleRegUser2(t *testing.T) {

	hTest := Handlers{
		users:    make([]User, 0),
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"email": "vitalian42@mail.ru", "password": "1234, "username": "Vitalian42"}`) // incorrect JSONy

	r := httptest.NewRequest("POST", "/registration/", bodyReader)
	w := httptest.NewRecorder()

	hTest.HandleRegUser(w, r)

	expectedJSON := `{"errorMessage":"incorrect json"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleRegUser3(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"email": "COM44@mail.su", "password": "1234", "username": "Vitalian42"}`) // mot unique email

	r := httptest.NewRequest("POST", "/registration/", bodyReader)
	w := httptest.NewRecorder()

	hTest.HandleRegUser(w, r)

	expectedJSON := `{"errorMessage":"not unique Email"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleListUsers1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}

	r := httptest.NewRequest("GET", "/users/", nil)
	w := httptest.NewRecorder()

	hTest.HandleListUsers(w, r)

	expectedJSON := `[{"username":"12d6","name":"Bob","surname":"","email":"NEO43@mail.su","age":"","status":"","isactive":""},{"username":"12d7","name":"Bob","surname":"","email":"COM44@mail.su","age":"","status":"","isactive":""},{"username":"12d8","name":"Bob","surname":"","email":"ABC45@mail.su","age":"","status":"","isactive":""}]`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLoginUser1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}
	bodyReader := strings.NewReader(`{"email": "COM44@mail.su", "password": "abcd"}`)

	r := httptest.NewRequest("POST", "/login/", bodyReader)
	w := httptest.NewRecorder()

	hTest.HandleLoginUser(w, r)

	expectedJSON := `{"username":"12d7","name":"Bob","surname":"","email":"COM44@mail.su","age":"","status":"","isactive":""}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLoginUser2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}
	bodyReader := strings.NewReader(`{"email": "COM44@mail.su", "password": "mypass"}`) // incorrect password

	r := httptest.NewRequest("POST", "/login/", bodyReader)
	w := httptest.NewRecorder()

	hTest.HandleLoginUser(w, r)

	expectedJSON := `{"errorMessage":"incorrect combination of Email and Password"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLoginUser3(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.ru",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "bob42@mail.com",
				Username: "12d8",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 12,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 5,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 16,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"email": "bob42@mail.su", "password": "abcd"}`)

	r := httptest.NewRequest("POST", "/login/", bodyReader)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "5h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleLoginUser(w, r)

	expectedJSON := `{"username":"12d7","name":"Bob","surname":"","email":"bob42@mail.su","age":"","status":"","isactive":""}
{"message":"successfully log in yet"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLoginUser4(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}
	bodyReader := strings.NewReader(`{"email": "COM44@mail.su, "password": "mypass"}`) // incorrect password

	r := httptest.NewRequest("POST", "/login/", bodyReader) // incorrect JSON
	w := httptest.NewRecorder()

	hTest.HandleLoginUser(w, r)

	expectedJSON := `{"errorMessage":"incorrect json"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLoginUser5(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       5,
				Name:     "Bob",
				Password: "abcd",
				Email:    "NEO43@mail.su",
				Username: "12d6",
			},
			{
				ID:       12,
				Name:     "Bob",
				Password: "abcd",
				Email:    "COM44@mail.su",
				Username: "12d7",
			},
			{
				ID:       16,
				Name:     "Bob",
				Password: "abcd",
				Email:    "ABC45@mail.su",
				Username: "12d8",
			},
		},
		sessions: make([]UserSession, 0),
		mu:       &sync.Mutex{},
	}
	bodyReader := strings.NewReader(`{"email": "_@mail.su", "password": "mypass"}`) // incorrect password

	r := httptest.NewRequest("POST", "/login/", bodyReader) // incorrect email
	w := httptest.NewRecorder()

	hTest.HandleLoginUser(w, r)

	expectedJSON := `{"errorMessage":"incorrect combination of Email and Password"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEditProfileUserData1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"username": "Andrey", "name": "Andrey", "surname": "dmitrievich", "password": "MyUniquePassword", "email": "Andrey@mail.ru", "age": "40", "status": "active", "isactive": "true"}`)

	r := httptest.NewRequest("GET", "/profile/data", bodyReader)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "6h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleEditProfileUserData(w, r)

	expectedJSON := `{"message":"data successfully saved"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEditProfileUserData2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"username": "And "name": "Andrey", "surname": "dmitrievich", "password": "MyUniquePassword", "email": "Andrey@mail.ru", "age": "40", "status": "active", "isactive": "true"}`)

	r := httptest.NewRequest("GET", "/profile/data", bodyReader) // incorrect json
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "6h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleEditProfileUserData(w, r)

	expectedJSON := `{"errorMessage":"incorrect json"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEditProfileUserData3(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"username": "Andrey", "name": "Andrey", "surname": "dmitrievich", "password": "MyUniquePassword", "email": "Andrey@mail.ru", "age": "40", "status": "active", "isactive": "true"}`)

	r := httptest.NewRequest("GET", "/profile/data", bodyReader)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "ummmm", // incorrect cookie
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleEditProfileUserData(w, r)

	expectedJSON := `{"errorMessage":"invalid cookie or user"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEditProfileUserData4(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"username": "Andrey", "name": "Andrey", "surname": "dmitrievich", "password": "MyUniquePassword", "email": "Liza@mail.com", "age": "40", "status": "active", "isactive": "true"}`)

	r := httptest.NewRequest("GET", "/profile/data", bodyReader) // not unique email
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "6h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleEditProfileUserData(w, r)

	expectedJSON := `{"errorMessage":"not unique Email"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEditProfileUserData5(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	bodyReader := strings.NewReader(`{"username": "Dima", "name": "Andrey", "surname": "dmitrievich", "password": "MyUniquePassword", "email": "Andrey@mail.ru", "age": "40", "status": "active", "isactive": "true"}`)

	r := httptest.NewRequest("GET", "/profile/data", bodyReader) // not unique username
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "6h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleEditProfileUserData(w, r)

	expectedJSON := `{"errorMessage":"not unique Username"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLogoutUser1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	r := httptest.NewRequest("GET", "/profile/data", nil)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "7h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleLogoutUser(w, r)

	expectedJSON := `{"infoMessage":"Session has been successfully deleted"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLogoutUser2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	r := httptest.NewRequest("GET", "/profile/data", nil)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "8h7x", // incorrect cookie
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleLogoutUser(w, r)

	expectedJSON := `{"errorMessage":"Session has not found"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleLogoutUser3(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	r := httptest.NewRequest("GET", "/profile/data", nil)
	w := httptest.NewRecorder()

	hTest.HandleLogoutUser(w, r)

	expectedJSON := `{"errorMessage":"Cookies have not found"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEditProfileUserPicture1(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	r := httptest.NewRequest("POST", "/", nil)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "7h7x",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleEditProfileUserPicture(w, r)

	expectedJSON := `{"errorMessage":"Cannot read profile picture"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}

func TestHandleEditProfileUserPicture2(t *testing.T) {

	hTest := Handlers{
		users: []User{
			{
				ID:       0,
				Name:     "Anton",
				Password: "123",
				Email:    "Anton@mail.ru",
				Username: "Anton",
			},
			{
				ID:       1,
				Name:     "Dima",
				Password: "abc",
				Email:    "Dima@mail.su",
				Username: "Dima",
			},
			{
				ID:       2,
				Name:     "Liza",
				Password: "xyz",
				Email:    "Liza@mail.com",
				Username: "Liza",
			},
		},
		sessions: []UserSession{
			{
				ID:     1,
				UserID: 1,
				UserCookie: UserCookie{
					Value:      "5h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     2,
				UserID: 0,
				UserCookie: UserCookie{
					Value:      "6h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
			{
				ID:     3,
				UserID: 2,
				UserCookie: UserCookie{
					Value:      "7h7x",
					Expiration: time.Now().Add(1 * time.Hour),
				},
			},
		},
		mu: &sync.Mutex{},
	}

	r := httptest.NewRequest("POST", "/", nil)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "xxxx", // incorrect cookie
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	r.AddCookie(&cookie)
	w := httptest.NewRecorder()

	hTest.HandleEditProfileUserPicture(w, r)

	expectedJSON := `{"errorMessage":"user not found or not valid cookies"}`

	bytes, _ := ioutil.ReadAll(w.Body)
	bodyJSON := strings.Trim(string(bytes), "\n")
	if bodyJSON != expectedJSON {
		t.Errorf("Test failed")
	}
}