package utils

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	//plaintext := "Hello, AES-CTR encryption!"
	key := "1234567890123456"

	//encrypted, err := encrypt(plaintext, key)
	//if err != nil {
	//	fmt.Println("Error encrypting:", err)
	//	return
	//}
	//fmt.Println("Encrypted:", encrypted)

	decrypted, err := AESDecrypt("kND0JKhLLNX3yHmRY8wOH9lQfc65I+uTU0Y=", key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted:", decrypted)
}
