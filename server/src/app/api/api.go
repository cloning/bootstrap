package api

import (
	"../service"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

type Api struct {
	service *core.Service
	port    int
}

func NewApi(service *core.Service, port int) *Api {
	return &Api{
		service,
		port,
	}
}

func (this *Api) Run() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Charset: "UTF-8",
	}))

	m.Get("/", func(args martini.Params, r render.Render) {
		user := this.service.GetUser()
		r.JSON(200, user)
	})

	log.Fatal(
		http.ListenAndServe(
			// Listen and serve expects eg. ':8080'
			fmt.Sprintf(":%d", this.port),
			m))
}
