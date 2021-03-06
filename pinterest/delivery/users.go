package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (h *HandlersStruct) HandleGetUserByUsername(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	username := ctx.Param("username")
	if username == "" {
		return errors.New("incorrect name")
	}
	isFollowee := false
	getUser := ctx.Get("User")
	if getUser != nil {
		user := getUser.(models.User)
		var err error
		isFollowee, err = h.PUsecase.GetMySubscribeByUsername(user.ID, username)
		if err != nil {
			return err
		}
	}
	ctx.Response().Header().Set("Content-Type", "application/json")

	userProfile, err := h.PUsecase.GetUserByUsername(username)
	if err != nil {
		return err
	}
	userProfile.IsFollowee = isFollowee

	pins, err := h.PUsecase.GetPinsByUsername(int(userProfile.ID))
	if err != nil {
		return err
	}


	//data := h.PUsecase.SetJSONData(userProfile, ctx.Get("token").(string), "OK")
	//err = encoder.Encode(data)
	//if err != nil {
	//	return err
	//}
	//return nil

	body := struct {
		User models.AnotherUser `json:"user"`
		Pins  []models.PinDisplay `json:"pins"`
		Info  string     `json:"info"`
	}{userProfile, pins, "OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleCreateSubscribe(ctx echo.Context) (Err error) {
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
	user := getUser.(models.User)
	followeeName := ctx.Param("username")
	if err := h.PUsecase.AddSubscribe(user.ID, followeeName); err != nil {
		return err
	}
	body := struct {
		Info string `json:"info"`
	}{"data successfully saved"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string), body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleDeleteSubscribe(ctx echo.Context) (Err error) {
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
	user := getUser.(models.User)
	followeeName := ctx.Param("username")
	if err := h.PUsecase.RemoveSubscribe(user.ID, followeeName); err != nil {
		return err
	}
	body := struct {
		Info string `json:"info"`
	}{"data successfully saved"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string), body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}
