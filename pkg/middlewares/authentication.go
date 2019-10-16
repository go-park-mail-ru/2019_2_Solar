package middlewares

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/labstack/echo"
)

var QueryReadUserByCookie string = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
	" U.avatardir, U.isactive from testschema.Users as U JOIN testschema.usersessions as s on U.id = s.userid " +
	"where s.cookiesvalue = "

var QueryCookiesExpiration string = "SELECT s.cookiesvalue, s.cookiesexpiration from testschema.usersessions" +
	" as s where s.cookiesvalue = "

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("  ", ctx.Request().URL.Path)
		cookie, err := ctx.Request().Cookie("session_key")
		if err != nil {
			return next(ctx)
		}
		fmt.Println(cookie.Value)
		var User repository.UsersSlice
		//var User []models.User
		err = repository.DBWorker.UniversalRead(QueryReadUserByCookie+"'"+cookie.Value+"'", &User)
		if err != nil {
			return next(ctx)
		}
		if len(User) != 1 {
			return next(ctx)
		}
		var Cookie repository.UserCookiesSlice
		//var Cookie []models.UserCookie
		err = repository.DBWorker.UniversalRead(QueryCookiesExpiration+"'"+cookie.Value+"'", &Cookie)
		if err != nil {
			return next(ctx)
		}
		if len(Cookie) != 1 {
			return next(ctx)
		}
		fmt.Println(User[0])
		fmt.Println(Cookie[0])
		ctx.Set("User", User[0])
		ctx.Set("Cookie", Cookie[0])
		return next(ctx)
	}
}
