package middlewares

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"time"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("session_key")
		if err != nil {
			return next(ctx)
		}
		dbWorker := repository.RepositoryStruct{}
		err = dbWorker.NewDataBaseWorker()
		if err != nil {
			return err
		}
		var user []models.User
		var params []interface{}
		params = append(params, cookie.Value)
		user, err = dbWorker.ReadUser(consts.ReadUserByCookieValueSQLQuery, params)
		if err != nil || len(user) != 1 {
			return err
		}

		var userCookie []models.UserCookie
		userCookie, err = dbWorker.ReadUserCookies(consts.ReadCookiesExpirationByCookieValueSQLQuery, params)
		if err != nil || len(userCookie) != 1 {
			return err
		}
		if userCookie[0].Expiration.Before(time.Now()) {
			//delete Coockie!!!!
			return next(ctx)
		}
		ctx.Set("User", user[0])
		ctx.Set("Cookie", userCookie[0])
		return next(ctx)
	}
}
