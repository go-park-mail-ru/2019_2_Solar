package middlewares

import (
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/usecase/middleware"
)

type MiddlewareStruct struct {
	MUsecase useCaseMiddleware.MUseCaseInterface
}

/*type MiddlewareInterface interface {
	NewMiddleware(e *echo.Echo)
	AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CustomHTTPErrorHandler(err error, ctx echo.Context)
	PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}*/
