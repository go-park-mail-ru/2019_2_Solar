package delivery

import (
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) HandleUpgradeWebSocket(ctx echo.Context) (Err error) {
	ws, err := webSocket.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := ws.Close(); err != nil {
			Err = errors.Wrap(Err, err.Error())
		}
	}()
	client := &webSocket.Client{Hub: h.PUsecase.ReturnHub(), Conn: ws, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
	return nil
}
