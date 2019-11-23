package repositoryMiddleware

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

func (MRS *MRepositoryStruct) DataBaseInit() error {
	MRS.connectionString = consts.ConnStr
	var err error

	MRS.DataBase, err = sql.Open("postgres", consts.ConnStr)
	if err != nil {
		return err
	}
	MRS.DataBase.SetMaxOpenConns(10)
	err = MRS.DataBase.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (MRS *MRepositoryStruct) SelectUsersByCookieValue(cookieValue string) (Users []models.User, Err error) {
	usersSlice := make([]models.User, 0)
	rows, err := MRS.DataBase.Query(consts.SELECTUserByCookieValue, cookieValue)
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

func (MRS *MRepositoryStruct) SelectSessionsByCookieValue(cookieValue string) (Sessions []models.UserSession, Err error) {
	userSessionsSlice := make([]models.UserSession, 0)
	rows, err := MRS.DataBase.Query(consts.SELECTSessionByCookieValue, cookieValue)
	if err != nil {
		return userSessionsSlice, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	for rows.Next() {
		userSession := models.UserSession{}
		err := rows.Scan(&userSession.ID, &userSession.UserID, &userSession.Value, &userSession.Expiration)
		if err != nil {
			return userSessionsSlice, err
		}
		userSessionsSlice = append(userSessionsSlice, userSession)
	}
	return userSessionsSlice, nil
}
