package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/consts"
	customMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/middlewares"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/support/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/support/repository"
	repositoryMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/repository/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/support/usecase"
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/usecase/middleware"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
	"github.com/labstack/echo"
	"sync"
)

func main() {
	e := echo.New()
	middlewares := customMiddleware.MiddlewareStruct{}
	mRep := repositoryMiddleware.MRepositoryStruct{}
	err := mRep.DataBaseInit()
	if err != nil {
		e.Logger.Error("can't connect to database " + err.Error())
		return
	}
	mUseCase := useCaseMiddleware.MUseCaseStruct{}
	mUseCase.NewUseCase(&mRep)
	middlewares.NewMiddleware(e, &mUseCase)
	//e.Use(customMiddleware.CORSMiddleware)
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	//e.Use(customMiddleware.PanicMiddleware)
	//e.Use(customMiddleware.AccessLogMiddleware)
	//e.Use(customMiddleware.AuthenticationMiddleware)
	//e.HTTPErrorHandler = customMiddleware.CustomHTTPErrorHandler

	e.Static("/static", "static")

	handlers := delivery.HandlersStruct{}
	var mutex sync.Mutex
	rep := repository.ReposStruct{}
	err = rep.DataBaseInit()
	if err != nil {
		e.Logger.Error("can't connect to database " + err.Error())
		return
	}
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	useCase := usecase.UseStruct{}
	useCase.NewUseCase(&mutex, &rep, &san, hub)
	err = handlers.NewHandlers(e, &useCase)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	if err := e.Start(consts.HostAddress); err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")
}
