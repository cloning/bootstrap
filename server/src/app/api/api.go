package api

import (
	"../service"
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

type Api struct {
	service *core.Service
	port    string
}

func NewApi(service *core.Service, port int) *Api {
	return &Api{
		service,
		fmt.Sprintf(":%d", port),
	}
}

func (this *Api) Run() {
	m := martini.Classic()

	m.Get("/", func() string {
		return this.service.GetName()
	})

	log.Fatal(http.ListenAndServe(this.port, m))
}
