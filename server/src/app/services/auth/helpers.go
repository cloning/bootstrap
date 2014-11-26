package auth

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/satori/go.uuid"
)

func hashPassword(password string) string {
	d := sha512.New()
	d.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(d.Sum(nil))
}

func generateTokenKey() string {
	return uuid.NewV4().String()
}
