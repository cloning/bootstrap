package api

import (
	"../services"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/hydrogen18/stoppableListener"
	"github.com/julianduniec/martini-jsonp"
	"github.com/martini-contrib/render"
	"net"
	"net/http"
	"sync"
)

type Api struct {
	service *services.Service
	port    int
	wg      sync.WaitGroup
	sl      *stoppableListener.StoppableListener
}

func NewApi(service *services.Service, port int, wg sync.WaitGroup) *Api {
	return &Api{
		service,
		port,
		wg,
		nil,
	}
}

func (this *Api) Run() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Charset: "UTF-8",
	}))

	m.Use(jsonp.JSONP(jsonp.Options{
		ParameterName: "jsonp",
	}))

	m.Get("/", func(args martini.Params, r render.Render) {
		user := this.service.GetUser()
		r.JSON(200, user)
	})

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

	http.Handle("/", m)
	server := http.Server{}
	server.Serve(this.sl)
}
