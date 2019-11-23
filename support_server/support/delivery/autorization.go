package delivery

/*import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
)

func (h *HandlersStruct) HandleLoginUser(ctx echo.Context) (Err error) {
	var err error
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	if user := ctx.Get("User"); user != nil {
		data := h.PUsecase.SetJSONData(user.(models.User), ctx.Get("token").(string),"OK")
		err := encoder.Encode(data)
		if err != nil {
			return err
		}
		return nil
	}
	decoder := json.NewDecoder(ctx.Request().Body)
	newUserLogin := new(models.UserLogin)
	if err := decoder.Decode(newUserLogin); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	var User models.User
	User, err = h.PUsecase.GetUserByEmail(newUserLogin.Email)
	if err != nil {
		return err
	}

	if err := h.PUsecase.ComparePassword(User.Password, User.Salt, newUserLogin.Password); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	cookies, err := h.PUsecase.AddNewUserSession(User.ID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	ctx.SetCookie(&cookies)
	data := h.PUsecase.SetJSONData(User, "","OK")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}*/