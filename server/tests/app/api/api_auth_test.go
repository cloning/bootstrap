package api

import (
	"../../../src/app/api"
	"regexp"
	"testing"
)

const (
	email, password, fullName = "testing", "testing", "Testy Testington"
	loginPostBody             = "{\"email\":\"" + email + "\", \"password\" : \"" + password + "\", \"fullName\" : \"" + fullName + "\"}"
	loginSuccess              = "{\"token\":\".*\",\"expires\":\".*\"}"
	loginFail                 = "{\"error\":\"Invalid credentials\"}"
	registerPostBody          = loginPostBody
	registerSuccess           = loginSuccess
)

func TestRegisterAndLogin(t *testing.T) {
	WithApi(func(api *api.Api) {

		// First - check that the login fails
		body, _, _ := Post("/auth/login", loginPostBody)

		if matched, err := regexp.Match(loginFail, body); !matched || err != nil {
			t.Errorf("Invalid response on failed login: %s %s", body, loginFail)
		}

		// Then we register the user
		body, _, _ = Post("/auth/register", registerPostBody)

		if matched, err := regexp.Match(registerSuccess, body); !matched || err != nil {
			t.Errorf("Invalid response on successful registration: %s %s", body, registerSuccess)
		}

		// First - check that the login fails
		body, _, _ = Post("/auth/login", loginPostBody)

		if matched, err := regexp.Match(loginSuccess, body); !matched || err != nil {
			t.Errorf("Invalid response on successful login: %s", body)
		}
	})
}

func TestRegisterWithInvalidData(t *testing.T) {
	//Invalid: missing fullName
	invalidPostBody := "{\"email\":\"some@email.com\",\"password\":\"password\"}"

	_, response, _ := Post("/auth/register", invalidPostBody)

	if response.StatusCode != 400 {
		t.Errorf("Expected http 400, but was %s", response.StatusCode)
	}

}

// TODO: Test login and get current user
// TODO: Test duplicate email

func TestValidateToken(t *testing.T) {
	WithApi(func(api *api.Api) {

		// Just a dummy method that allows the consumer to check the token
		// without doing anything.
		_, res, _ := GetWithToken("/auth/token/validate", "FAKE TOKEN")

		if res.StatusCode != 403 {
			t.Errorf("Expected status to be 403, but was %v", res.StatusCode)
		}

		//TODO: Verify that a real token is valid
	})
}
