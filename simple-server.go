package main

import (
	"net/http"
	"sync"
)

var handlers = Handlers{
	users:    make([]User, 0),
	sessions: make([]UserSession, 0),
	mu:       &sync.Mutex{},
}

func main() {

	http.Handle("/", CORSMiddleware(http.HandlerFunc(HandleRoot)))
	http.Handle("/users/", CORSMiddleware(http.HandlerFunc(HandleUsers)))
	http.Handle("/registration/", CORSMiddleware(http.HandlerFunc(HandleRegistration)))
	http.Handle("/login/", CORSMiddleware(http.HandlerFunc(HandleLogin)))
	http.Handle("/logout/", CORSMiddleware(http.HandlerFunc(HandleLogout)))
	http.Handle("/profile/data", CORSMiddleware(http.HandlerFunc(HandleProfileData)))
	http.Handle("/profile/picture", CORSMiddleware(http.HandlerFunc(HandleProfilePicture)))

	http.ListenAndServe(":8080", nil)
}
