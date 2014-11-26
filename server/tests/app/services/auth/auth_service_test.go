package auth

import (
	"../../../../src/app/services/auth"
	"testing"
)

/*
	Asserts that valid accounts give valid tokens
*/
func TestAdminAccounts(t *testing.T) {
	createAuthService(t, func(authService *auth.AuthService) {

		t1 := verifyLogin("admin", "pwd", authService, t)
		t2 := verifyLogin("otheradmin", "pwd", authService, t)
		t3 := verifyLogin("non-existing", "non-existing", authService, t)

		if t1.Key == t2.Key {
			t.Log("Token keys are not unique")
			t.Fail()
		}

		verifyToken(t1, t)
		verifyToken(t2, t)

		if t3 != nil {
			t.Log("User should not be valid")
			t.Fail()
		}
	})
}

/*
	A valid token key should be valid
	and an invalid should now =)
*/
func TestTokenValidate(t *testing.T) {
	createAuthService(t, func(authService *auth.AuthService) {

		token := verifyLogin("admin", "pwd", authService, t)

		if valid, _ := authService.ValidateToken(token.Key); !valid {
			t.Log("Token should be valid")
			t.Fail()
		}

		token.Key = "invalid"

		if valid, _ := authService.ValidateToken(token.Key); valid {
			t.Log("Token shouldn't be valid")
			t.Fail()
		}
	})
}

/*
	Asserts that a token is persisted across instances of services
*/
func TestTokenPersistance(t *testing.T) {
	createAuthService(t, func(authService1 *auth.AuthService) {
		createAuthService(t, func(authService2 *auth.AuthService) {
			token := verifyLogin("admin", "pwd", authService1, t)

			verifyToken(token, t)

			validInFirst, _ := authService1.ValidateToken(token.Key)
			validInSecond, _ := authService2.ValidateToken(token.Key)

			if validInFirst == false {
				t.Log("Token was not valid in first service instance")
				t.Fail()
			}

			if validInSecond == false {
				t.Log("Token was not valid in first service instance")
				t.Fail()
			}
		})
	})
}

func TestHashPassword(t *testing.T) {
	createAuthService(t, func(authService *auth.AuthService) {
		expected := "7hBn0sVNiwlbt7OTeqQJaMw0deQ2BDOov4FiF+gjJx/MnnIi3Z5Xr7VnXZmbiPSVdO2OajgzsUN5EOmrp7ZXXw=="
		hashed := authService.HashPassword("pwd")
		if hashed != expected {
			t.Log("Hashed password did not match expected password")
			t.Fail()
		}
	})
}

func TestRegister(t *testing.T) {
	createAuthService(t, func(authService *auth.AuthService) {
		userName, pwd := "name", "password"
		authService.Register(userName, pwd)

		token := verifyLogin(userName, pwd, authService, t)

		verifyToken(token, t)

		valid, _ := authService.ValidateToken(token.Key)

		if valid == false {
			t.Log("The token that was created was not valid")
			t.Fail()
		}
	})
}
