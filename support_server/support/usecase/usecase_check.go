package usecase

import "errors"

func (use *UseStruct) CompareAdminPassword(password string, autorizedPassword string) (Err error) {
	if password != autorizedPassword {
		return errors.New("not equal")
	}
	return nil
}
