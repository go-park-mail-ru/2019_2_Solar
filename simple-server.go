package main

import (
	"2019_2_Solar/pkg/handls"
	"net/http"
)

func main() {
	http.Handle("/", handls.CORSMiddleware(http.HandlerFunc(handls.HandleRoot)))
	http.Handle("/users/", handls.CORSMiddleware(http.HandlerFunc(handls.HandleUsers)))
	http.Handle("/registration/", handls.CORSMiddleware(http.HandlerFunc(handls.HandleRegistration)))
	http.Handle("/login/", handls.CORSMiddleware(http.HandlerFunc(handls.HandleLogin)))
	http.Handle("/logout/", handls.CORSMiddleware(http.HandlerFunc(handls.HandleLogout)))
	http.Handle("/profile/data", handls.CORSMiddleware(http.HandlerFunc(handls.HandleProfileData)))
	http.Handle("/profile/picture", handls.CORSMiddleware(http.HandlerFunc(handls.HandleProfilePicture)))

	http.ListenAndServe(":8080", nil)
}
