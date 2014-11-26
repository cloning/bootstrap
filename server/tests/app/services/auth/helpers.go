package auth

import (
	"../../../../src/app/services/auth"
	"gopkg.in/mgo.v2"
	"testing"
	"time"
)

const (
	DATABASE_HOST  = "localhost"
	DATABASE       = "intellizen-test-auth"
	AUTH_FILE_PATH = "./mock/.admin_accounts"
)

type AuthServiceCallback func(*auth.AuthService)

func verifyLogin(email, password string, authService *auth.AuthService, t *testing.T) *auth.Token {
	token, err := authService.Login(email, password)

	if err != nil {
		t.Error(err)
	}

	return token
}

func verifyToken(token *auth.Token, t *testing.T) {
	if token == nil {
		t.Log("Token was nil")
		t.Fail()
	}

	if token.Key == "" {
		t.Log("Token key was nil")
		t.Fail()
	}

	now := time.Now()

	if token.Expires.Before(now) || token.Expires.Equal(now) {
		t.Log("Token already expired")
		t.Fail()
	}
}

func createAuthService(t *testing.T, callback AuthServiceCallback) {

	authService, err := auth.NewAuthService(AUTH_FILE_PATH, DATABASE_HOST, DATABASE)

	if err != nil {
		t.Errorf("Coudn't create auth service %v", err)
	}
	defer cleanup()
	callback(authService)
}

func cleanup() {
	s, _ := mgo.Dial(DATABASE_HOST)
	s.DB(DATABASE).DropDatabase()
}
