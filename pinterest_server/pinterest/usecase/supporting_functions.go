package usecase

import (
	"crypto/sha1"
	"errors"
	"golang.org/x/crypto/pbkdf2"
)

func (USC *UseStruct) ExtractFormatFile(fileName string) (string, error) {
	for i := 0; i < len(fileName); i++ {
		if string(fileName[i]) == "." {
			return fileName[i:], nil
		}
}
	return "", errors.New("invalid file name")
}

func HashPassword(password, salt string) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha1.New)
}
