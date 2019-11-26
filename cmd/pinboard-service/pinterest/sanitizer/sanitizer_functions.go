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
	user.Email = San.sanit.Sanitize(user.Email)
	user.Username = San.sanit.Sanitize(user.Username)
	user.Status = San.sanit.Sanitize(user.Status)
}

func (San *SanitStruct) SanitPin(pin *models.Pin) {
	pin.Description = San.sanit.Sanitize(pin.Description)
	pin.Title = San.sanit.Sanitize(pin.Title)
}

func (San *SanitStruct) SanitPinForSearchResult(pin *models.PinForSearchResult) {
	pin.Title = San.sanit.Sanitize(pin.Title)
}

func (San *SanitStruct) SanitComment(comment *models.CommentDisplay) {
	comment.Text = San.sanit.Sanitize(comment.Text)
	comment.Author = San.sanit.Sanitize(comment.Author)
}

func (San *SanitStruct) SanitBoard(board *models.Board) {
	board.Title = San.sanit.Sanitize(board.Title)
	board.Description = San.sanit.Sanitize(board.Description)
	board.Category = San.sanit.Sanitize(board.Category)
}

func (San *SanitStruct) SanitPinDisplay(pin *models.PinDisplay) {
	pin.Title = San.sanit.Sanitize(pin.Title)
}

func (San *SanitStruct) SanitFullPin(pin *models.FullPin) {
	pin.Title = San.sanit.Sanitize(pin.Title)
	pin.Description = San.sanit.Sanitize(pin.Description)
	pin.AuthorUsername = San.sanit.Sanitize(pin.AuthorUsername)
	pin.OwnerUsername = San.sanit.Sanitize(pin.OwnerUsername)
}
