package delivery

import (
	"context"
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"time"

	//"github.com/go-park-mail-ru/2019_2_Solar/cmd/authorization-service/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
)

type HandlersStruct struct {
	PUsecase usecase.UseInterface
}

type PinBoardService struct {
	Usecase 	usecase.UseInterface
	Host		string
}

func NewPinBoardService(use usecase.UseInterface,
	port string) *PinBoardService {
	return &PinBoardService{
		Usecase: use,
		Host: port,
	}
}

func (pnb *PinBoardService) CreateBoard(ctx context.Context, in *pinboard_service.NewBoard) (*pinboard_service.LastID, error) {

	newBoard := models.NewBoard{
		Title:      in.Title,
		Description: in.Description,
		Category:    in.Category,
	}

	if err := pnb.Usecase.CheckBoardData(newBoard); err != nil {
		return &pinboard_service.LastID{}, err
	}
	board := models.Board{
		OwnerID:     in.OwnerID,
		Title:       newBoard.Title,
		Description: newBoard.Description,
		Category:    newBoard.Category,
		CreatedTime: time.Now(),
	}
	lastID, err := pnb.Usecase.AddBoard(board)
	if err != nil {
		return &pinboard_service.LastID{}, err
	}

	
	return &pinboard_service.LastID{
		LastID:               lastID,
	}, nil
}

func (pnb *PinBoardService) GetBoard(ctx context.Context, in *pinboard_service.BoardID) (*pinboard_service.BoardAndPins, error) {




	board, err := pnb.Usecase.GetBoard(in.BoardID)
	if err != nil {
		return &pinboard_service.BoardAndPins{}, err
	}
	pins, err := pnb.Usecase.GetPinsDisplay(in.BoardID)
	if err != nil {
		return &pinboard_service.BoardAndPins{}, err
	}

	boardMessage :=  pinboard_service.Board{
		ID:                   board.ID,
		OwnerID:              board.OwnerID,
		Title:                board.Title,
		Description:          board.Description,
		Category:             board.Category,
		CreatedTime:          board.CreatedTime.String(),
		IsDeleted:            board.IsDeleted,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	var pinsMessage []*pinboard_service.BoardAndPins_Pin

	for _, element := range pins {
		pinsMessage = append(pinsMessage, &pinboard_service.BoardAndPins_Pin{
			ID:                   element.ID,
			PinDir:               element.PinDir,
			Title:                element.Title,
		})
	}



	return &pinboard_service.BoardAndPins{
		Board:                &boardMessage,
		Pins:                 pinsMessage,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

func (pnb *PinBoardService) CreatePin(ctx context.Context, in *pinboard_service.NewPin) (*pinboard_service.LastID, error) {

	newPin := models.NewPin{
		BoardID:     in.BoardID,
		Title:       in.Title,
		Description: in.Description,
		PinDir:      in.PinDir,
	}

	if err := pnb.Usecase.CheckPinData(newPin); err != nil {
		return &pinboard_service.LastID{}, err
	}
	pin := models.Pin{
		OwnerID:     in.UserID,
		AuthorID:    in.UserID,
		BoardID:     newPin.BoardID,
		Title:       newPin.Title,
		Description: newPin.Description,
		PinDir:      newPin.PinDir,
		CreatedTime: time.Now(),
	}
	lastID, err := pnb.Usecase.AddPin(pin)
	if err != nil {
		return &pinboard_service.LastID{}, err
	}
	err = pnb.Usecase.AddTags(pin.Description, lastID)
	if err != nil {
		return &pinboard_service.LastID{}, err
	}
	pin.ID = lastID
	pin.IsDeleted = false

	return &pinboard_service.LastID{
		LastID:               lastID,
	}, nil
}

func (pnb *PinBoardService) GetPin(ctx context.Context, in *pinboard_service.PinID) (*pinboard_service.PinAndComments, error) {


	pin, err := pnb.Usecase.GetPin(in.PinID)
	if err != nil {
		return &pinboard_service.PinAndComments{}, err
	}
	comments, err := pnb.Usecase.GetComments(in.PinID)
	if err != nil {
		return &pinboard_service.PinAndComments{}, err
	}


	pinMessage :=  pinboard_service.FullPin{
		ID:                   pin.ID,
		OwnerUsername:        pin.OwnerUsername,
		AuthorUsername:       pin.AuthorUsername,
		BoardID:              pin.BoardID,
		PinDir:               pin.PinDir,
		Title:                pin.Title,
		Description:          pin.Description,
		CratedTime:           pin.CreatedTime.String(),

	}

	var commentsMessage []*pinboard_service.PinAndComments_CommentDisplay

	for _, element := range comments {
		commentsMessage = append(commentsMessage, &pinboard_service.PinAndComments_CommentDisplay{
			Text:                element.Text,
			CreatedTime:          element.CreatedTime.String(),
			Author:               element.Author,
			AuthorPincture:       element.AuthorPicture,
		})
	}



	return &pinboard_service.PinAndComments{
		Pin:                &pinMessage,
		Comments:                 commentsMessage,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}