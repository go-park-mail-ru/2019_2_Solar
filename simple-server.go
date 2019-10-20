package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	customMiddlewares "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"sync"
)

func main() {
	e := echo.New()
	e.Use(customMiddlewares.CORSMiddleware)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	e.Use(customMiddlewares.PanicMiddleware)
	//e.Use(customMiddleware.AccessLogMiddleware)
	e.Use(customMiddlewares.AuthenticationMiddleware)
	e.HTTPErrorHandler = customMiddlewares.CustomHTTPErrorHandler

	handlers := delivery.HandlersStruct{}
	var mutex sync.Mutex
	rep := repository.RepositoryStruct{}
	err := rep.NewDataBaseWorker()
	if err != nil {
		return
	}
	useCase := usecase.UsecaseStruct{}
	useCase.NewUseCase(&mutex, &rep)
	handlers.NewHandlers(e, &useCase)
	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	err = e.Start(consts.HostAddress)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
