package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	middleware "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSMiddleware)
	//e.Use(middleware.AuthenticationMiddleware)
	//e.Use(echomiddleware.Logger())
	//e.Use(middleware.PanicMiddleware)
	//e.HTTPErrorHandler = middleware.ErrorHandler

	delivery.NewHandlers(e)

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

/*func NewHandlers(e *echo.Echo, Handler delivery.Handlers) {
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
*/