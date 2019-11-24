package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/authorization-service/pinterest/usecase"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) NewHandlers(e *echo.Echo, useCase usecase.UseInterface) error {
	h.PUsecase = useCase

	e.POST("/registration", h.HandleRegUser)
	e.POST("/login", h.HandleLoginUser)
	e.POST("/logout", h.HandleLogoutUser)

	e.GET()

	return nil
}
