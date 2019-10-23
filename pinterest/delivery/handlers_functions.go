package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/labstack/echo"
)

func (h *HandlersStruct)NewHandlers(e *echo.Echo, IUsecase usecase.UsecaseInterface) {
	h.PUsecase = IUsecase
	e.GET("/", h.HandleEmpty)

	e.GET("/users", h.HandleListUsers)
	e.GET("/users/:email", h.HandleGetUserByEmail)

	e.POST("/registration", h.HandleRegUser)
	e.POST("/login", h.HandleLoginUser)
	e.POST("/logout", h.HandleLogoutUser)

	e.GET("/profile/data", h.HandleGetProfileUserData)
	//e.GET("/profile/picture", h.HandleGetProfileUserPicture)

	e.POST("/profile/data", h.HandleEditProfileUserData)
	e.POST("/profile/picture", h.HandleEditProfileUserPicture)

	e.POST("/board", h.HandleCreateBoard)
	//e.GET( "/board/:id", h.HandleGetPin)

	e.POST("/pin", h.HandleCreatePin)
	e.GET( "/pin/:id", h.HandleGetPin)
}
