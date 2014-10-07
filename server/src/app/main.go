package main

import (
	"./api"
	"./core"
)

func main() {
	service := core.NewService("Bootstrap Service")
	api := api.NewApi(service)
	api.Run()
}
