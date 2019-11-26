package delivery

import (
	"context"
	"encoding/json"
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func (h *HandlersStruct) ServiceGetBoard(ctx echo.Context) (Err error) {
defer func() {
if err := ctx.Request().Body.Close(); err != nil {
Err = err
}
}()
ctx.Response().Header().Set("Content-Type", "application/json")
	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			Err = err
		}
	}()
	ctx.Response().Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Response())
	getUser := ctx.Get("User")
	if getUser == nil {
		return errors.New("not authorized")
	}
	//user := getUser.(models.User)

	id := ctx.Param("id")
	if id == "" {
		return errors.New("incorrect id")
	}
	boardID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	boardIDMessage := pinboard_service.BoardID{
		BoardID:              uint64(boardID),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	boardAndPins, err := h.PinBoardService.GetBoard(context.Background(), &boardIDMessage)
	if err != nil {
		return err
	}

	board := models.Board{
		ID:          boardAndPins.Board.ID,
		OwnerID:     boardAndPins.Board.OwnerID,
		Title:       boardAndPins.Board.Title,
		Description: boardAndPins.Board.Description,
		Category:    boardAndPins.Board.Category,
		CreatedTime: time.Now(),
		IsDeleted:   boardAndPins.Board.IsDeleted,
	}

	pins := []models.PinDisplay{}

	for _, element := range boardAndPins.Pins {
		pins = append(pins, models.PinDisplay{
			ID:          element.ID,
			PinDir:      element.PinDir,
			Title:       element.Title,
		})
	}

	data := struct {
		CSRFToken string `json:"csrf_token"`
		Body      struct {
			Board models.Board        `json:"board"`
			Pins  []models.PinDisplay `json:"pins"`
			Info  string              `json:"info"`
		} `json:"body"`
	}{CSRFToken: ctx.Get("token").(string),
		Body: struct {
			Board models.Board        `json:"board"`
			Pins  []models.PinDisplay `json:"pins"`
			Info  string              `json:"info"`
		}{Info: "OK", Board: board, Pins: pins}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) ServiceGetPin(ctx echo.Context) (Err error) {
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
	id := ctx.Param("id")
	if id == "" {
		return errors.New("incorrect id")
	}
	pinId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	pinIDMessage := pinboard_service.PinID{
		PinID:                uint64(pinId),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	pinAndComments, err := h.PinBoardService.GetPin(context.Background(), &pinIDMessage)
	if err != nil {
		return err
	}

	pin := models.FullPin{
		ID:             pinAndComments.Pin.ID,
		OwnerUsername:  pinAndComments.Pin.OwnerUsername,
		AuthorUsername: pinAndComments.Pin.AuthorUsername,
		BoardID:        pinAndComments.Pin.BoardID,
		PinDir:         pinAndComments.Pin.PinDir,
		Title:          pinAndComments.Pin.Title,
		Description:    pinAndComments.Pin.Description,
		CreatedTime:    time.Now(),
		IsDeleted:      false,
	}

	comments := []models.CommentDisplay{}

	for _, element := range 	pinAndComments.Comments {
		comments = append(comments, models.CommentDisplay{
			Text:          element.Text,
			CreatedTime:   time.Now(),
			Author:        element.Author,
			AuthorPicture: element.AuthorPincture,
		})
	}

	body := struct {
		Pin  models.FullPin `json:"pins"`
		Comments []models.CommentDisplay `json:"comments"`
		Info  string     `json:"info"`
	}{pin, comments ,"OK"}
	data := models.ValeraJSONResponse{ctx.Get("token").(string),body}
	if err := ctx.JSON(200, data); err != nil {
		return err
	}

	return nil
}