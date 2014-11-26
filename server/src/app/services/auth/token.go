package auth

import (
	"time"
)

type Token struct {
	Email   string    `json:"-" bson:"email"`
	Key     string    `json:"token"`
	Expires time.Time `json:"expires"`
}
