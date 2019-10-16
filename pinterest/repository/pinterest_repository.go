package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	_ "github.com/lib/pq"
)

var ConnStr string = "user=postgres password=7396 dbname=testdatabase sslmode=disable"

func init() {
	DBWorker.connectionString = ConnStr
}

type DataBaseWorker struct {
	connectionString string
	DataBase         *sql.DB
}

var DBWorker = DataBaseWorker{
	connectionString: "",
	DataBase:         nil,
}

func (dbw *DataBaseWorker) NewDataBaseWorker (){
	dbw.connectionString = ConnStr
	dbw.DataBase = nil
}

func (dbw *DataBaseWorker) WriteData(executeQuery string) error {
	var err error = nil
	if dbw.DataBase == nil {
		dbw.DataBase, err = sql.Open("postgres", ConnStr)
		if err != nil {
			return err
		}
	}
	defer dbw.DataBase.Close()
	result, err := dbw.DataBase.Exec(executeQuery)
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
)

func (US *UsersSlice) DBRead(rows *sql.Rows) error {
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

func (dbw *DataBaseWorker) UniversalRead(executeQuery string, readSlice DBReader) error {
	var err error = nil
	if dbw.DataBase == nil {
		dbw.DataBase, err = sql.Open("postgres", ConnStr)
		if err != nil {
			return err
		}
	}
	defer dbw.DataBase.Close()
	rows, err := dbw.DataBase.Query(executeQuery)
	if err != nil {
		return err
	}
	err = readSlice.DBRead(rows)
	if err != nil {
		return err
	}
	return nil
}
