package api

import (
	"../core"
	"github.com/go-martini/martini"
)

type Api struct {
	service *core.Service
}

func NewApi(service *core.Service) *Api {
	return &Api{service}
}

func (this *Api) Run() {
	m := martini.Classic()
	m.Get("/", func() string {
		return this.service.GetName()
	})
	m.Run()
}
