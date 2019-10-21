package usecase

import "errors"

func (USC *UsecaseStruct) ExtractFormatFile(fileName string) (string, error) {
	for i := 0; i < len(fileName); i++ {
		if string(fileName[i]) == "." {
			return fileName[i:], nil
		}
	}
	return "", errors.New("invalid file name")
}
