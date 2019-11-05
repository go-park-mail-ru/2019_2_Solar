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
	data := h.PUsecase.SetJSONData(nil, ctx.Get("token").(string), "Empty handler has been done")
	if err := encoder.Encode(data); err != nil {
		return err
	}
	log.Printf("Empty handler has been done")
	return nil
}
