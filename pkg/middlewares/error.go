package middlewares

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("panicMiddleware", ctx.Request().URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(ctx.Response(), "Internal server error", 500)
			}
		}()
		return next(ctx)
	}
}
