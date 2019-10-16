package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	_ "github.com/lib/pq"
)

var ConnStr string = "user=postgres password=7396 dbname=testdatabase sslmode=disable"

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

func (RS *RepositoryStruct) WriteData(executeQuery string, params []interface{}) error {
	result, err := RS.DataBase.Exec(executeQuery, params...)
	if err != nil {
		return err
	}
	fmt.Println(result.LastInsertId())
	return nil
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
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Age,
			&user.Status, &user.AvatarDir, &user.IsActive)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*US = append(*US, user)
	}
	return nil
}

func (USC *UserCookiesSlice) DBRead(rows *sql.Rows) error {
	defer rows.Close()
	for rows.Next() {
		userCookie := models.UserCookie{}
		err := rows.Scan(&userCookie.Value, &userCookie.Expiration)
		if err != nil {
			fmt.Println(err)
			continue
		}
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

func (RS *RepositoryStruct) UniversalRead(executeQuery string, readSlice DBReader, params []interface{}) error {
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
