package delivery

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"log"
)

func (h *HandlersStruct) HandleEmpty(ctx echo.Context) (Err error) {
	ctx.Response()
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	encoder := json.NewEncoder(ctx.Response())
	data := h.PUsecase.SetJsonData(nil, "Empty handler has been done")
	encoder.Encode(data)
	log.Printf("Empty handler has been done")
	return nil
}