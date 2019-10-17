package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	middleware "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"github.com/labstack/echo"
	"sync"
)

func main() {
	//var userSlice repository.UsersSlice
	//err := repository.DBWorker.DBDataRead(middleware.QueryReadUserByCookie, &userSlice)
	//fmt.Println(err)
	//fmt.Println(userSlice)
	e := echo.New()
	e.Use(middleware.CORSMiddleware)
	//e.Use(middleware.PanicMiddleware)
	//e.Use(echomiddleware.Logger())

	e.Use(middleware.AuthenticationMiddleware)
	//e.HTTPErrorHandler = middleware.ErrorHandler

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

	//e.Logger.Warnf("start listening on %s", listenAddr)
	err = e.Start("127.0.0.1:8080")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
