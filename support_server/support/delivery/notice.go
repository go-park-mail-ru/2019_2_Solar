package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func (h *HandlersStruct) HandleCreateNotice(ctx echo.Context) (Err error) {
	defer func() {
		if bodyErr := ctx.Request().Body.Close(); bodyErr != nil {
			Err = errors.Wrap(Err, bodyErr.Error())
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	user := getUser.(models.User)

	decoder := json.NewDecoder(ctx.Request().Body)

	newNotice := new(models.NewNotice)

	if err := decoder.Decode(newNotice); err != nil {
		return err
	}

	id := ctx.Param("receiver_id")
	if id == "" {
		return errors.New("incorrect id")
	}
	receiverdID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	notice := models.Notice{
		UserID:      user.ID,
		ReceiverID:  uint64(receiverdID),
		Message:     newNotice.Message,
		CreatedTime: time.Now(),
	}
	lastID, err := h.PUsecase.AddNotice(notice)
	if err != nil {
		return err
	}
	notice.ID = lastID
	notice.IsRead = false

	data := struct {
		CSRFToken string `json:"csrf_token"`
		Body struct {
			Notice models.Notice `json:"notice"`
			Info   string        `json:"info"`
		} `json:"body"`
	}{ CSRFToken: ctx.Get("token").(string),
		Body: struct {
		Notice models.Notice `json:"notice"`
		Info   string        `json:"info"`
	}{Info: "data successfully saved", Notice: notice}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}


func (h *HandlersStruct) HandleGetNotices(ctx echo.Context) (Err error) {
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
	var notices []models.Notice
	notices, err := h.PUsecase.GetMyNotices(user.ID)
	if err != nil {
		return err
	}
	body := struct {
		Notices  []models.Notice `json:"notices"`
		Info  string     `json:"info"`
	}{notices, "OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}
	return nil
}