package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) HandlerFindPinByTag(ctx echo.Context) (Err error) {
	ctx.Response()
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")

	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}

	tagName := ctx.Param("tag")

	pins, err := h.PUsecase.SearchPinsByTag(tagName)
	if err != nil {
		return err
	}
	jsonStruct := models.JSONResponse{Body: pins}
	if err := ctx.JSON(200, jsonStruct); err != nil {
		return err
	}
	return nil

}