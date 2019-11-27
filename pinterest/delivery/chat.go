package delivery

import (
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"strconv"
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
	h.PUsecase.CreateClient(ws, user)

	/*	body := models.BodyInfo{Info: "OK"}
		jsonStruct := models.JSONResponse{Body: body}
		if err := ctx.JSON(200, jsonStruct); err != nil {
			return err
		}*/
	return nil
}

func (h *HandlersStruct) HandleGetMessages(ctx echo.Context) (Err error) {
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")

	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	user := getUser.(models.User)

	senderIdStr := ctx.Param("senderId")
	senderIdInt, err := strconv.Atoi(senderIdStr)
	if err != nil {
		return nil
	}
	senderId := uint64(senderIdInt)

	receiverIdStr := ctx.Param("receiverId")
	receiverIdInt, err := strconv.Atoi(receiverIdStr)
	if err != nil {
		return err
	}
	receiverId := uint64(receiverIdInt)

	if user.ID != senderId && user.ID != receiverId {
		return errors.New("Not your chat!")
	}

	messages, err := h.PUsecase.GetMessages(senderId, receiverId)
	if err != nil {
		return err
	}

	body := struct {
		Messages []models.OutputMessage `json:"messages"`
		Info     string                 `json:"info"`
	}{messages, "OK"}

	data := models.ValeraJSONResponse{ctx.Get("token").(string), body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}
