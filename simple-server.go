package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
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
	handlers := delivery.HandlersStruct{}
	if err := handlers.NewHandlers(e); err != nil {
		e.Logger.Errorf("server error: %s", err)
		return
	}
	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	if err := e.Start(consts.HostAddress); err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
