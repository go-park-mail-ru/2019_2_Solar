package middlewares

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/consts"
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/usecase/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (MS *MiddlewareStruct) NewMiddleware(e *echo.Echo, mRep useCaseMiddleware.MUseCaseInterface) {
	MS.MUsecase = mRep
	e.Use(MS.CORSMiddleware)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	e.Use(MS.PanicMiddleware)
	e.Use(MS.AuthenticationMiddleware)
	e.HTTPErrorHandler = MS.CustomHTTPErrorHandler
}
