package middlewares

import (
	"github.com/labstack/echo"
)

func CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error  {
		ctx.Response().Header().Set("Content-Type", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "http://solar-env.v2zxh2s3me.us-east-2.elasticbeanstalk.com")
		ctx.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		return next(ctx)
	}
}
