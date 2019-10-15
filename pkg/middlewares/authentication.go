package middlewares

import (
	"fmt"
	"github.com/labstack/echo"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("  ", ctx.Request().URL.Path)
		_, err := ctx.Request().Cookie("session_id")

		// учебный пример! это не проверка авторизации!
		if err != nil {
			fmt.Println("no auth at", ctx.Request().URL.Path)
			//http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}
		return next(ctx)
	}
}
