package validation

import (
	"regexp"
)

var (
	EmailIsCorrect    = regexp.MustCompile(`^.+@.+\..+$`)
	UsernameIsCorrect = regexp.MustCompile(`^[\w\d]{3,30}$`)

	PasswordIsCorrect       = regexp.MustCompile(`^[\w\d0-9]{8,30}$`)
	PasswordHasDownCaseChar = regexp.MustCompile(`^.*[a-z]+.*$`)
	PasswordHasAperCaseChar = regexp.MustCompile(`^.*[A-Z]+.*$`)

	NameIsCorrect    = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	SurnameIsCorrect = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	StatusIsCorrect  = regexp.MustCompile(`^.*$`)
	AgeIsCorrect     = regexp.MustCompile(`^[0-9]{1,3}$`)

	BoardTitle       = regexp.MustCompile(`^.{3,50}$`)
	BoardDescription = regexp.MustCompile(`^.{0,300}$`)

	PinTitle       = regexp.MustCompile(`^.{3,30}$`)
	PinDescription = regexp.MustCompile(`^.{0,200}$`)

	FindTags = regexp.MustCompile(`#\w+`)
)
