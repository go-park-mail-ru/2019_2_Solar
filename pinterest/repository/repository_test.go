package repository

import (
	"database/sql"
	_ "database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	_ "github.com/lib/pq"
	"reflect"
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

func TestSelectFullUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var username string = "Mari"

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
		WithArgs(username).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	user, err := repo.SelectUsersByUsername(username)

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
	cookie, err := repo.SelectCookiesByCookieValue(CookieValue)

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




func TestDeleteSession(t *testing.T) {
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
	err = repo.DeleteSessionByKey(sessionKey)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
}

func TestDeleteSubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var userID = 1
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
	err = repo.DeleteSubscribeByName(uint64(userID), followeeName)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
}

func TestSelectBoard(t *testing.T) {
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

	var boardID uint64 = 1

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
	_, err = repo.SelectBoardsByID(boardID)

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

func TestSelectCategory(t *testing.T) {
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
		ExpectQuery("SELECT c.name FROM sunrise.category as c;").
		WithArgs().
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params)
	categories, err := repo.SelectCategories()

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if len(categories) != 2 {
		t.Errorf("result not match, want 2 categoies")
	}
}


func TestReposStruct_SelectComments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"text", "username", "avatardir", "created_time"})
	expect := []*models.CommentDisplay{
		{"vooooo",time.Now(), "Name", "/dir"},
	}

	for _, comment := range expect {
		rows = rows.AddRow(comment.Text, comment.Author, comment.AuthorPicture, comment.CreatedTime)
	}

	var pinID uint64 = 1

	mock.
		ExpectQuery("SELECT c.text, u.username, u.avatardir, c.created_time FROM sunrise.comment").
		WithArgs(pinID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	comments, err := repo.SelectCommentsByPinId(pinID)

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


func TestSelectPin(t *testing.T) {
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

	var pinID uint64 = 1

	mock.
		ExpectQuery("SELECT p.id, o.username, a.username, p.board_id, p.title, p.description, p.pindir, p.createdTime, p.isDeleted ").
		WithArgs(pinID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, pinID)
	pins, err := repo.SelectPinsById(pinID)

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

func TestSelectPinsByTag(t *testing.T) {
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
	pins, err := repo.SelectPinsByTag(tagName)

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
		ExpectQuery("SELECT s.id, s.userid, s.cookiesvalue, s.cookiesexpiration FROM sunrise.usersession as s").
		WithArgs(cookieValue).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, cookieValue)
	sessions, err := repo.SelectSessionsByCookieValue(cookieValue)

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

func TestSelectIDUsernameEmailUser(t *testing.T) {
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
	users, err := repo.SelectIDUsernameEmailUser(username, email)

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

func TestSelectUsersByCookieValue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "username", "name", "surname", "hashpassword", "email", "age", "status",
		"avatardir", "isactive", "salt", "created_time"})
	expect := []*models.DBUser{
		{1, "Vo1",  sql.NullString{"Vov", true}, sql.NullString{"Voi", true}, "123", "emil1@com.er",
			sql.NullInt32{32, true}, sql.NullString{"", true}, sql.NullString{"", true},
			true, "500", time.Now()},
	}

	for _, user:= range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Name, user.Surname, user.Password, user.Email, user.Age, user.Status,
			user.AvatarDir, user.IsActive, user.Salt, user.CreatedTime)
	}

	var cookieValue string = "FF"

	mock.
		ExpectQuery("SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status,").
		WithArgs(cookieValue).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	users, err := repo.SelectUsersByCookieValue(cookieValue)

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

func TestSelectUsersByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "username", "name", "surname", "hashpassword", "email", "age", "status",
		"avatardir", "isactive", "salt", "created_time"})
	expect := []*models.DBUser{
		{1, "Vo1",  sql.NullString{"Vov", true}, sql.NullString{"Voi", true}, "123", "emil1@com.er",
			sql.NullInt32{32, true}, sql.NullString{"", true}, sql.NullString{"", true},
			true, "500", time.Now()},
	}

	for _, user:= range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Name, user.Surname, user.Password, user.Email, user.Age, user.Status,
			user.AvatarDir, user.IsActive, user.Salt, user.CreatedTime)
	}

	var email string = expect[0].Email

	mock.
		ExpectQuery("SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status").
		WithArgs(email).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	users, err := repo.SelectUsersByEmail(email)

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

func TestSelectCategoryByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"name"})
	expect := []string{"cars"}

	for _, category := range expect {
		rows = rows.AddRow(category)
	}

	var categoryName string = "cars"

	mock.
		ExpectQuery("SELECT c.name FROM sunrise.category as c WHERE").
		WithArgs().
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}
	categories, err := repo.SelectCategoryByName(categoryName)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if len(categories) != 1 {
		t.Errorf("result not match, want 1 categoies")
	}
}


func TestSelectPinsDisplayByBoardId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "title", "pindir"})
	expect := []*models.PinDisplay{
		{uint64(1), "/dir", "hello"},
	}

	for _, pin := range expect {
		rows = rows.AddRow(pin.ID, pin.Title, pin.PinDir)
	}

	var boardID uint64 = 1

	mock.
		ExpectQuery("SELECT p.id, p.title, p.pindir FROM sunrise.pin as p WHERE").
		WithArgs(boardID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	pins, err := repo.SelectPinsDisplayByBoardId(boardID)

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



func TestSelectPinsDisplayByUsernam(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "title", "pindir"})
	expect := []*models.PinDisplay{
		{uint64(1), "/dir", "hello"},
	}

	for _, pin := range expect {
		rows = rows.AddRow(pin.ID, pin.Title, pin.PinDir)
	}

	var userID int = 1

	mock.
		ExpectQuery("SELECT p.id, p.title, p.pindir FROM sunrise.pin as p WHERE").
		WithArgs(userID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	pins, err := repo.SelectPinsDisplayByUsername(userID)

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


func TestSelectBoardsByOwnerId(t *testing.T) {
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

	var ownerID uint64 = 1

	mock.
		ExpectQuery("SELECT b.id, b.owner_id, b.title, b.description, b.category, b.createdTime, b.isDeleted " +
			"FROM sunrise.board as b WHERE").
		WithArgs(ownerID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	boards, err := repo.SelectBoardsByOwnerId(ownerID)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if len(boards) != 1 {
		t.Errorf("result not match, want 1 board")
	}
}


func TestSSelectAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()
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
		ExpectQuery("SELECT *").
		WithArgs().
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	users, err := repo.SelectAllUsers()

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(users[0], expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], users[0])
	}
}


func TestSelectNewPinsDisplayByNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "title", "pindir"})
	expect := []*models.PinDisplay{
		{uint64(1), "/dir", "hello"},
	}

	for _, pin := range expect {
		rows = rows.AddRow(pin.ID, pin.Title, pin.PinDir)
	}

	var first int = 0;
	var last int = 9

	mock.
		ExpectQuery("SELECT p.id, p.pindir, p.title FROM").
		WithArgs(first, last).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	pins, err := repo.SelectNewPinsDisplayByNumber(first, last)

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



func TestSelectMyPinsDisplayByNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "title", "pindir"})
	expect := []*models.PinDisplay{
		{uint64(1), "/dir", "hello"},
	}

	for _, pin := range expect {
		rows = rows.AddRow(pin.ID, pin.Title, pin.PinDir)
	}

	var userID uint64 = 1
	var number int = 20

	mock.
		ExpectQuery("SELECT p.id, p.pindir, p.title FROM").
		WithArgs(number, userID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	pins, err := repo.SelectMyPinsDisplayByNumber(userID, number)

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



func TestSelectSubscribePinsDisplayByNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	// good query
	rows := sqlmock.NewRows([]string{"id", "title", "pindir"})
	expect := []*models.PinDisplay{
		{uint64(1), "/dir", "hello"},
	}

	for _, pin := range expect {
		rows = rows.AddRow(pin.ID, pin.Title, pin.PinDir)
	}

	var userID uint64 = 1
	var first int = 0;
	var last int = 9

	mock.
		ExpectQuery("SELECT p.id, p.pindir, p.title FROM").
		WithArgs(first, last, userID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	pins, err := repo.SelectSubscribePinsDisplayByNumber(userID, first, last)

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



func TestSelectUsersByUsernameSearch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()
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

	username := expect[0].Username

	mock.
		ExpectQuery("SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status").
		WithArgs("%" + username + "%").
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	users, err := repo.SelectUsersByUsernameSearch(username)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(users[0], expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], users[0])
	}
}


func TestSelectMySubscribeByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()
	// good query
	rows := sqlmock.NewRows([]string{"id", "subscriber_id", "followee_id"})
	expect := []*models.Subscribe{
		{1, 2, 1},
	}

	for _, sub := range expect {
		rows = rows.AddRow(sub.Id, sub.IdSubscriber, sub.FolloweeId)
	}

	var userID uint64 = 1
	var followeeName string = "Add"

	mock.
		ExpectQuery("SELECT s.id, s.subscriber_id, s.followee_id FROM sunrise.subscribe as s").
		WithArgs(userID, followeeName).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	sub, err := repo.SelectMySubscribeByUsername(userID, followeeName)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if reflect.DeepEqual(sub, expect[0]) {
		t.Errorf("result not match, want %v, have %v", expect[0], sub)
	}
}


func TestSelectAllTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()
	// good query
	rows := sqlmock.NewRows([]string{"name"})
	expect := []string{"car", "forest"}

	for _, tag := range expect {
		rows = rows.AddRow(tag)
	}

	mock.
		ExpectQuery("SELECT t.name from sunrise.tag as t").
		WithArgs().
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	tags, err := repo.SelectAllTags()

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if tags[0] != expect[0] {
		t.Errorf("result not match, want %v, have %v", expect[0], tags[0])
	}
	if tags[1] != expect[1] {
		t.Errorf("result not match, want %v, have %v", expect[1], tags[1])
	}
}


func TestSelectNoticesByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()
	// good query
	rows := sqlmock.NewRows([]string{"id", "user_id", "receiver_id", "message", "createdtime", "isread"})
	expect := []*models.Notice{
		{1, 1, 1, "Text", time.Now(), false},
	}

	for _, notice := range expect {
		rows = rows.AddRow(notice.ID, notice.UserID, notice.ReceiverID, notice.Message, notice.CreatedTime, notice.IsRead)
	}

	var userID uint64 = 1

	mock.
		ExpectQuery("SELECT n.id, n.user_id, n.receiver_id, n.message, n.createdtime, isread FROM").
		WithArgs(userID).
		WillReturnRows(rows)

	repo := &ReposStruct{
		connectionString: consts.ConnStr,
		DataBase:         db,
	}

	notices, err := repo.SelectNoticesByUserID(userID)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}
	if notices[0] != *expect[0] {
		t.Errorf("result not match, want %v, have %v", expect[0], notices[0])
	}
}