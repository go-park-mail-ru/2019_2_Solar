package middlewares

import "github.com/labstack/echo"

type MiddlewareStruct struct {
}

type MiddlewareInterface interface {
	NewMiddleware(e *echo.Echo)
	AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CustomHTTPErrorHandler(err error, ctx echo.Context)
	PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}
