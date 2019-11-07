package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func (h *HandlersStruct) HandleCreateBoard(ctx echo.Context) (Err error) {
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

	newBoard := new(models.NewBoard)

	if err := decoder.Decode(newBoard); err != nil {
		return err
	}

	if err := h.PUsecase.CheckBoardData(*newBoard); err != nil {
		return err
	}
	board := models.Board{
		OwnerID:     user.ID,
		Title:       newBoard.Title,
		Description: newBoard.Description,
		Category:    newBoard.Category,
		CreatedTime: time.Now(),
	}
	lastID, err := h.PUsecase.AddBoard(board)
	if err != nil {
		return err
	}
	board.ID = lastID
	board.IsDeleted = false

	data := struct {
		CSRFToken string `json:"csrf_token"`
		Body      struct {
			Board models.Board `json:"board"`
			Info  string       `json:"info"`
		} `json:"body"`
	}{CSRFToken: ctx.Get("token").(string),
		Body: struct {
			Board models.Board `json:"board"`
			Info  string       `json:"info"`
		}{Info: "data successfully saved", Board: board}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (h *HandlersStruct) HandleGetBoard(ctx echo.Context) (Err error) {
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

	board, err := h.PUsecase.GetBoard(uint64(boardID))
	if err != nil {
		return err
	}
	pins, err := h.PUsecase.GetPinsDisplay(uint64(boardID))
	if err != nil {
		return err
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

func (h *HandlersStruct) HandleGetMyBoards(ctx echo.Context) (Err error) {
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
	user := getUser.(models.User)

	boards, err := h.PUsecase.GetMyBoards(user.ID)
	if err != nil {
		return err
	}

	data := struct {
		CSRFToken string `json:"csrf_token"`
		Body      struct {
			Boards []models.Board `json:"boards"`
			Info   string         `json:"info"`
		} `json:"body"`
	}{CSRFToken: ctx.Get("token").(string),
		Body: struct {
			Boards []models.Board `json:"boards"`
			Info   string         `json:"info"`
		}{Info: "OK", Boards: boards}}

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
