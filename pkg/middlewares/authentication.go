package middlewares

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"time"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokens, _ := functions.NewAesCryptHashToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
		cookie, err := ctx.Cookie("session_key")
		if err != nil {
			return next(ctx)
		}
		dbWorker := repository.ReposStruct{}
		defer func() {
			dbWorker.CloseDB()
		}()
		err = dbWorker.NewDataBaseWorker()
		if err != nil {
			return err
		}
		var user []models.User
		user, err = dbWorker.SelectUsersByCookieValue(cookie.Value)
		if err != nil {
			return err
		}
		if len(user) == 0 {
			return errors.New("cookie not found")
		}
		if len(user) > 1 {
			return errors.New("several same cookies")
		}

		var userCookie []models.UserCookie
		userCookie, err = dbWorker.SelectCookiesByCookieValue(cookie.Value)
		if err != nil {
			return err
		}
		if len(user) == 0 {
			return errors.New("cookie not found")
		}
		if len(user) > 1 {
			return errors.New("several same cookies")
		}

		if userCookie[0].Expiration.Before(time.Now()) {
			//delete Coockie!!!!
			return next(ctx)
		}

		var userSessions []models.UserSession
		userSessions, err = dbWorker.SelectSessionsByCookieValue(cookie.Value)
		if err != nil {
			return err
		}
		sess := functions.Session{
			UserID: uint(userSessions[0].UserID),
			ID:     string(userSessions[0].ID),
		}

		if ctx.Request().URL.Path != "/login" &&
			ctx.Request().URL.Path != "/registration" &&
			ctx.Request().Method != "GET" {

			CSRFToken := ctx.Request().Header.Get("csrf-token")
			_, err = tokens.Check(&sess, CSRFToken)
			if err != nil {
				return err
			}
		}

		token, err := tokens.Create(&sess, time.Now().Add(30*time.Minute).Unix())
		if err != nil {
			return err
		}

		ctx.Set("token", token)
		ctx.Set("User", user[0])
		ctx.Set("Cookie", userCookie[0])
		return next(ctx)
	}
}
