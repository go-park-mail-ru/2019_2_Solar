package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"io/ioutil"
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

	if err := RS.LoadSchemaSQL(); err != nil {
		err, ok := err.(*pq.Error)
		if !ok {
			return err
		}
		if err.Code != pq.ErrorCode("42P06") {
			return err
		}
	}

	return nil
}

func (RS *RepositoryStruct) LoadSchemaSQL() error {

	dbSchema := "sunrise_db.sql"

	content, err := ioutil.ReadFile(dbSchema)
	if err != nil {
		return err
	}
	tx, err := RS.DataBase.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	println(string(content))

	if _, err = tx.Exec(string(content)); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
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
	return strconv.FormatUint(id, 10), nil
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

func (RS *RepositoryStruct) DeleteSubscribe(executeQuery string, params []interface{}) error {
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

func (RS *RepositoryStruct) SelectPin(executeQuery string, params []interface{}) (Pins []models.Pin, Err error) {
	pins := make([]models.Pin, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return pins, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()

	for rows.Next() {
		scanPin := models.Pin{}
		err := rows.Scan(&scanPin.ID, &scanPin.OwnerID, &scanPin.AuthorID, &scanPin.BoardID, &scanPin.Title,
			&scanPin.Description, &scanPin.PinDir, &scanPin.CreatedTime, &scanPin.IsDeleted)
		if err != nil {
			return pins, err
		}
		pins = append(pins, scanPin)
	}
	return pins, nil
}

func (RS *RepositoryStruct) SelectBoard(executeQuery string, params []interface{}) (Board models.Board, Err error) {
	var board models.Board
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return board, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()

	scanBoard := models.Board{}
	for rows.Next() {
		err := rows.Scan(&scanBoard.ID, &scanBoard.OwnerID, &scanBoard.Title,
			&scanBoard.Description, &scanBoard.Category, &scanBoard.CreatedTime, &scanBoard.IsDeleted)
		if err != nil {
			return board, err
		}
		board = scanBoard
	}
	return board, nil
}

func (RS *RepositoryStruct) SelectIdDirPins(executeQuery string, params []interface{}) (Pins []models.PinForMainPage, Err error) {
	var pins []models.PinForMainPage
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return pins, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	scanPin := models.PinForMainPage{}
	for rows.Next() {
		err := rows.Scan(&scanPin.ID, &scanPin.PinDir, &scanPin.IsDeleted)
		if err != nil {
			return pins, err
		}
		pins = append(pins, scanPin)
	}
	return pins, nil
}

func (RS *RepositoryStruct) SelectComments(executeQuery string, params []interface{}) (Comments []models.CommentForSend, Err error) {
	var comments []models.CommentForSend
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return comments, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	scanComment := models.CommentForSend{}
	for rows.Next() {
		err := rows.Scan(&scanComment.Text, &scanComment.Author, &scanComment.CreatedTime)
		if err != nil {
			return comments, err
		}
		comments = append(comments, scanComment)
	}
	return comments, nil
}
