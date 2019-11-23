package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
	"github.com/gorilla/websocket"
)

func (USC *UseStruct) CreateClient(conn *websocket.Conn, userId uint64) {
	client := &webSocket.Client{Hub: USC.ReturnHub(), Conn: conn, Send: make(chan models.ChatMessage), UserId: userId}
	client.Hub.Register <- client
	go client.ReadPump(USC.PRepository)
	go client.WritePump()
}

func (USC *UseStruct) ReturnHub() *webSocket.HubStruct {
	return &USC.Hub
}
