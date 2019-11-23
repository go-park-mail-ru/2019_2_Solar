package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/support/usecase"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) NewHandlers(e *echo.Echo, useCase usecase.UseInterface) error {
	h.PUsecase = useCase
	e.GET("/chat", h.HandleUpgradeWebSocket)
	/*e.POST("/login", h.HandleLoginUser)
	e.GET("/profile/data", h.HandleGetProfileUserData)
*/
	return nil
}
