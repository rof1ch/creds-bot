package crypto

import (
	"crypto/sha256"
	"fmt"
)

func Hash(str string) string {
	data := []byte(str)
	hash := sha256.Sum256(data)
    return fmt.Sprintf("%x", hash)
}

func CheckHash(str string, oldHashStr string) bool {
    newData := []byte(str)
	newHash := sha256.Sum256(newData)
    newHashStr := fmt.Sprintf("%x", newHash)

    return newHashStr == oldHashStr
}