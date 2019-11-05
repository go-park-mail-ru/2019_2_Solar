package sanitizer

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	"github.com/microcosm-cc/bluemonday"
)

func (San *SanitStruct) NewSanitizer() {
	San.sanit = bluemonday.UGCPolicy()
}

func (San *SanitStruct) SanitUser(user *models.User) {
	user.Name = San.sanit.Sanitize(user.Name)
	user.Surname = San.sanit.Sanitize(user.Surname)
	user.Username = San.sanit.Sanitize(user.Username)
}

func (San *SanitStruct) SanitPin(pin *models.Pin) {
	pin.Description = San.sanit.Sanitize(pin.Description)
	pin.Title = San.sanit.Sanitize(pin.Title)
}

func (San *SanitStruct) SanitComment(comment *models.CommentForSend) {
	comment.Text = San.sanit.Sanitize(comment.Text)
}

func (San *SanitStruct) SanitBoard(board *models.Board) {
	board.Title = San.sanit.Sanitize(board.Title)
	board.Description = San.sanit.Sanitize(board.Description)
	board.Category = San.sanit.Sanitize(board.Category)
}
