package secutil

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func HashStr(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashBytes := h.Sum(nil)

	return fmt.Sprintf("%x", hashBytes)
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeBase64(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	return string(decodedBytes), err
}
