package usecase

func (USC *UseStruct) RemoveOldUserSession(sessionKey string) error {
	err := USC.PRepository.DeleteSessionByKey(sessionKey)
	if err != nil {
		return err
	}
	return nil
}

