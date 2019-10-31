package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	customMiddlewares "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(customMiddlewares.CORSMiddleware)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	e.Use(customMiddlewares.PanicMiddleware)
	//e.Use(customMiddleware.AccessLogMiddleware)
	e.Use(customMiddlewares.AuthenticationMiddleware)
	e.HTTPErrorHandler = customMiddlewares.CustomHTTPErrorHandler
	e.Static("/static", "static")
	useCase := usecase.UsecaseStruct{}
	if err := useCase.NewUseCase(); err != nil {
		e.Logger.Errorf("server error: %s", err)
		return
	}
	handlers := delivery.HandlersStruct{}
	handlers.NewHandlers(e, &useCase)
	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	if err := e.Start(consts.HostAddress); err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
