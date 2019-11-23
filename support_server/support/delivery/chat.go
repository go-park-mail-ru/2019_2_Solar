package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) HandleUpgradeWebSocket(ctx echo.Context) (Err error) {
	getUser := ctx.Get("User")
	getAdmin := ctx.Get("Admin")
	if getAdmin == nil && getUser == nil {
		return errors.New("not authorized")
	}
	role := ""
	if getAdmin == nil {
		role = "admin"
	} else {
		role = "user"
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
	user := models.User{ID:getUser.(models.User).ID}
	h.PUsecase.CreateClient(ws, user.ID, role)

/*	body := models.BodyInfo{Info: "OK"}
	jsonStruct := models.JSONResponse{Body: body}
	if err := ctx.JSON(200, jsonStruct); err != nil {
		return err
	}*/
	return nil
}
