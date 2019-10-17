package middlewares

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/labstack/echo"
	"time"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("session_key")
		if err != nil {
			return next(ctx)
		}
		DBWorker := repository.RepositoryStruct{}
		err = DBWorker.NewDataBaseWorker()
		if err != nil {
			return err
		}
		var User repository.UsersSlice
		var params []interface{}
		params = append(params, cookie.Value)
		err = DBWorker.DBDataRead(consts.ReadUserByCookieValueSQLQuery, &User, params)
		if err != nil || len(User) != 1 {
			return err
		}

		var userCookie repository.UserCookiesSlice
		err = DBWorker.DBDataRead(consts.ReadCookiesExpirationByCookieValueSQLQuery, &userCookie, params)
		if err != nil || len(userCookie) != 1 {
			return err
		}
		if userCookie[0].Expiration.Before(time.Now()) {
			//delete Coockie!!!!
			return next(ctx)
		}
		ctx.Set("User", User[0])
		ctx.Set("Cookie", userCookie[0])
		return next(ctx)
	}
}
