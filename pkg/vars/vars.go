package vars

import (
	"regexp"
)

var (
	emailIsCorrect    = regexp.MustCompile(`^\w+@\w+\.\w+$`)
	usernameIsCorrect = regexp.MustCompile(`^[\w\d]*$`)

	passwordIsCorrect       = regexp.MustCompile(`^[\w\d!?_#&^%]{8,30}$`)
	passwordHasDownCaseChar = regexp.MustCompile(`^.*[a-z]+.*$`)
	passwordHasAperCaseChar = regexp.MustCompile(`^.*[A-Z]+.*$`)
	passwordHasSpecChar     = regexp.MustCompile(`^.*[!?_#&^%]+.*$`)

	nameIsCorrect    = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	surnameIsCorrect = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	statusIsCorrect  = regexp.MustCompile(`^.*$`)
	ageIsCorrect     = regexp.MustCompile(`^[0-9]{1,3}$`)
)
