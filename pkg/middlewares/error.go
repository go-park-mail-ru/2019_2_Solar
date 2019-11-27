package middlewares

import (
	echov4 "github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (MS *MiddlewareStruct) PanicMiddleware(next echov4.HandlerFunc) echov4.HandlerFunc {
	return func(ctx echov4.Context) (err error) {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				ctx.Logger().Error("recovered ", panicErr)
				err = &echov4.HTTPError{Code: 500, Message: "Internal server error"}
			}
		}()
		err = next(ctx)
		return err
	}
}

func (MS *MiddlewareStruct) CustomHTTPErrorHandler(err error, ctx echov4.Context) {
	var jsonError error
	switch err := errors.Cause(err); err.(type) {
	case *echov4.HTTPError:
		ctx.Logger().Warn(err)
		jsonError = ctx.JSON(err.(*echov4.HTTPError).Code, struct {
			Body string `json:"body"`
		}{Body: err.(*echov4.HTTPError).Message.(string)})
	case nil:
		return
	default:
		//ctx.Logger().Error(err)
		//ctx.Logger().Warn(err)
		ctx.Logger().Info(err)
		//ctx.Logger().Fatal(err)
		jsonError = ctx.JSON(400, struct {
			Body struct {
				Info string `json:"info"`
			} `json:"body"`
		}{Body: struct {
			Info string `json:"info"`
		}{Info: err.Error()}})
	}
	if jsonError != nil {
		ctx.Logger().Error("Server cant repay response")
	}
}
