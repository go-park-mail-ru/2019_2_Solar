package sanitizer

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/microcosm-cc/bluemonday"
)

func (San *SanitizerStruct) NewSanitizer() {
	San.sanitizer = bluemonday.UGCPolicy()
}

func (San *SanitizerStruct) SanitizeUser(user models.User) models.User {
	user.Name = San.sanitizer.Sanitize(user.Name)
	user.Surname = San.sanitizer.Sanitize(user.Surname)
	user.Username = San.sanitizer.Sanitize(user.Username)
	return user
}

func (San *SanitizerStruct) SanitizePin(pin models.Pin) models.Pin {
	pin.Description = San.sanitizer.Sanitize(pin.Description)
	pin.Title = San.sanitizer.Sanitize(pin.Title)
	return pin
}

func (San *SanitizerStruct) SanitizeComment(comment models.Comment) models.Comment {
	comment.Text = San.sanitizer.Sanitize(comment.Text)
	return comment
}

func (San *SanitizerStruct) SanitizeBoard(board models.Board) models.Board {
	board.Title = San.sanitizer.Sanitize(board.Title)
	board.Description = San.sanitizer.Sanitize(board.Description)
	board.Category = San.sanitizer.Sanitize(board.Category)
	return board
}
