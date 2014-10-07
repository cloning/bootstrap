package main

import (
	"./api"
	"./services"
)

func main() {
	service := services.NewService("Bootstrap Service")
	api := api.NewApi(service, 8080)
	api.Run()
}
