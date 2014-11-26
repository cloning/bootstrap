package api

/*
   This file contains the authentication endpoints
   of the API, ie. anything under /auth
*/
import (
	"../services/user"
	"github.com/martini-contrib/render"
)

func (this *Api) auth_login(request AuthLoginRequest, r render.Render) {

	token, err := this.authService.Login(request.Email, request.Password)

	if err != nil {
		r.JSON(500, err)
	} else if token == nil {
		r.JSON(400, &ErrorResponse{"Invalid credentials"})
	} else {
		r.JSON(200, token)
	}
}

func (this *Api) auth_register(request AuthRegisterRequest, r render.Render) {

	if !request.Validate() {
		r.JSON(400, &ErrorResponse{"Invalid request"})
		return
	}

	// Create the user object
	_, err := this.userService.Create(request.Email, request.FullName)

	// If the email is already taken, we bail
	if _, ok := err.(user.EmailAlreadyExists); ok {
		r.JSON(400, &ErrorResponse{"Email already exists"})
		return
	}

	if err != nil {
		r.JSON(500, &ErrorResponse{err.Error()})
		return
	}

	// Register the auth credentials
	err = this.authService.Register(request.Email, request.Password)

	if err != nil {
		r.JSON(500, &ErrorResponse{err.Error()})
		return
	}

	// Create the login token and return
	token, err := this.authService.Login(request.Email, request.Password)

	if err != nil || token == nil {
		r.JSON(500, &ErrorResponse{err.Error()})
	} else {
		r.JSON(200, token)
	}
}

func (this *Api) auth_token_validate(user *user.User, r render.Render) {
	r.Status(200)
}
