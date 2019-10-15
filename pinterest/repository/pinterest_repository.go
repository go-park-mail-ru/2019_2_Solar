package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	_ "github.com/lib/pq"
)

var connStr string = "user=postgres password=mypass dbname=productdb sslmode=disable"

func init() {
	DBWorker.connectionString = connStr
}

type DataBaseWorker struct {
	connectionString string
	dataBase         *sql.DB
}

var DBWorker = DataBaseWorker{
	connectionString: "",
	dataBase:         nil,
}

func (dbw *DataBaseWorker) WriteData(executeQuery string) error {
	var err error = nil
	if dbw.dataBase == nil {
		dbw.dataBase, err = sql.Open("postgres", connStr)
		if err != nil {
			return err
		}
	}
	defer dbw.dataBase.Close()
	result, err := dbw.dataBase.Exec(executeQuery)
	if err != nil {
		return err
	}
	fmt.Println(result.LastInsertId())
	return nil
}

func (dbw *DataBaseWorker) ReadData(executeQuery string, userSlice *[]models.User) error {
	var err error = nil
	if dbw.dataBase == nil {
		dbw.dataBase, err = sql.Open("postgres", connStr)
		if err != nil {
			return err
		}
	}
	defer dbw.dataBase.Close()
	rows, err := dbw.dataBase.Query(executeQuery)
	if err != nil {
		return err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.AvatarDir, &user.IsActive, &user.Name,
			&user.Password, &user.Status, &user.Surname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*userSlice = append(*userSlice, user)
	}
	return nil
}
