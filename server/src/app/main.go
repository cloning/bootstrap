package main

import (
	"./api"
	"./service"
)

func main() {
	service := core.NewService("Bootstrap Service")
	api := api.NewApi(service, 8080)
	api.Run()
}
