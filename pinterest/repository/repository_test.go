package repository

import (
	_ "database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	_ "github.com/lib/pq"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestReposStruct_DataBaseInit(t *testing.T) {
	repo := ReposStruct{}
	if err := repo.DataBaseInit(); err != nil {
		t.Fatalf("cannot init db: %s", err)
	}
}

func TestReposStruct_NewDataBaseWorker(t *testing.T) {
	repo := ReposStruct{}
	if err := repo.NewDataBaseWorker(); err != nil {
		t.Fatalf("cannot create db worker: %s", err)
	}
}

func TestReposStruct_WriteData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var email string = "mari@mail.ru"
	var username string = "Mari"
	var password string = "Qw12##!NFkq"

	// good query
	rows := sqlmock.NewRows([]string{"id",})
	expect := []uint64{
		68,
	}

	for _, id := range expect {
		rows = rows.AddRow(id)
	}

	mock.
		ExpectQuery("INSERT INTO").
		WithArgs(username, email, password).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}
	var params []interface{}
	params = append(params, username, email, password)
	id, err := repo.Insert(consts.INSERTRegistration, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if id != strconv.Itoa(int(expect[0])) {
		t.Errorf("result not match, want %v, have %v", expect[0], id)
	}
}


func TestReposStruct_ReadUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var email string = "mari@mail.ru"

	// good query
	rows := sqlmock.NewRows([]string{"id", "username", "name", "surname", "hashpassword", "email", "age",
		"status", "avatardir", "isactive", "salt", "created_time"})
	expect := []*models.User{
		{1, "Mari", "Mari", "Frolova", "Qw12##!NFkq",
			"mari@mail.ru", 32, "I'am okey", "img/p1.png", true, "", time.Now()},
	}

	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Name, user.Surname, user.Password,
			user.Email, user.Age, user.Status, user.AvatarDir, user.IsActive, user.Salt, user.CreatedTime)
	}

	mock.
		ExpectQuery("SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive, U.salt, U.created_time from sunrise.User as U where").
		WithArgs(email).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, email)
	user, err := repo.SelectFullUser(consts.SELECTUserByEmail, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(user, expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], user)
	}
}

func TestReposStruct_ReadUserCookies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var CookieValue string = "FF"

	// good query
	rows := sqlmock.NewRows([]string{"cookiesvalue", "cookiesexpiration"})
	expect := []*models.UserSession{
		{10, 1, models.UserCookie{"FF", time.Now().Add(1 * time.Hour)}}}

	for _, session:= range expect {
		rows = rows.AddRow(session.Value, session.Expiration)
	}

	mock.
		ExpectQuery("SELECT s.cookiesvalue, s.cookiesexpiration from sunrise.usersession" +
		" as s where").
		WithArgs(CookieValue).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, CookieValue)
	cookie, err := repo.SelectUserCookies(consts.SELECTCookiesExpirationByCookieValue, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if cookie[0].Value != expect[0].Value {
		t.Errorf("result not match, want %v, have %v", expect[0].Value, cookie[0].Value)
	}
	if cookie[0].Expiration != expect[0].Expiration {
		t.Errorf("result not match, want %v, have %v", expect[0].Expiration, cookie[0].Expiration)
	}
}

func TestReposStruct_ReadOneCol(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var email string = "my@mail.ru"

	// good query
	rows := sqlmock.NewRows([]string{"id"})
	expect := []uint64{1}

	for _, id := range expect {
		rows = rows.AddRow(id)
	}

	mock.
		ExpectQuery("SELECT u.id from sunrise.user as u where").
		WithArgs(email).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, email)
	id, err := repo.SelectOneCol(consts.SELECTUserIDByEmail, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if id[0] != strconv.Itoa(int(expect[0])) {
		t.Errorf("result not match, want %v, have %v", expect[0], id[0])
	}
}

func TestReposStruct_DeleteSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var sessionKey string = "FF"

	mock.
		ExpectQuery("DELETE FROM sunrise.usersession as s WHERE").
		WithArgs(sessionKey).
		WillReturnRows(nil)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, sessionKey)
	err = repo.DeleteSession(consts.DELETESessionByKey, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
}

func TestReposStruct_DeleteSubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var userID string = "1"
	var followeeName = "Max"

	mock.
		ExpectQuery("DELETE FROM sunrise.subscribe as s WHERE").
		WithArgs(userID, followeeName).
		WillReturnRows(nil)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, userID, followeeName)
	err = repo.DeleteSubscribe(consts.DELETESubscribeByName, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
}

func TestReposStruct_SelectBoard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "owner_id", "title", "description", "category", "createdTime",
		"isDeleted"})
	expect := []*models.Board{
		{1, 1, "MyBoard", "MyDesc", "cars",
			time.Now(), false},
	}

	for _, board := range expect {
		rows = rows.AddRow(board.ID, board.OwnerID, board.Title, board.Description, board.Category, board.CreatedTime,
			board.IsDeleted)
	}

	var boardID string = "1"

	mock.
		ExpectQuery("SELECT b.id, b.owner_id, b.title, b.description, b.category, b.createdTime, b.isDeleted " +
		"FROM sunrise.board as b WHERE").
		WithArgs(boardID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, boardID)
	_, err = repo.SelectBoard(consts.SELECTBoardByID, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(boardID, expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], boardID)
	}
}

func TestReposStruct_SelectCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"name"})
	expect := []string{"cars", "natural"}

	for _, category := range expect {
		rows = rows.AddRow(category)
	}

	mock.
		ExpectQuery("SELECT c.name FROM sunrise.category as c").
		WithArgs().
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params)
	categories, err := repo.SelectCategory("SELECT c.name FROM sunrise.category as c", params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if categories[0] != expect[0] {
		t.Errorf("result not match, want %v, have %v", expect[0], categories[0])
	}
	if categories[1] != expect[1] {
		t.Errorf("result not match, want %v, have %v", expect[1], categories[1])
	}
}

/*
func TestReposStruct_SelectComments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "pin_id", "text", "created_time", "author_id"})
	expect := []*models.Comment{
		{1, 1, "Text", time.Now(), 1},
	}

	for _, comment := range expect {
		rows = rows.AddRow(comment.ID, comment.PinID, comment.Text, comment.CreatedTime, comment.AuthorID)
	}

	var pinID string = "1"

	mock.
		ExpectQuery("SELECT c.text, u.username, c.created_time FROM comment as c").
		WithArgs(pinID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, pinID)
	comments, err := repo.SelectComments(consts.SELECTComments, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(comments[0], expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], comments[0])
	}
}
*/

func TestReposStruct_SelectPin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "owner_id", "author_id", "board_id",
		"title", "description", "pindir", "createdTime", "isDeleted"})
	expect := []*models.Pin{
		{1, 1, 1, 1, "Title","Desc", "/dir/", time.Now(), false},
	}

	for _, pin := range expect {
		rows = rows.AddRow(pin.ID, pin.OwnerID, pin.AuthorID, pin.BoardID, pin.Title, pin.Description, pin.PinDir,
			pin.CreatedTime, pin.IsDeleted)
	}

	var pinID string = "1"

	mock.
		ExpectQuery("SELECT p.id, p.owner_id, p.author_id, p.board_id, p.title," +
		" p.description, p.pindir, p.createdTime, p.isDeleted " +
		"FROM sunrise.pin as p WHERE").
		WithArgs(pinID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, pinID)
	pins, err := repo.SelectPin(consts.SELECTPinByID, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(pins[0], expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], pins[0])
	}
}

func TestReposStruct_SelectPinsByTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "pindir", "title"})
	expect := []*models.PinForSearchResult{
		{1, "/dir/", "title"},
	}

	for _, pin := range expect {
		rows = rows.AddRow(pin.ID, pin.PinDir, pin.Title, )
	}

	var tagName string = "car"

	mock.
		ExpectQuery("SELECT DISTINCT p.id, p.pindir, p.title FROM sunrise.pin as p " +
		"JOIN sunrise.pinandtag as pt ON p.id = pt.pin_id WHERE").
		WithArgs(tagName).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, tagName)
	pins, err := repo.SelectPinsByTag(consts.SELECTPinsByTag, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(pins[0], expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], pins[0])
	}
}

func TestReposStruct_SelectSessions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "userID"})
	expect := []*models.UserSession{
		{1, 1, models.UserCookie{"g", time.Now().Add(24 * time.Hour)}},
	}

	for _, session := range expect {
		rows = rows.AddRow(session.ID, session.UserID)
	}

	var cookieValue string = "val"

	mock.
		ExpectQuery("SELECT s.id, s.userid FROM sunrise.usersession as s WHERE").
		WithArgs(cookieValue).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, cookieValue)
	sessions, err := repo.SelectSessions(consts.SELECTSessionByCookieValue, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if sessions[0].ID != expect[0].ID || sessions[0].UserID != expect[0].UserID {
		t.Errorf("result not match, want %v, have %v", expect[0], sessions[0])
	}
}

func TestReposStruct_SelectIDUsernameEmailUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "username", "email"})
	expect := []*models.UserUnique{
		{1, "vova@mail.com", "Nani"},
	}

	for _, user:= range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Email)
	}

	var username string = "Nani"
	var email = "vova@mail.com"

	mock.
		ExpectQuery("SELECT u.id, u.username, u.email from sunrise.user as u where").
		WithArgs(username, email).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, username, email)
	users, err := repo.SelectIDUsernameEmailUser(consts.SELECTUserIDUsernameEmailByUsernameOrEmail, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if users[0].ID != expect[0].ID || users[0].Email != expect[0].Email || users[0].Username != expect[0].Username {
		t.Errorf("result not match, want %v, have %v", expect[0], users[0])
	}
}