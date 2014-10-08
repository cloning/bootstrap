package main

import (
	"./api"
	"./configuration"
	"./services"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

var configurationFileFlag = flag.String("configurationFile", "", "Location of configuration file")

func main() {
	flag.Parse()

	// Load the configuration
	conf, err := loadConfiguration()

	if err != nil {
		panic(fmt.Sprintf("Could not load configuration: %s", err))
	}

	// Initialize any services here
	service := services.NewService("Bootstrap Service")

	// Waitgroup used for graceful cleanup on exit
	var wg sync.WaitGroup

	// Initialize the API
	api := api.NewApi(service, conf.Api.Port, wg)

	// Run the api
	go api.Run()

	// Block thread until OS interrupts
	blockUntilOsStop()

	// OS interrupted, stop running api
	api.Stop()

	// Wait for cleanup
	wg.Wait()
}

func blockUntilOsStop() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT)
	select {
	case signal := <-stop:
		fmt.Printf("Caught stop signal: %v", signal)
	}
}

/*
   Loads configuration
*/
func loadConfiguration() (*configuration.Configuration, error) {
	var configurationFile string

	// If no configuration file flag was set, we use the default
	if configurationFile = *configurationFileFlag; configurationFile == "" {
		var err error
		configurationFile, err = getDefaultConfiguration()

		// Unable to load configuration if we can't get access to any configuration file path
		if err != nil {
			return nil, err
		}
	}
	return configuration.Load(configurationFile)
}

/*
   Get the default configuration file (same directory as app)
*/
func getDefaultConfiguration() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir + "/configuration.yaml", err
}
