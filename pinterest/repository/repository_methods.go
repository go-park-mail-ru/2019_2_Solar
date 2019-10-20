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

func (RS *RepositoryStruct) DBWriteData(executeQuery string, params []interface{}) (string, error) {
	var id uint64
	err := RS.DataBase.QueryRow(executeQuery, params...).Scan(&id)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}

func (RS *RepositoryStruct) DBReadDataUser(executeQuery string, params []interface{}) ([]models.User, error) {
	usersSlice := make([]models.User, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return usersSlice, err
	}
	defer rows.Close()
	for rows.Next() {
		dbuser := models.DBUser{}
		err := rows.Scan(&dbuser.ID, &dbuser.Username, &dbuser.Name, &dbuser.Surname, &dbuser.Password, &dbuser.Email, &dbuser.Age,
			&dbuser.Status, &dbuser.AvatarDir, &dbuser.IsActive)
		if err != nil {
			fmt.Println(err)
			continue
		}
		user := models.User {
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
		usersSlice = append(usersSlice, user)
	}
	return usersSlice, nil
}

func (RS *RepositoryStruct) DBReadDataUserCookies(executeQuery string, params []interface{}) ([]models.UserCookie, error) {
	userCookiesSlice := make([]models.UserCookie, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return userCookiesSlice, err
	}
	defer rows.Close()
	for rows.Next() {
		userCookie := models.UserCookie{}
		err := rows.Scan(&userCookie.Value, &userCookie.Expiration)
		if err != nil {
			fmt.Println(err)
			continue
		}
		userCookiesSlice = append(userCookiesSlice, userCookie)
	}
	return userCookiesSlice, nil
}

func (RS *RepositoryStruct) DBReadDataString(executeQuery string, params []interface{}) ([]string, error) {
	stringSlice := make([]string, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return stringSlice, err
	}
	defer rows.Close()
	for rows.Next() {
		var str string
		err := rows.Scan(&str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		stringSlice = append(stringSlice, str)
	}
	return stringSlice, nil
}

func (RS *RepositoryStruct) DELETE_SESSION(executeQuery string, params []interface{}) error {
	_, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return err
	}
	return nil
}