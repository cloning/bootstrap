package main

import (
	"./api"
)

func main() {
	api := api.NewApi()
	api.Run()
}
