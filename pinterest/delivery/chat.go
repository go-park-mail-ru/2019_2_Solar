package delivery

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/vars"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) HandleUpgradeWebSocket(ctx echo.Context) (Err error) {
	ws, err := vars.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := ws.Close(); err != nil {
			Err = errors.Wrap(Err, err.Error())
		}
	}()
	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			ctx.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			ctx.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
	return nil
}


