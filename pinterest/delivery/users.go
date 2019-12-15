package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
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

	body := struct {
		User models.AnotherUser  `json:"user"`
		Pins []models.PinDisplay `json:"pins"`
		Info string              `json:"info"`
	}{userProfile, pins, "OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string), body}
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

func (h *HandlersStruct) HandleGetFolloweeUser(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	user := getUser.(models.User)
	ctx.Response().Header().Set("Content-Type", "application/json")

	followeeUsers, err := h.PUsecase.GetFolloweeUserBySubscriberId(user.ID)
	if err != nil {
		return err
	}
	data := models.ValeraJSONResponse{ctx.Get("token").(string), followeeUsers}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}

func (h *HandlersStruct) HandleFeedback(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	user := getUser.(models.User)
	ctx.Response().Header().Set("Content-Type", "application/json")
	newFeedBack := models.NewFeedBack{}
	if err := ctx.Bind(&newFeedBack); err != nil {
		ctx.Logger().Warn(err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	newFeedBack.UserId = user.ID
	if err:= h.PUsecase.AddFeedBack(newFeedBack); err != nil {
		ctx.Logger().Warn(err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	data := models.ValeraJSONResponse{ctx.Get("token").(string), "ok"}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}
