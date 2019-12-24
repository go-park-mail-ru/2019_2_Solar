package functions

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"time"
)

type CryptToken struct {
	Secret []byte
}

type TokenData struct {
	SessionID string
	UserID    uint
	Exp       int64
}

type Session struct {
	UserID uint
	ID     string
}

func NewAesCryptHashToken(secret string) (*CryptToken, error) {
	key := []byte(secret)
	_, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("cypher problem %v", err)
	}
	return &CryptToken{Secret: key}, nil
}

func (tk *CryptToken) Create(s *Session, tokenExpTime int64) (string, error) {
	block, err := aes.NewCipher(tk.Secret)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	td := &TokenData{SessionID: s.ID, UserID: s.UserID, Exp: tokenExpTime}
	data, _ := json.Marshal(td)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	res := append([]byte(nil), nonce...)
	res = append(res, ciphertext...)

	token := base64.StdEncoding.EncodeToString(res)
	return token, nil
}

func (tk *CryptToken) Check(s *Session, inputToken string) (bool, error) {
	block, err := aes.NewCipher(tk.Secret)
	if err != nil {
		return false, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(inputToken)
	if err != nil {
		return false, err
	}
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return false, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, fmt.Errorf("decrypt fail: %v", err)
	}

	td := TokenData{}
	err = json.Unmarshal(plaintext, &td)
	if err != nil {
		return false, fmt.Errorf("bad json: %v", err)
	}

	if td.Exp < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	expected := TokenData{SessionID: s.ID, UserID: s.UserID}
	td.Exp = 0
	return td == expected, nil
}


func GenSessionKey(length int) (string, error) {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			var err error
			randomBytes, err = SecureRandomBytes(bufferSize)
			if err != nil {
				return "", err
			}
		}
		if idx := int(randomBytes[j%length] & consts.LetterIdxMask); idx < len(consts.LetterBytes) {
			result[i] = consts.LetterBytes[idx]
			i++
		}
	}

	return string(result), nil
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) ([]byte, error) {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return []byte(""), err
	}
	return randomBytes, nil
}

func HashPassword(password, salt string) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha1.New)
}