package delivery

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	//"github.com/go-park-mail-ru/2019_2_Solar/cmd/authorization-service/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"net/http"
)

type HandlersStruct struct {
	PUsecase usecase.UseInterface
}

type AuthorizationService struct {
	Usecase 	usecase.UseInterface
	Host		string
}

func NewAuthorizationService(use usecase.UseInterface,
	port string) *AuthorizationService {
	return &AuthorizationService{
		Usecase: use,
		Host: port,
	}
}

func (auth *AuthorizationService) CheckSession(ctx context.Context, cookie *services.Cookie) (*services.UserSession, error) {


	value := ctx.Value( "session-key")
	serviceCookie, ok := value.(http.Cookie)
	if !ok {
		return &services.UserSession{}, errors.New("session_key is nil")
	}


	//user, err := auth.Usecase.GetUserByCookieValue(serviceCookie.Value)
	//if err != nil {
	//	return err
	//}

	userSession, err := auth.Usecase.GetSessionsByCookieValue(serviceCookie.Value)
	if err != nil {
		return &services.UserSession{}, err
	}

	serviceCookie.Expires = userSession.Expiration

	ctx = context.WithValue(ctx, "session-key", serviceCookie)

	return &services.UserSession{
		ID:                   userSession.ID,
		UserID:               userSession.UserID,
		Value:                userSession.Value,
		Exp:                  userSession.Expiration.String(),

	}, nil
}

func (auth *AuthorizationService) RegUser(ctx context.Context, in *services.UserReg) (*services.Cookie, error) {

	//user := ctx.Value("userReg")
	//newUserReg, ok := user.(models.UserReg)
	//if !ok {
	//	return &services.Cookie{}, errors.New("can not convert userReg to models.UserReg")
	//}

	newUserReg := models.UserReg{
		Email:   in.Email,
		Password: in.Password,
		Username: in.Username,
	}

	if err := auth.Usecase.CheckRegDataValidation(&newUserReg); err != nil {
		return &services.Cookie{}, &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	if err := auth.Usecase.CheckRegUsernameEmailIsUnique(newUserReg.Username, newUserReg.Email); err != nil {
		return &services.Cookie{}, err
	}

	newUserID, err := auth.Usecase.AddNewUser(newUserReg.Username, newUserReg.Email, newUserReg.Password)
	if err != nil {
		return &services.Cookie{}, err
	}

	cookies, err := auth.Usecase.AddNewUserSession(newUserID)
	if err != nil {
		return &services.Cookie{}, &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	return &services.Cookie{
		Key:                  cookies.Name,
		Value:                cookies.Value,
		Exp:                  cookies.Expires.String(),
	}, nil
}

func (auth *AuthorizationService) LoginUser(ctx context.Context, userLogin *services.UserLogin) (*services.Cookie, error) {

	return &services.Cookie{
		Key:                  "key",
		Value:                "123",
		Exp:                  "23:00",
	}, nil
}

func (auth *AuthorizationService) LogoutUser(ctx context.Context, cookie *services.Cookie) (*services.Nothing, error) {

	return &services.Nothing{
		Dummy:                false,
	}, nil
}