package api

/*
   Provides the core functionality of the
   api, such as running/serving http
*/

import (
	"../services/auth"
	"../services/user"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/hydrogen18/stoppableListener"
	"net"
	"net/http"
	"sync"
)

type Api struct {
	port        int
	wg          *sync.WaitGroup
	authService *auth.AuthService
	userService *user.UserService
	sl          *stoppableListener.StoppableListener
}

func (this *Api) Run() {
	m := martini.Classic()

	this.middleware(m)

	this.route(m)

	this.listenAndServe(m)
}

func (this *Api) Stop() {
	this.sl.Stop()
}

func (this *Api) listenAndServe(m *martini.ClassicMartini) {
	// Notify waitgroup that we have one task
	this.wg.Add(1)
	defer this.wg.Done()

	port := fmt.Sprintf(":%d", this.port)
	listener, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	this.sl, err = stoppableListener.New(listener)

	if err != nil {
		panic(err)
	}

	server := http.Server{}
	server.Handler = m
	server.Serve(this.sl)
}
