package delivery

import (
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"sort"
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

	senderId := user.ID

	receiverIdStr := ctx.Param("recipientId")

	receiverIdInt, err := strconv.Atoi(receiverIdStr)
	if err != nil {
		return err
	}

	receiverId := uint64(receiverIdInt)

	messages, err := h.PUsecase.GetMessages(senderId, receiverId)
	if err != nil {
		return err
	}

	body := struct {
		Messages []models.OutputMessage `json:"messages"`
		Info     string                 `json:"info"`
	}{messages, "OK"}

	data := models.ValeraJSONResponse{ctx.Get("token").(string), body}
	return ctx.JSON(http.StatusOK, data)
}

func (h *HandlersStruct) HandleChatRecipient(ctx echo.Context) (Err error) {
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
	messages, err := h.PUsecase.GetRecipients(user.ID)
	if err != nil {
		return err
	}
	var uniqueMes []models.MessageWithUsername
	for i := 0; i < len(messages)-1; i++ {
		mes := messages[i]
		for j := i + 1; j < len(messages); j++ {
			if messages[i].SenderUserName == messages[j].RecipientUserName && messages[i].RecipientUserName == messages[j].SenderUserName && messages[i].SendTime.Before(messages[j].SendTime) {
				mes = messages[j]
				continue
			}
		}

		elem := sort.Search(len(uniqueMes), func(elem int) bool { return uniqueMes[elem] == mes })
		if elem < len(uniqueMes) && uniqueMes[elem] == mes {
			continue
		} else {
			uniqueMes = append(uniqueMes, mes)
		}
	}

	data := models.ValeraJSONResponse{CSRF: ctx.Get("token").(string), Body: uniqueMes}
	return ctx.JSON(http.StatusOK, data)
}
