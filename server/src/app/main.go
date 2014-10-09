package main

import (
	"flag"
)

var configurationFileFlag = flag.String("configurationFile", "", "Location of configuration file")

func main() {
	flag.Parse()
	app, err := NewApp(*configurationFileFlag)

	if err != nil {
		panic(err)
	}

	app.Run()
}
