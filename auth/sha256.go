package auth

import (
	"crypto/sha256"
	"fmt"
)

func PasswordEncode(s string) string {
	str := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", str)
}