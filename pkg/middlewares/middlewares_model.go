package middlewares

import (
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
)

type MiddlewareStruct struct {
	MUsecase useCaseMiddleware.MUseCaseInterface
	MAuth functions.Auth
}

/*type MiddlewareInterface interface {
	NewMiddleware(e *echo.Echo)
	AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CustomHTTPErrorHandler(err error, ctx echo.Context)
	PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}*/
