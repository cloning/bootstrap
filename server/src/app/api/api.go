package api

import (
	"../services"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/julianduniec/martini-jsonp"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

type Api struct {
	service *services.Service
	port    int
}

func NewApi(service *services.Service, port int) *Api {
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

	m.Use(jsonp.JSONP(jsonp.Options{
		ParameterName: "jsonp",
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
