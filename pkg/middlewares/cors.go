package middlewares

import (
	"github.com/labstack/echo"
)

func (MS *MiddlewareStruct) CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set("Content-Type", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, csrf-token, Authorization")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "http://solarsunrise.ru")
		ctx.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request().Method == "OPTIONS" {
			return nil
		}
		return next(ctx)
	}
}
