package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
)

func (USC *UsecaseStruct) RemoveOldUserSession(sessionKey string) error {
	var params []interface{}
	params = append(params, sessionKey)

	err := USC.PRepository.DeleteSession(consts.DELETESessionByKey, params)
	if err != nil {
		return err
	}
	return nil
}

func (USC *UsecaseStruct) RemoveSubscribe(userID, followeeName string) error {
	var params []interface{}
	params = append(params, userID, followeeName)

	err := USC.PRepository.DeleteSubscribe(consts.DELETESubscribeByName, params)
	if err != nil {
		return err
	}
	return nil
}
