package middlewares

import (
	//"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/functions"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	"github.com/labstack/echo"
	"time"
)

func (MS *MiddlewareStruct) AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//tokens, _ := functions.NewAesCryptHashToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
		cookie, err := ctx.Cookie("session_key")
		if err != nil {
			return next(ctx)
		}

		user, err := MS.MUsecase.GetUserByCookieValue(cookie.Value)
		if err != nil {
			return err
		}

		userSession, err := MS.MUsecase.GetSessionsByCookieValue(cookie.Value)
		if err != nil {
			return err
		}

		userCookie := models.UserCookie{
			Value:      userSession.Value,
			Expiration: userSession.Expiration,
		}

		if userCookie.Expiration.Before(time.Now()) {
			//delete Coockie!!!!
			return next(ctx)
		}

		//sess := functions.Session{
		//	UserID: uint(userSession.UserID),
		//	ID:     string(userSession.ID),
		//}

		//if ctx.Request().URL.Path != "/login" &&
		//	ctx.Request().URL.Path != "/registration" &&
		//	ctx.Request().Method != "GET" {
		//
		//	CSRFToken := ctx.Request().Header.Get("csrf-token")
		//	_, err = tokens.Check(&sess, CSRFToken)
		//	if err != nil {
		//		return err
		//	}
		//}
		//
		//token, err := tokens.Create(&sess, time.Now().Add(30*time.Minute).Unix())
		//if err != nil {
		//	return err
		//}
		//
		//ctx.Set("token", token)

		ctx.Set("User", user)
		ctx.Set("Cookie", userCookie)
		return next(ctx)
	}
}


func (MS *MiddlewareStruct) AuthenticationAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//tokens, _ := functions.NewAesCryptHashToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
		cookie, err := ctx.Cookie("admin_session_key")
		if err != nil {
			return next(ctx)
		}

		admin, err := MS.MUsecase.GetAdminByCookieValue(cookie.Value)
		if err != nil {
			return err
		}

		adminSession, err := MS.MUsecase.GetAdminSessionsByCookieValue(cookie.Value)
		if err != nil {
			return err
		}

		adminCookie := models.UserCookie{
			Value:      adminSession.Value,
			Expiration: adminSession.Expiration,
		}

		if adminCookie.Expiration.Before(time.Now()) {
			//delete Coockie!!!!
			return next(ctx)
		}

		ctx.Set("Admin", admin)
		ctx.Set("Cookie", adminCookie)
		return next(ctx)
	}
}
