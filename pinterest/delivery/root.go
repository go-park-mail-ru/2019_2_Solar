package delivery

import (
	"encoding/json"
	"github.com/labstack/echo"
	"log"
)

func (h *HandlersStruct) HandleEmpty(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	encoder := json.NewEncoder(ctx.Response())
	data := h.PUsecase.SetJsonData(nil, "Empty handler has been done")
	encoder.Encode(data)
	log.Printf("Empty handler has been done")
	return nil
}