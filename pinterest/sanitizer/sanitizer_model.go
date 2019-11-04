package sanitizer

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/microcosm-cc/bluemonday"
)

type SanitizerStruct struct {
	sanitizer *bluemonday.Policy
}

type SanitizerInterface interface {
	SanitizeUser(user *models.User)
	SanitizePin(pin *models.Pin)
	SanitizeComment(comment *models.CommentForSend)
	SanitizeBoard(board *models.Board)
}
