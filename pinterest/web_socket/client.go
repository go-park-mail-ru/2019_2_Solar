package webSocket

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

/*var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)*/

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the Hub.
type Client struct {
	Hub  *HubStruct
	User models.User
	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan models.ChatMessage
}

// ReadPump pumps messages from the websocket connection to the Hub.
//
// The application runs ReadPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump(PRepository repository.ReposInterface) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	err := c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(err)
				log.Printf("error: %v", err)
			}
			break
		}
		newChatMessage := new(models.NewChatMessage)
		err = json.Unmarshal(message, newChatMessage)
		if err != nil {
			fmt.Println(err)
			break
		}

		userRecipients, err := PRepository.SelectUsersByUsername(newChatMessage.UserNameRecipient)
		if err != nil {
			fmt.Println(err)
			break
		}
		if len(userRecipients) != 1 {
			fmt.Println("Нет юзера")
			break
		}

		userRecipient := userRecipients[0]

		saveMessage := models.SaveMessage{
			IdSender:    c.User.ID,
			IdRecipient: userRecipient.ID,
			Message:     newChatMessage.Message,
			SendTime:    time.Now(),
			IsDeleted:   false,
		}

		_, err = PRepository.InsertChatMessage(saveMessage)
		if err != nil {
			fmt.Println(err)
			break
		}
		if saveMessage.IdRecipient == saveMessage.IdSender {
			continue
		}
		chatMessage := models.ChatMessage{
			UserNameSender: c.User.Username,
			IdRecipient:    saveMessage.IdRecipient,
			Message:        saveMessage.Message,
			SendTime:       saveMessage.SendTime,
			IsDeleted:      saveMessage.IsDeleted,
		}

		for client := range c.Hub.Clients {
			if client.User.ID == chatMessage.IdRecipient {
				client.Send <- chatMessage
			}
		}
		//c.Hub.Broadcast <- chatMessage
	}
}

// WritePump pumps messages from the Hub to the websocket connection.
//
// A goroutine running WritePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The Hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteJSON(message)
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
