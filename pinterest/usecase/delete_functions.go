package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
)

func (USC *UsecaseStruct) DeleteOldUserSession(sessionKey string) error {
	var params []interface{}
	params = append(params, sessionKey)

	err := USC.PRepository.DeleteSession(consts.DeleteSessionByKey, params)
	if err != nil {
		return err
	}
	return nil
}