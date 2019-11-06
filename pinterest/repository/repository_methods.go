package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
)


func (RS *ReposStruct) DataBaseInit() error {
	RS.connectionString = consts.ConnStr
	var err error

	RS.DataBase, err = sql.Open("postgres", consts.ConnStr)
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

func (RS *ReposStruct) LoadSchemaSQL() (Err error) {
	dbSchema := "sunrise_db.sql"

	content, err := ioutil.ReadFile(dbSchema)
	if err != nil {
		return err
	}
	tx, err := RS.DataBase.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			Err = errors.Wrap(Err, err.Error())
		}
	}()

	if _, err = tx.Exec(string(content)); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (RS *ReposStruct) NewDataBaseWorker() error {
	RS.connectionString = consts.ConnStr
	var err error

	RS.DataBase, err = sql.Open("postgres", consts.ConnStr)
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

func (RS *ReposStruct) Insert(executeQuery string, params []interface{}) (string, error) {
	var id uint64
	err := RS.DataBase.QueryRow(executeQuery, params...).Scan(&id)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}

func (RS *ReposStruct) Update(executeQuery string, params []interface{}) (int, error) {
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

func (RS *ReposStruct) SelectFullUser(executeQuery string, params []interface{}) (Sl []models.User, Err error) {
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
			&dbuser.Status, &dbuser.AvatarDir, &dbuser.IsActive, &dbuser.Salt, &dbuser.CreatedTime)
		if err != nil {
			return usersSlice, err
		}
		user := models.User{
			ID:          dbuser.ID,
			Username:    dbuser.Username,
			Name:        dbuser.Name.String,
			Surname:     dbuser.Surname.String,
			Password:    dbuser.Password,
			Email:       dbuser.Email,
			Age:         uint(dbuser.Age.Int32),
			Status:      dbuser.Status.String,
			AvatarDir:   dbuser.AvatarDir.String,
			IsActive:    dbuser.IsActive,
			Salt:        dbuser.Salt,
			CreatedTime: dbuser.CreatedTime,
		}
		usersSlice = append(usersSlice, user)
	}
	return usersSlice, nil
}

func (RS *ReposStruct) SelectIDUsernameEmailUser(executeQuery string, params []interface{}) (Sl []models.UserUnique, Err error) {
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
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return userUniqueSlice, err
		}
		userUniqueSlice = append(userUniqueSlice, user)
	}
	return userUniqueSlice, nil
}

func (RS *ReposStruct) SelectUserCookies(executeQuery string, params []interface{}) (Sl []models.UserCookie, Err error) {
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

func (RS *ReposStruct) SelectOneCol(executeQuery string, params []interface{}) (Sl []string, Err error) {
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

func (RS *ReposStruct) DeleteSession(executeQuery string, params []interface{}) error {
	_, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return err
	}
	return nil
}

func (RS *ReposStruct) DeleteSubscribe(executeQuery string, params []interface{}) error {
	_, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return err
	}
	return nil
}

func (RS *ReposStruct) SelectCategory(executeQuery string, params []interface{}) (categories []string, Err error) {
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

func (RS *ReposStruct) SelectPin(executeQuery string, params []interface{}) (Pins []models.Pin, Err error) {
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

func (RS *ReposStruct) SelectPinsByTag(executeQuery string, params []interface{}) (Pins []models.PinForSearchResult, Err error) {
	pins := make([]models.PinForSearchResult, 0)
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
		scanPin := models.PinForSearchResult{}
		err := rows.Scan(&scanPin.ID, &scanPin.PinDir, &scanPin.Title)
		if err != nil {
			return pins, err
		}
		pins = append(pins, scanPin)
	}
	return pins, nil
}

func (RS *ReposStruct) SelectBoard(executeQuery string, params []interface{}) (Board models.Board, Err error) {
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

func (RS *ReposStruct) SelectIDDirPins(executeQuery string, params []interface{}) (Pins []models.PinForMainPage, Err error) {
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
		err := rows.Scan(&scanPin.ID, &scanPin.PinDir, &scanPin.Title)
		if err != nil {
			return pins, err
		}
		pins = append(pins, scanPin)
	}
	return pins, nil
}

func (RS *ReposStruct) SelectComments(executeQuery string, params []interface{}) (Comments []models.CommentForSend, Err error) {
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

func (RS *ReposStruct) SelectSessions(executeQuery string, params []interface{}) (Sessions []models.UserSession, Err error) {
	var sessions []models.UserSession
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return sessions, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	scanSession := models.UserSession{}
	for rows.Next() {
		err := rows.Scan(&scanSession.ID, &scanSession.UserID)
		if err != nil {
			return sessions, err
		}
		sessions = append(sessions, scanSession)
	}
	return sessions, nil
}

func (RS *ReposStruct) SelectBoards(executeQuery string, params []interface{}) (Boards []models.Board, Err error) {
	var boards []models.Board
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return boards, err
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
			return boards, err
		}
		boards = append(boards, scanBoard)
	}
	return boards, nil
}