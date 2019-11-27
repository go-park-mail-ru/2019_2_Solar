package usecase

import (
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/gorilla/websocket"
)

func (USC *UseStruct) CreateClient(conn *websocket.Conn, user models.User) {
	client := &webSocket.Client{Hub: USC.ReturnHub(), Conn: conn, Send: make(chan models.ChatMessage), User: user}
	client.Hub.Register <- client
	go client.ReadPump(USC.PRepository)
	client.WritePump()
}
