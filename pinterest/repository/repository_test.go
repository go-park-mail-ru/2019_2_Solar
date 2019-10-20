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

func TestRepositoryStruct_WriteData(t *testing.T) {
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

	repo := &RepositoryStruct{
		connectionString: ConnStr,
		DataBase:         db,
	}
	var params []interface{}
	params = append(params, username, email, password)
	id, err := repo.WriteData(consts.InsertRegistrationQuery, params)

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


func TestRepositoryStruct_ReadUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var email string = "mari@mail.ru"

	// good query
	rows := sqlmock.NewRows([]string{"id", "username", "name", "surname", "hashpassword", "email", "age",
		"status", "avatardir", "isactive"})
	expect := []*models.User{
		{1, "Mari", "Mari", "Frolova", "Qw12##!NFkq",
			"mari@mail.ru", 32, "I'am okey", "img/p1.png", true},
	}

	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Name, user.Surname, user.Password,
			user.Email, user.Age, user.Status, user.AvatarDir, user.IsActive)
	}

	mock.
		ExpectQuery("SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive from sunrise.Users as U where").
		WithArgs(email).
		WillReturnRows(rows)

	repo := &RepositoryStruct{
		connectionString: ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, email)
	user, err := repo.ReadUser(consts.ReadUserByEmailSQLQuery, params)

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

func TestRepositoryStruct_ReadUserCookies(t *testing.T) {
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
		ExpectQuery("SELECT s.cookiesvalue, s.cookiesexpiration from sunrise.usersessions" +
	" as s where").
		WithArgs(CookieValue).
		WillReturnRows(rows)

	repo := &RepositoryStruct{
		connectionString: ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, CookieValue)
	cookie, err := repo.ReadUserCookies(consts.ReadCookiesExpirationByCookieValueSQLQuery, params)

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

func TestRepositoryStruct_ReadOneCol(t *testing.T) {
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
		ExpectQuery("SELECT u.id from sunrise.users as u where").
		WithArgs(email).
		WillReturnRows(rows)

	repo := &RepositoryStruct{
		connectionString: ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, email)
	id, err := repo.ReadOneCol(consts.ReadUserIdByEmailSQLQuery, params)

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

func TestRepositoryStruct_DeleteSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cannot create mock: %s", err)
	}
	defer db.Close()

	var sessionKey string = "FF"

	mock.
		ExpectQuery("DELETE FROM sunrise.usersessions as s WHERE").
		WithArgs(sessionKey).
		WillReturnRows(nil)

	repo := &RepositoryStruct{
		connectionString: ConnStr,
		DataBase:         db,
	}

	var params []interface{}
	params = append(params, sessionKey)
	err = repo.DeleteSession(consts.DeleteSessionByKey, params)

	if err != nil {
		t.Errorf("inexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infilfilled expectations: %s", err)
		return
	}

}
