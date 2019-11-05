package webSocket

import "github.com/go-park-mail-ru/2019_2_Solar/pkg/models"

type HubStruct struct {
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the Clients.
	Broadcast chan models.ChatMessage

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	Unregister chan *Client
}

func (h *HubStruct) NewHub() {
	h.Broadcast = make(chan models.ChatMessage)
	h.Register = make(chan *Client)
	h.Unregister = make(chan *Client)
	h.Clients = make(map[*Client]bool)
}

func (h *HubStruct) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case chatMessage := <-h.Broadcast:
			for client := range h.Clients {
				if client.UserId == chatMessage.IdRecipient {
					select {
					case client.Send <- chatMessage:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}
