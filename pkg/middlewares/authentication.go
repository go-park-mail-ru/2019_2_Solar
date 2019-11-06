package middlewares

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
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
		err = dbWorker.NewDataBaseWorker()
		if err != nil {
			return err
		}
		var user []models.User
		var params []interface{}
		params = append(params, cookie.Value)
		user, err = dbWorker.SelectFullUser(consts.SELECTUserByCookieValue, params)
		if err != nil || len(user) != 1 {
			return err
		}

		var userCookie []models.UserCookie
		userCookie, err = dbWorker.SelectUserCookies(consts.SELECTCookiesExpirationByCookieValue, params)
		if err != nil || len(userCookie) != 1 {
			return err
		}
		if userCookie[0].Expiration.Before(time.Now()) {
			//delete Coockie!!!!
			return next(ctx)
		}

		var userSessions []models.UserSession
		userSessions, err = dbWorker.SelectSessions(consts.SELECTSessionByCookieValue, params)

		sess := functions.Session{
			UserID: uint(userSessions[0].UserID),
			ID: string(userSessions[0].ID),
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

		token, err := tokens.Create(&sess, time.Now().Add(24*time.Hour).Unix())
		if err != nil {
			return err
		}

		ctx.Set("token", token)
		ctx.Set("User", user[0])
		ctx.Set("Cookie", userCookie[0])
		return next(ctx)
	}
}
