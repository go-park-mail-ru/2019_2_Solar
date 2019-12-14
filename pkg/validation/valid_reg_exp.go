package validation

import (
	"regexp"
)

var (
	EmailIsCorrect    = regexp.MustCompile(`^.+@.+\..+$`)
	UsernameIsCorrect = regexp.MustCompile(`^[\w\d]{3,30}$`)

	PasswordIsCorrect       = regexp.MustCompile(`^.{8,30}$`)
	PasswordHasChar 		= regexp.MustCompile(`^.*[A-Za-z]+.*$`)
	PasswordHasNumber 		= regexp.MustCompile(`^.*[0-9]+.*$`)

	NameIsCorrect    = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	SurnameIsCorrect = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	StatusIsCorrect  = regexp.MustCompile(`^.*$`)
	AgeIsCorrect     = regexp.MustCompile(`^[0-9]{1,3}$`)

	BoardTitle       = regexp.MustCompile(`^.{3,50}$`)
	BoardDescription = regexp.MustCompile(`^.{0,300}$`)

	PinTitle       = regexp.MustCompile(`^.{3,30}$`)
	PinDescription = regexp.MustCompile(`^.{0,200}$`)

	FindTags = regexp.MustCompile(`#\w+`)

	FindJpg = regexp.MustCompile(`https:\/\/.{1,100}.jpg`)
)
