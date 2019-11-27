package middlewares

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"time"
)

func (MS *MiddlewareStruct) AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokens, _ := functions.NewAesCryptHashToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
		cookie, err := ctx.Cookie("session_key")
		if err != nil {
			return next(ctx)
		}

		sCookie := services.Cookie{
			Key:                  cookie.Name,
			Value:                cookie.Value,
			Exp:                  cookie.Expires.String(),
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}

		sctx := context.Background()
		sUserSession, err := MS.MAuth.CheckSession(sctx, &sCookie)
		if err != nil {
			return err
		}


		userSession := models.UserSession{
			ID:         sUserSession.ID,
			UserID:     sUserSession.UserID,
		}

		user, err := MS.MUsecase.GetUserByCookieValue(cookie.Value)
		if err != nil {
			return err
		}

		//userSession, err := MS.MUsecase.GetSessionsByCookieValue(cookie.Value)
		//if err != nil {
		//	return err
		//}

		userCookie := models.UserCookie{
			Value:      cookie.Value,
			Expiration: cookie.Expires,
		}

		sess := functions.Session{
			UserID: uint(userSession.UserID),
			ID:     string(userSession.ID),
		}

		if ctx.Request().URL.Path != "/login" &&
			ctx.Request().URL.Path != "/registration" &&
			ctx.Request().URL.Path != "/logout" &&
			ctx.Request().Method != "GET" {

			CSRFToken := ctx.Request().Header.Get("csrf-token")
			_, err = tokens.Check(&sess, CSRFToken)
			if err != nil {
				return err
			}
		}

		token, err := tokens.Create(&sess, time.Now().Add(30*time.Minute).Unix())
		if err != nil {
			return err
		}

		ctx.Set("token", token)
		ctx.Set("User", user)
		ctx.Set("Cookie", userCookie)
		return next(ctx)
	}
}
