package validation

import (
	"regexp"
)

var (
	EmailIsCorrect    = regexp.MustCompile(`^\w+@\w+\.\w+$`)
	UsernameIsCorrect = regexp.MustCompile(`^[\w\d]*$`)

	PasswordIsCorrect       = regexp.MustCompile(`^[\w\d!?_#&^%]{8,30}$`)
	PasswordHasDownCaseChar = regexp.MustCompile(`^.*[a-z]+.*$`)
	PasswordHasAperCaseChar = regexp.MustCompile(`^.*[A-Z]+.*$`)
	PasswordHasSpecChar     = regexp.MustCompile(`^.*[!?_#&^%]+.*$`)

	NameIsCorrect    = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	SurnameIsCorrect = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	StatusIsCorrect  = regexp.MustCompile(`^.*$`)
	AgeIsCorrect     = regexp.MustCompile(`^[0-9]{1,3}$`)
)
