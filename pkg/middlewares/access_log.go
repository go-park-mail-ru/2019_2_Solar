package middlewares

import (
	"github.com/labstack/echo"
	logger "github.com/sirupsen/logrus"
	"time"
)

func AccessLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()
		err := next(ctx)
		logger.WithFields(logger.Fields{
			"method":      ctx.Request().Method,
			"remote_addr": ctx.Request().RemoteAddr,
			"work_time":   time.Since(start),
		}).Info(ctx.Request().URL.Path)
		return err
	}
}
