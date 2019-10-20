package middlewares

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)


func PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				ctx.Logger().Error("recovered ", panicErr)
				err = &echo.HTTPError{Code: 500, Message: "Internal server error"}
			}
		}()
		err = next(ctx)
		return err
	}
}

func CustomHTTPErrorHandler(err error, ctx echo.Context) {
	var jsonError error
	switch err := errors.Cause(err); err.(type) {
	case *echo.HTTPError:
		jsonError = ctx.JSON(err.(*echo.HTTPError).Code, struct {
			Body string
		}{Body: err.(*echo.HTTPError).Message.(string)})
	case nil:
		return
	default:
		ctx.Logger().Warn(err)
		jsonError = ctx.JSON(400, struct {
			Body struct{ Info string }
		}{Body: struct{ Info string }{Info: err.Error()}})
	}
	if jsonError != nil {
		ctx.Logger().Error("Server cant repay response")
	}
}
