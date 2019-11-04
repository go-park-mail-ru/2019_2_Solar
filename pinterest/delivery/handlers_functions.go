package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) NewHandlers(e *echo.Echo) error {
	useCase := usecase.UseStruct{}
	if err := useCase.NewUseCase(); err != nil {
		return err
	}
	h.PUsecase = &useCase
	e.GET("/", h.HandleEmpty)

	e.GET("/users", h.HandleListUsers)
	e.GET("/users/:username", h.HandleGetUserByUsername)

	e.POST("/subscribe/:username", h.HandleCreateSubscribe)
	e.DELETE("/subscribe/:username", h.HandleDeleteSubscribe)

	e.POST("/registration", h.HandleRegUser)
	e.POST("/login", h.HandleLoginUser)
	e.POST("/logout", h.HandleLogoutUser)

	e.GET("/profile/data", h.HandleGetProfileUserData)
	//e.GET("/profile/picture", h.HandleGetProfileUserPicture)

	e.POST("/profile/data", h.HandleEditProfileUserData)
	e.POST("/profile/picture", h.HandleEditProfileUserPicture)

	e.POST("/board", h.HandleCreateBoard)
	e.GET("/board/:id", h.HandleGetBoard)

	e.POST("/pin", h.HandleCreatePin)
	e.POST("/pin/:id/comment", h.HandleCreateComment)
	e.GET("/pin/:id", h.HandleGetPin)
	e.GET("/pin/list/new", h.HandleGetNewPins)
	e.GET("/pin/list/my", h.HandleGetMyPins)
	e.GET("/pin/list/subscribe", h.HandleGetSubscribePins)

	e.POST("/notice/:receiver_id", h.HandleCreateNotice)
	return nil
}
