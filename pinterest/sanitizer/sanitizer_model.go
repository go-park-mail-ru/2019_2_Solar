package sanitizer

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/microcosm-cc/bluemonday"
)

type SanitStruct struct {
	sanit *bluemonday.Policy
}

type SanitInterface interface {
	SanitUser(user *models.User)
	SanitPin(pin *models.Pin)
	SanitPinForSearchResult(pin *models.PinForSearchResult)
	SanitComment(comment *models.CommentDisplay)
	SanitBoard(board *models.Board)
	SanitPinDisplay(pin *models.PinDisplay)
	SanitFullPin(pin *models.FullPin)
}
