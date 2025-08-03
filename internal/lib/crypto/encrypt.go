package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

/*
    Input: 
        - Text to encrypt
        - Key
    Return:
        - Enctyption text
        - Nonce
        - Error
*/
func Encrypt(plaintext, key string) (string, string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	ciphertext := aesgsm.Seal(nil, nonce, []byte(plaintext), nil)

	return hex.EncodeToString(ciphertext), hex.EncodeToString(nonce), nil
}

func Decrypt(ciphertextStr, nonce, key string) (string, error) {
	ciphertext, err := hex.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}
	noncetext, err := hex.DecodeString(nonce)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesgsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgsm.Open(nil, noncetext, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), err
}
