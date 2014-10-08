package api_test

import (
	"../../../src/app/api"
	"../../../src/app/services"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"testing"
)

const (
	PORT = 4001
)

var apiUrl = fmt.Sprintf("http://localhost:%d", PORT)
var service = services.NewService("Test")

func TestJsonPShouldWrapResponse(t *testing.T) {
	var wg sync.WaitGroup

	api := api.NewApi(service, PORT, wg)

	go api.Run()

	body, _, _ := get("/?jsonp=callback")

	if matched, err := regexp.Match("callback(.*)", body); !matched || err != nil {
		t.Errorf("Not wrapped in callback:%s", body)
	}

	api.Stop()
	wg.Wait()
}

func get(path string) ([]byte, *http.Response, error) {
	resp, err := http.Get(apiUrl + path)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, resp, err
}
