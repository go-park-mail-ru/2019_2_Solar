package middlewares

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
)

func PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(ctx.Response(), "Internal server error", 500)
			}
		}()
		err := next(ctx)
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
	default:
		ctx.Logger().Error(err)
	}
	if jsonError != nil {

	}
	/*
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		errorPage := fmt.Sprintf("%d.html", code)
		if err := ctx.File(errorPage); err != nil {
			ctx.Logger().Error(err)
		}
		ctx.Logger().Error(err)*/
}
