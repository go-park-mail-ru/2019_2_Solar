package delivery

import (
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) HandleUpgradeWebSocket(ctx echo.Context) (Err error) {
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	ws, err := webSocket.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := ws.Close(); err != nil {
			Err = errors.Wrap(Err, err.Error())
		}
	}()
	user := getUser.(models.User)
	h.PUsecase.CreateClient(ws, user.ID)

	body := models.BodyInfo{Info: "OK"}
	jsonStruct := models.JSONResponse{Body: body}
	if err := ctx.JSON(200, jsonStruct); err != nil {
		return err
	}
	return nil
}
