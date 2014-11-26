package user

import (
	"../../../../src/app/services/user"
	"testing"
)

const (
	email    = "test@test.com"
	fullName = "Bengt Nilsson"
)

// Create user
func TestCreateUser(t *testing.T) {
	withUserService(t, func(userService *user.UserService) {
		user, err := userService.Create(email, fullName)

		if err != nil {
			t.Errorf("Error while creating user %v", err)
			return
		}

		if user == nil {
			t.Errorf("Created user was nil")
			return
		}

		if user.Email != email {
			t.Errorf("User email was incorrect %s", user.Email)
		}

		if user.FullName != fullName {
			t.Errorf("User fullname was incorrect %s", user.FullName)
		}
	})
}

func TestDuplicateEmail(t *testing.T) {
	withUserService(t, func(userService *user.UserService) {
		_, err := userService.Create(email, fullName)

		if err != nil {
			t.Errorf("Unexpected error %v", err)
			return
		}

		_, err = userService.Create(email, fullName)

		if _, ok := err.(user.EmailAlreadyExists); !ok {
			t.Errorf("Error expected to be of typ EmailAlreadyExists")
		}
	})
}

// Create user - duplicate email

// Find user by email

// Update single property
