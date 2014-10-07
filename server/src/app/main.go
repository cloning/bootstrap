package main

import (
	"./api"
	"./service"
)

func main() {
	service := service.NewService("Bootstrap Service")
	api := api.NewApi(service, 8080)
	api.Run()
}
