package delivery

import (
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	user_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/user-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/labstack/echo"
)


func (h *HandlersStruct) NewHandlers(e *echo.Echo, useCase usecase.UseInterface, auth services.AuthorizationServiceClient,
	pinBoardService pinboard_service.PinBoardServiceClient, userService user_service.UserServiceClient) error {
	h.PUsecase = useCase

	h.AuthSessManager = auth
	h.PinBoardService = pinBoardService
	h.UserService = userService


	e.GET("/", h.HandleEmpty)

	// ==============================================================


	e.GET("/users", h.HandleListUsers)
	e.GET("/users/:username", h.HandleGetUserByUsername)
	e.POST("/subscribe/:username", h.HandleCreateSubscribe)
	e.DELETE("/subscribe/:username", h.HandleDeleteSubscribe)

	e.GET("/service/users", h.ServiceGetUsers)
	e.GET("/service/users/:username", h.ServiceGetUserByUsername)
	e.POST("/service/subscribe/:username", h.ServiceCreateSubscribe)
	e.DELETE( "/service/subscribe/:username", h.ServiceDeleteSubscribe)


	// ==============================================================
	// ==============================================================

	e.POST("/registration", h.HandleRegUser)
	e.POST("/login", h.HandleLoginUser)
	e.POST("/logout", h.HandleLogoutUser)

	e.POST("/service/registration", h.ServiceRegUser)
	e.POST("/service/login", h.ServiceLoginUser)
	e.POST( "/service/logout", h.ServiceLogoutUser)

	// ==============================================================

	e.GET("/profile/data", h.HandleGetProfileUserData)

	e.POST("/profile/data", h.HandleEditProfileUserData)
	e.POST("/profile/picture", h.HandleEditProfileUserPicture)

	// ==============================================================

	e.POST("/board", h.HandleCreateBoard)
	e.GET("/board/:id", h.HandleGetBoard)
	e.POST("/pin", h.HandleCreatePin)
	e.GET("/pin/:id", h.HandleGetPin)


	e.POST("/service/board", h.ServiceCreateBoard)
	e.GET( "/service/board/:id", h.ServiceGetBoard)
	e.POST("/service/pin", h.ServiceCreatePin)
	e.GET("/service/pin/:id", h.ServiceGetPin)

	// ==============================================================

	e.GET("/board/list/my", h.HandleGetMyBoards)

	e.POST("/pin/:id/comment", h.HandleCreateComment)
	e.GET("/pin/list/new", h.HandleGetNewPins)
	e.GET("/pin/list/my", h.HandleGetMyPins)
	e.GET("/pin/list/subscribe", h.HandleGetSubscribePins)

	e.POST("/notice/:receiver_id", h.HandleCreateNotice)
	e.GET( "/notice", h.HandleGetNotices)

	e.GET("/chat", h.HandleUpgradeWebSocket)

	e.GET( "/find/pins/by/tag/:tag", h.HandlerFindPinByTag)
	e.GET( "/find/users/by/username/:username", h.HandlerFindUserByUsername)

	e.GET ("/categories", h.HandleGetCategories)

	return nil
}
