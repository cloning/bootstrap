package main

import (
	"./api"
	"./configuration"
	"./services/auth"
	"./services/user"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

type App struct {
	Configuration *configuration.Configuration
	Api           *api.Api
	Wg            sync.WaitGroup
}

func NewApp(configurationFile string) (*App, error) {
	// Load the configuration
	conf, err := loadConfiguration(configurationFile)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not load configuration: %s", err))
	}

	// Waitgroup used for graceful cleanup on exit
	var wg sync.WaitGroup

	// Initialize any services here
	authService := createAuthService(conf)
	userService := createUserService(conf)

	// Initialize the API
	api := api.NewApi(conf.Api.Port, &wg, authService, userService)

	app := &App{
		Configuration: conf,
		Api:           api,
		Wg:            wg,
	}

	return app, nil
}

func createUserService(conf *configuration.Configuration) *user.UserService {
	userService, err := user.NewUserService(conf.UserService.DatabaseHost, conf.UserService.Database)

	if err != nil {
		panic(err)
	}

	return userService
}

func createAuthService(conf *configuration.Configuration) *auth.AuthService {
	appLocation, err := getAppLocation()

	if err != nil {
		panic(err)
	}

	adminFile := appLocation + "/" + conf.AuthService.AccountsFile

	authService, err := auth.NewAuthService(adminFile, conf.AuthService.DatabaseHost, conf.AuthService.Database)

	if err != nil {
		panic(err)
	}

	return authService

}

func (this *App) Run() {
	go this.Api.Run()

	this.blockUntilOsStop()

	// OS interrupted, stop running api
	this.Api.Stop()

	// Wait for cleanup
	this.Wg.Wait()
}

func (this *App) blockUntilOsStop() {
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
func loadConfiguration(configurationFile string) (*configuration.Configuration, error) {

	// If no configuration file flag was set, we use the default
	if configurationFile == "" {
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
	dir, err := getAppLocation()
	return dir + "/configuration.yaml", err
}

func getAppLocation() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
