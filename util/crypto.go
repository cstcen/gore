package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func CFBEncrypt(text string, secret string) (string, error) {
	key := []byte(secret)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("new cipher, key: %s, error: %w", key, err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("io read full, iv: %s, error: %w", iv, err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return fmt.Sprintf("%x", ciphertext), nil
}

func CFBDecrypt(text string, secret string) (string, error) {
	key := []byte(secret)
	ciphertext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("new cipher, key: %s, error: %w", key, err)
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ECBDecrypt(ciphertext, key string) string {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		fmt.Printf("failed to base64 std encoding decode string, data: %s", data)
		return ""
	}
	block, _ := aes.NewCipher([]byte(key))

	plaintext := make([]byte, len(data))

	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		if len(plaintext) < be || len(data) < be {
			fmt.Printf("failed to ecb decrypt, `be` out array")
			break
		}
		block.Decrypt(plaintext[bs:be], data[bs:be])
	}

	return string(pkcs5UnPadding(plaintext))
}

func ECBEncrypt(plaintext, key string) string {
	data := []byte(plaintext)
	block, _ := aes.NewCipher([]byte(key))
	size := block.BlockSize()
	data = pkcs5Padding(data, size)
	decrypted := make([]byte, len(data))

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return base64.StdEncoding.EncodeToString(decrypted)
}
