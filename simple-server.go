package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	repositoryMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase/middleware"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
	customMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"github.com/labstack/echo-contrib/prometheus"
	//"github.com/labstack/echo"
	echov4 "github.com/labstack/echo/v4"
	"sync"
)

func main() {

	authorizationService := functions.AuthServiceCreate("authorization-service")
	pinBoardService := functions.PinBoardServiceCreate("pinboard-service")
	userService := functions.UserServiceCreate("user-service")

	e := echov4.New()
	//ev4 := echov4.New()
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	middlewares := customMiddleware.MiddlewareStruct{}
	mRep := repositoryMiddleware.MRepositoryStruct{}
	err := mRep.DataBaseInit()
	if err != nil {
		e.Logger.Error("can't connect to database " + err.Error())
		return
	}
	mUseCase := useCaseMiddleware.MUseCaseStruct{}
	mUseCase.NewUseCase(&mRep)
	middlewares.NewMiddleware(e, &mUseCase, authorizationService)
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
	err = handlers.NewHandlers(e, &useCase, authorizationService, pinBoardService, userService)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	// Enable metrics middleware
	//go func() { ev4.Logger.Fatal(e.Start(":8082")) }()

	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	if err := e.Start(consts.HostAddress); err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")
}
