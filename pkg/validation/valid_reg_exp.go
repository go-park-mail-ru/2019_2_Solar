package validation

import (
	"regexp"
)

var (
	EmailIsCorrect    = regexp.MustCompile(`^.+@.+\..+$`)
	UsernameIsCorrect = regexp.MustCompile(`^[\w\d]{3,30}$`)

	PasswordIsCorrect       = regexp.MustCompile(`^.{6,30}$`)		// is not using
	PasswordHasChar 		= regexp.MustCompile(`^.*[A-Za-z]+.*$`) // is not using
	PasswordHasNumber 		= regexp.MustCompile(`^.*[0-9]+.*$`)	// is not using

	NameIsCorrect    = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	SurnameIsCorrect = regexp.MustCompile(`^[^\d_!@#$%^&*,.:~|\\\/\<\>=\+\?"'\[\]\{\}]*$`)
	StatusIsCorrect  = regexp.MustCompile(`^.*$`)
	AgeIsCorrect     = regexp.MustCompile(`^[0-9]{1,3}$`)

	BoardTitle       = regexp.MustCompile(`^.{3,50}$`)
	BoardDescription = regexp.MustCompile(`^.{0,300}$`)

	PinTitle       = regexp.MustCompile(`^.{3,30}$`)
	PinDescription = regexp.MustCompile(`^.{0,200}$`)

	FindTags = regexp.MustCompile(`#\w+`)

	//FindJpg = regexp.MustCompile(`(image_large_url|image_cover_hd_url)\\":\\"https:.{1,100}.jpg`)
	FindJpg = regexp.MustCompile(`(orig.{1,150}.jpg|((image_large_url|image_cover_hd_url)\\":\\"https:.{1,100}.jpg)|3x, https:.{1,100}.jpg)`)
	FindPinUrl = regexp.MustCompile(`https:.{1,100}.jpg`)
)
