package repository

import (
	"database/sql"
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

func (RS *RepositoryStruct) Insert(executeQuery string, params []interface{}) (string, error) {
	var id uint64
	err := RS.DataBase.QueryRow(executeQuery, params...).Scan(&id)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}

func (RS *RepositoryStruct) InsertBoard(executeQuery string, params []interface{}) (string, error) {
	var id uint64
	err := RS.DataBase.QueryRow(executeQuery, params...).Scan(&id)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}

func (RS *RepositoryStruct) InsertPin(executeQuery string, params []interface{}) (string, error) {
	var id uint64
	err := RS.DataBase.QueryRow(executeQuery, params...).Scan(&id)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}

func (RS *RepositoryStruct) Update(executeQuery string, params []interface{}) (int, error) {
	result, err := RS.DataBase.Exec(executeQuery, params...)
	if err != nil {
		return 0, err
	}
	rowsEdit, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsEdit), nil
}

func (RS *RepositoryStruct) SelectFullUser(executeQuery string, params []interface{}) (Sl []models.User, Err error) {
	usersSlice := make([]models.User, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return usersSlice, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	for rows.Next() {
		dbuser := models.DBUser{}
		err := rows.Scan(&dbuser.ID, &dbuser.Username, &dbuser.Name, &dbuser.Surname, &dbuser.Password, &dbuser.Email, &dbuser.Age,
			&dbuser.Status, &dbuser.AvatarDir, &dbuser.IsActive)
		if err != nil {
			return usersSlice, err
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
		usersSlice = append(usersSlice, user)
	}
	return usersSlice, nil
}

func (RS *RepositoryStruct) SelectIdUsernameEmailUser(executeQuery string, params []interface{}) (Sl []models.UserUnique, Err error) {
	userUniqueSlice := make([]models.UserUnique, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return userUniqueSlice, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	for rows.Next() {
		user := models.UserUnique{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			return userUniqueSlice, err
		}
		userUniqueSlice = append(userUniqueSlice, user)
	}
	return userUniqueSlice, nil
}

func (RS *RepositoryStruct) SelectUserCookies(executeQuery string, params []interface{}) (Sl []models.UserCookie, Err error) {
	userCookiesSlice := make([]models.UserCookie, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return userCookiesSlice, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	for rows.Next() {
		userCookie := models.UserCookie{}
		err := rows.Scan(&userCookie.Value, &userCookie.Expiration)
		if err != nil {
			return userCookiesSlice, err
		}
		userCookiesSlice = append(userCookiesSlice, userCookie)
	}
	return userCookiesSlice, nil
}

func (RS *RepositoryStruct) SelectOneCol(executeQuery string, params []interface{}) (Sl []string, Err error) {
	stringSlice := make([]string, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return stringSlice, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	for rows.Next() {
		var str string
		err := rows.Scan(&str)
		if err != nil {
			return stringSlice, err
		}
		stringSlice = append(stringSlice, str)
	}
	return stringSlice, nil
}

func (RS *RepositoryStruct) DeleteSession(executeQuery string, params []interface{}) error {
	_, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return err
	}
	return nil
}

func (RS *RepositoryStruct) SelectCategory(executeQuery string, params []interface{}) (categories []string, Err error) {
	categories = make([]string, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return categories, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	var category *string
	for rows.Next() {
		err := rows.Scan(&category)
		if err != nil {
			return categories, err
		}

		categories = append(categories, *category)
	}
	return categories, nil
}
