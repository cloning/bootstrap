package api

import (
	"../services/auth"
	"../services/user"
	"sync"
)

func NewApi(
	port int,
	wg *sync.WaitGroup,
	authService *auth.AuthService,
	userService *user.UserService) *Api {

	return &Api{
		port:        port,
		wg:          wg,
		authService: authService,
		userService: userService,
		sl:          nil,
	}
}
