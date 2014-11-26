package api

import (
	"../../../src/app/api"
	"../../../src/app/services/auth"
	"../../../src/app/services/user"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

const (
	PORT         = 4001
	TEST_DB_HOST = "localhost"
	TEST_DB      = "intellizen-api-test"
)

type apifn func(*api.Api)

func Get(path string) ([]byte, *http.Response, error) {
	return GetWithToken(path, "")
}

func GetWithToken(path, token string) ([]byte, *http.Response, error) {
	return Request("GET", path, "", map[string]string{
		"Authorization": token,
	})
}

func Post(path, postbody string) ([]byte, *http.Response, error) {
	return PostWithToken(path, postbody, "")
}

func PostWithToken(path, postBody, token string) ([]byte, *http.Response, error) {
	return Request("POST", path, postBody, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	})
}

func Request(method, path, data string, headers map[string]string) ([]byte, *http.Response, error) {
	url := fmt.Sprintf("http://localhost:%d%s", PORT, path)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(data))

	if err != nil {
		return nil, nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, resp, err

}

func createAuthService() *auth.AuthService {
	authService, err := auth.NewAuthService("./mock/.admin-accounts", TEST_DB_HOST, TEST_DB)
	if err != nil {
		panic(err)
	}
	return authService
}

func createUserService() *user.UserService {
	userService, err := user.NewUserService(TEST_DB_HOST, TEST_DB)

	if err != nil {
		panic(err)
	}

	return userService
}

func WithApi(fn apifn) {
	var wg sync.WaitGroup

	defer cleanup()

	authService := createAuthService()
	userService := createUserService()

	//TODO: init company service
	api := api.NewApi(PORT, &wg, authService, userService)
	defer func() {
		userService.Close()
		api.Stop()
		wg.Wait()
	}()
	go api.Run()
	fn(api)
}
func cleanup() {
	s, _ := mgo.Dial(TEST_DB_HOST)
	s.DB(TEST_DB).DropDatabase()

}
