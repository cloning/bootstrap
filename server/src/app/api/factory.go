package api

import (
	"sync"
)

func NewApi(
	port int,
	wg *sync.WaitGroup) *Api {

	return &Api{
		port: port,
		wg:   wg,
		sl:   nil,
	}
}
