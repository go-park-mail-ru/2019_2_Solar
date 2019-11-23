package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/support/repository"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type UseStruct struct {
	PRepository repository.ReposInterface
	Hub         webSocket.HubStruct
	Mu          *sync.Mutex
}

type UseInterface interface {
	CreateClient(conn *websocket.Conn, user models.User)
	NewUseCase(mu *sync.Mutex, rep repository.ReposInterface,
		hub webSocket.HubStruct)

	GetUserByID(userID uint64) (User models.User, Err error)
	//GetAllUsers() ([]models.AnotherUser, error)

	GetAdminByLogin(login string) (Admin models.Admin, Err error)
	CompareAdminPassword(password string, autorizedPassword string) (Err error)
	AddNewAdminSession(adminID uint64) (Cookie http.Cookie, Err error)

	GetHubListActiveUsers() ([]uint64, error)
}

func (USC *UseStruct) NewUseCase(mu *sync.Mutex, rep repository.ReposInterface,
	hub webSocket.HubStruct) {
	USC.Mu = mu
	USC.PRepository = rep
	USC.Hub = hub
	go USC.Hub.Run()
}