package middlewares

import (
	//"github.com/labstack/echo"
	echov4 "github.com/labstack/echo/v4"
)

func (MS *MiddlewareStruct) CORSMiddleware(next echov4.HandlerFunc) echov4.HandlerFunc {
	return func(ctx echov4.Context) error {
		ctx.Response().Header().Set("Content-Type", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, csrf-token, Authorization")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request().Method == "OPTIONS" {
			return nil
		}
		return next(ctx)
	}
}
