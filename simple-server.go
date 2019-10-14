package main

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	"net/http"
)

func main() {
	http.Handle("/", middlewares.CORSMiddleware(http.HandlerFunc(delivery.HandleRoot)))
	http.Handle("/users/", middlewares.CORSMiddleware(http.HandlerFunc(delivery.HandleUsers)))
	http.Handle("/registration/", middlewares.CORSMiddleware(http.HandlerFunc(delivery.HandleRegistration)))
	http.Handle("/login/", middlewares.CORSMiddleware(http.HandlerFunc(delivery.HandleLogin)))
	http.Handle("/logout/", middlewares.CORSMiddleware(http.HandlerFunc(delivery.HandleLogout)))
	http.Handle("/profile/data", middlewares.CORSMiddleware(http.HandlerFunc(delivery.HandleProfileData)))
	http.Handle("/profile/picture", middlewares.CORSMiddleware(http.HandlerFunc(delivery.HandleProfilePicture)))

	http.ListenAndServe(":8080", nil)
}
