package usecase

func (USC *UseStruct) RemoveOldUserSession(sessionKey string) error {
	err := USC.PRepository.DeleteSessionByKey(sessionKey)
	if err != nil {
		return err
	}
	return nil
}

func (USC *UseStruct) RemoveSubscribe(userID uint64, followeeName string) error {
	err := USC.PRepository.DeleteSubscribeByName(userID, followeeName)
	if err != nil {
		return err
	}
	return nil
}
