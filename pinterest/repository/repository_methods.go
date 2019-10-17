package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	_ "github.com/lib/pq"
	"strconv"
)

var ConnStr string = "user=postgres password=7396 dbname=sunrise_db sslmode=disable"

func (RS *RepositoryStruct) NewDataBaseWorker() error {
	RS.connectionString = ConnStr
	var err error = nil

	RS.DataBase, err = sql.Open("postgres", ConnStr)
	if err != nil {
		return err
	}
	RS.DataBase.SetMaxOpenConns(10)
	err = RS.DataBase.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (RS RepositoryStruct) WriteData(executeQuery string, params []interface{}) (string, error) {
	var id uint64
	err := RS.DataBase.QueryRow(executeQuery, params...).Scan(&id)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}

type DBReader interface {
	DBRead(rows *sql.Rows) error
}

type (
	UsersSlice       []models.User
	UserCookiesSlice []models.UserCookie
	StringSlice      []string
)

func (US *UsersSlice) DBRead(rows *sql.Rows) error {
	defer rows.Close()
	for rows.Next() {
		dbuser := models.DBUser{}
		err := rows.Scan(&dbuser.ID, &dbuser.Username, &dbuser.Name, &dbuser.Surname, &dbuser.Password, &dbuser.Email, &dbuser.Age,
			&dbuser.Status, &dbuser.AvatarDir, &dbuser.IsActive)
		if err != nil {
			fmt.Println(err)
			continue
		}
		user := models.User{
			ID:        dbuser.ID,
			Username:  dbuser.Username,
			Name:      dbuser.Name.String,
			Surname:   dbuser.Surname.String,
			Password:  dbuser.Password,
			Email:     dbuser.Email,
			Age:       uint(dbuser.Age.Int32),
			Status:    dbuser.Status.String,
			AvatarDir: dbuser.AvatarDir.String,
			IsActive:  dbuser.IsActive,
		}

		*US = append(*US, user)
	}
	return nil
}

func (USC *UserCookiesSlice) DBRead(rows *sql.Rows) error {
	defer rows.Close()
	for rows.Next() {
		userCookie := models.UserCookie{}
		//var expirationString string
		err := rows.Scan(&userCookie.Value, &userCookie.Expiration)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//userCookie.Expiration = time.Parse(expirationString)//expirationString
		*USC = append(*USC, userCookie)
	}
	return nil
}

func (SS *StringSlice) DBRead(rows *sql.Rows) error {
	defer rows.Close()
	for rows.Next() {
		var str string
		err := rows.Scan(&str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*SS = append(*SS, str)
	}
	return nil
}

func (RS *RepositoryStruct) DBDataRead(executeQuery string, readSlice DBReader, params []interface{}) error {
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return err
	}
	err = readSlice.DBRead(rows)
	if err != nil {
		return err
	}
	return nil
}
