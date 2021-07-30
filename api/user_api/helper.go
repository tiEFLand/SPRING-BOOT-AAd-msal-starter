package user_api

import (
	"crypto/sha256"
	"encoding/hex"
)

func checkPassword(passwordPlainText, passwordHashed, salt string) bool {
	h := sha256.New()
	h.Write([]byte(passwordPlainText + salt))
	if hex.EncodeToString(h.Sum(nil)) == passwordHashed {
		return true
	}
	return false
}
