package middlewares

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/labstack/echo"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("  ", ctx.Request().URL.Path)
		cookie, err := ctx.Request().Cookie("session_key")
		if err != nil {
			return next(ctx)
		}

		var User repository.UsersSlice
		err = repository.DBWorker.UniversalRead(consts.QueryReadUserByCookie+"'"+cookie.Value+"'", &User)
		if err != nil || len(User) != 1 {
			return err
		}

		var Cookie repository.UserCookiesSlice
		err = repository.DBWorker.UniversalRead(consts.QueryCookiesExpiration+"'"+cookie.Value+"'", &Cookie)
		if err != nil || len(Cookie) != 1 {
			return err
		}

		ctx.Set("User", User[0])
		ctx.Set("Cookie", Cookie[0])
		return next(ctx)
	}
}
