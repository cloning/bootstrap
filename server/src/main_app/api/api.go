package api

import (
	"github.com/go-martini/martini"
)

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

func (this *Api) Run() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
