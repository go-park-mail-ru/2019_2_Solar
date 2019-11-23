package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/support/repository"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
	"github.com/gorilla/websocket"
	"sync"
)

type UseStruct struct {
	PRepository repository.ReposInterface
	Hub         webSocket.HubStruct
	Mu          *sync.Mutex
}

type UseInterface interface {
	CreateClient(conn *websocket.Conn, userId uint64)
	NewUseCase(mu *sync.Mutex, rep repository.ReposInterface,
		hub webSocket.HubStruct)
}

func (USC *UseStruct) NewUseCase(mu *sync.Mutex, rep repository.ReposInterface,
	hub webSocket.HubStruct) {
	USC.Mu = mu
	USC.PRepository = rep
	USC.Hub = hub
	go USC.Hub.Run()
}