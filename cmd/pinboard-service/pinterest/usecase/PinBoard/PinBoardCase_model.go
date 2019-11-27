package PinBoard

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	"sync"
)

type PinBoardCase struct {
	PinBoardRep repository.PinBoard
	Sanitizer   sanitizer.SanitInterface
	Mu          *sync.Mutex
}

type IPinBoardCase interface {

}
