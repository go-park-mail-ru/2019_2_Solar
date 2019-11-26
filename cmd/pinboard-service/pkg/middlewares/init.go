package middlewares

import (
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
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
