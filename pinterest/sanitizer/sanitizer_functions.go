package sanitizer

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/microcosm-cc/bluemonday"
)

func (San *SanitizerStruct) NewSanitizer() {
	San.sanitizer = bluemonday.UGCPolicy()
}

func (San *SanitizerStruct) SanitizeUser(user *models.User) {
	user.Name = San.sanitizer.Sanitize(user.Name)
	user.Surname = San.sanitizer.Sanitize(user.Surname)
	user.Username = San.sanitizer.Sanitize(user.Username)
}

func (San *SanitizerStruct) SanitizePin(pin *models.Pin) {
	pin.Description = San.sanitizer.Sanitize(pin.Description)
	pin.Title = San.sanitizer.Sanitize(pin.Title)
}

func (San *SanitizerStruct) SanitizeComment(comment *models.CommentForSend) {
	comment.Text = San.sanitizer.Sanitize(comment.Text)
}

func (San *SanitizerStruct) SanitizeBoard(board *models.Board) {
	board.Title = San.sanitizer.Sanitize(board.Title)
	board.Description = San.sanitizer.Sanitize(board.Description)
	board.Category = San.sanitizer.Sanitize(board.Category)
}
