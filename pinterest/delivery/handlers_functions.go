package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) NewHandlers(e *echo.Echo, useCase usecase.UseInterface) error {
	h.PUsecase = useCase
	prefix := "/api/v1"

	e.GET(prefix + "/", h.HandleEmpty)

	e.GET(prefix + "/users", h.HandleListUsers)
	e.GET(prefix + "/users/:username", h.HandleGetUserByUsername)

	e.POST(prefix + "/subscribe/:username", h.HandleCreateSubscribe)
	e.DELETE(prefix +"/subscribe/:username", h.HandleDeleteSubscribe)

	e.POST(prefix +"/registration", h.HandleRegUser)
	e.POST(prefix +"/login", h.HandleLoginUser)
	e.POST(prefix +"/logout", h.HandleLogoutUser)

	e.GET(prefix +"/profile/data", h.HandleGetProfileUserData)

	e.POST(prefix +"/profile/data", h.HandleEditProfileUserData)
	e.POST(prefix + "/profile/picture", h.HandleEditProfileUserPicture)

	e.POST(prefix +"/board", h.HandleCreateBoard)
	e.GET(prefix +"/board/:id", h.HandleGetBoard)
	e.GET(prefix +"/board/list/my", h.HandleGetMyBoards)

	e.POST(prefix +"/pin", h.HandleCreatePin)
	e.POST(prefix +"/add/pin", h.HandleAddPin)
	e.POST(prefix +"/pin/:id/comment", h.HandleCreateComment)
	e.GET(prefix +"/pin/:id", h.HandleGetPin)
	e.GET(prefix +"/pin/list/new", h.HandleGetNewPins)
	e.GET(prefix +"/pin/list/my", h.HandleGetMyPins)
	e.GET(prefix +"/pin/list/subscribe", h.HandleGetSubscribePins)

	e.POST(prefix +"/notice/:receiver_id", h.HandleCreateNotice)
	e.GET(prefix +"/notice", h.HandleGetNotices)

	e.GET(prefix +"/chat", h.HandleUpgradeWebSocket)
	e.GET(prefix +"/chat/messages/:recipientId", h.HandleGetMessages)
	e.GET(prefix +"/chat/recipients", h.HandleChatRecipient)

	e.GET(prefix +"/find/pins/by/tag/:tag", h.HandlerFindPinByTag)
	e.GET(prefix +"/find/users/by/username/:username", h.HandlerFindUserByUsername)

	e.GET (prefix +"/categories", h.HandleGetCategories)

	e.POST(prefix +"/admin/fill", h.HandleAdminFill)
	e.POST(prefix +"/admin/fill/christmas", h.HandleAdminFillChristmas)



	return nil
}
