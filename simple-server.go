package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	middleware "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"sync"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSMiddleware)
	//e.Use(middleware.AuthenticationMiddleware)
	//e.Use(echomiddleware.Logger())
	//e.Use(middleware.PanicMiddleware)
	//e.HTTPErrorHandler = middleware.ErrorHandler
	handlers := delivery.Handlers{
		PUsecase: &usecase.PinterestUseCase{},
		Users:    make([]models.User, 0),
		Sessions: make([]models.UserSession, 0),
		Mu:       &sync.Mutex{},
	}
	NewHandlers(e, handlers)

	//e.Logger.Warnf("start listening on %s", listenAddr)
	err := e.Start("127.0.0.1:8080")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

	/*	http.Handle("/", middleware.CORSMiddleware(http.HandlerFunc(delivery.HandleRoot)))
		http.Handle("/users/", middleware.CORSMiddleware(http.HandlerFunc(delivery.HandleUsers)))
		http.Handle("/registration/", middleware.CORSMiddleware(http.HandlerFunc(delivery.HandleRegistration)))
		http.Handle("/login/", middleware.CORSMiddleware(http.HandlerFunc(delivery.HandleLogin)))
		http.Handle("/logout/", middleware.CORSMiddleware(http.HandlerFunc(delivery.HandleLogout)))
		http.Handle("/profile/data", middleware.CORSMiddleware(http.HandlerFunc(delivery.HandleProfileData)))
		http.Handle("/profile/picture", middleware.CORSMiddleware(http.HandlerFunc(delivery.HandleProfilePicture)))

		http.ListenAndServe(":8080", nil)*/
}

func NewHandlers(e *echo.Echo, Handler delivery.Handlers) {
	e.GET("/", delivery.HandleRoot)
	e.GET("/users/", delivery.HandleUsers)
	e.POST("/registration/", Handler.HandleRegUser)
	e.POST("/login/", Handler.HandleLoginUser)
	e.GET("/logout/", Handler.HandleLogoutUser)
	e.GET("/profile/data", Handler.HandleGetProfileUserData)
	e.GET("/profile/picture", Handler.HandleEditProfileUserPicture)
	e.POST("/profile/data", Handler.HandleEditProfileUserData)
	e.POST("/profile/picture", Handler.HandleGetProfileUserPicture)
}
