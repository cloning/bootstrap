package configuration

import (
	"io/ioutil"
	"launchpad.net/goyaml"
)

/*
   Configuration related to the API
*/
type Api struct {
	Port int `yaml:"port"`
}

/*
   Root object for configuration file
*/
type Configuration struct {
	Api Api `yaml:"api"`
}

/*
   Loads configuration from the supplied path
   and parses the yaml into a Configuration struct
*/
func Load(path string) (*Configuration, error) {

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var c Configuration
	err = goyaml.Unmarshal(bytes, &c)

	return &c, err
}
