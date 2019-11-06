package webSocket

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/gorilla/websocket"
	"log"
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
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the Hub.
type Client struct {
	Hub    *HubStruct
	UserId uint64
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
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		newChatMessage := models.NewChatMessage{}
		json.Unmarshal(message, newChatMessage)
		var params1 []interface{}
		params1 = append(params1, newChatMessage.IdSender, newChatMessage.UserNameRecipient, newChatMessage.Message, time.Now())
		_, _ = PRepository.Insert(consts.INSERTChatMessage, params1)
		var params2 []interface{}
		params2 = append(params2, newChatMessage.IdSender, newChatMessage.UserNameRecipient)
		idRecipient, _ := PRepository.SelectFullUser(consts.SELECTUserByUsername, params2)
		chatMessage := models.ChatMessage{
			IdSender:    newChatMessage.IdSender,
			IdRecipient: idRecipient[0].ID,
			Message:     newChatMessage.Message,
			SendTime:    time.Now(),
			IsDeleted:   false,
		}
		c.Hub.Broadcast <- chatMessage
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

			/*			w, err := c.Conn.NextWriter(websocket.TextMessage)
						if err != nil {
							return
						}

						w.Write(message)

						// Add queued chat messages to the current websocket message.
						n := len(c.Send)
						for i := 0; i < n; i++ {
							w.Write(newline)
							w.Write(<-c.Send)
						}

						if err := w.Close(); err != nil {
							return
						}*/
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
