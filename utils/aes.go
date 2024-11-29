package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const (
	TgidAesKey = "bd53ac8639784e4a89b64da02a045f53"
)

func AESEncrypt(plaintext, key string) (string, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AESDecrypt(ciphertext, key string) (string, error) {
	keyBytes := []byte(key)
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	plaintext := make([]byte, len(ciphertextBytes))
	stream.XORKeyStream(plaintext, ciphertextBytes)

	return string(plaintext), nil
}
