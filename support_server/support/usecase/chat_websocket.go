package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
	"github.com/gorilla/websocket"
)

func (USC *UseStruct) CreateClient(conn *websocket.Conn, userId uint64, role string) {
	client := &webSocket.Client{Hub: USC.ReturnHub(), Conn: conn, Send: make(chan models.ChatMessage), UserId: userId, Role:role}
	client.Hub.Register <- client
	go client.ReadPump(USC.PRepository)
	//time.Sleep(1*time.Second)
	client.WritePump()
	//time.Sleep(1*time.Second)
}

func (USC *UseStruct) ReturnHub() *webSocket.HubStruct {
	return &USC.Hub
}
