package middlewares

import (
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase/middleware"
)

type MiddlewareStruct struct {
	MUsecase useCaseMiddleware.MUseCaseInterface
	MAuth services.AuthorizationServiceClient
	PinBoardService pinboard_service.PinBoardServiceClient
}

/*type MiddlewareInterface interface {
	NewMiddleware(e *echo.Echo)
	AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CustomHTTPErrorHandler(err error, ctx echo.Context)
	PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}*/
