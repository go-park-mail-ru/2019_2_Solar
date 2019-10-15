package main
import (
	//"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	//"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	middleware "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"sync"
)

func main() {
	//var userSlice repository.UsersSlice
	//err := repository.DBWorker.UniversalRead(middleware.QueryReadUserByCookie, &userSlice)
	//fmt.Println(err)
	//fmt.Println(userSlice)
	e := echo.New()
	e.Use(middleware.CORSMiddleware)
	e.Use(middleware.AuthenticationMiddleware)
	//e.Use(echomiddleware.Logger())
	//e.Use(middleware.PanicMiddleware)
	//e.HTTPErrorHandler = middleware.ErrorHandler

	delivery.NewHandlers(e, usecase.NewPinterestUsecase([]models.User{}, []models.UserSession{}, &sync.Mutex{}))

	//e.Logger.Warnf("start listening on %s", listenAddr)
	err := e.Start("127.0.0.1:8080")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
