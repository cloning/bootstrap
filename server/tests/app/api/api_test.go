package api_test

import (
	"../../../src/app/api"
	"regexp"
	"testing"
)

func TestJsonPShouldWrapResponse(t *testing.T) {
	withApi(func(api *api.Api) {
		body, _, _ := get("/?jsonp=callback")

		if matched, err := regexp.Match("callback(.*)", body); !matched || err != nil {
			t.Errorf("Not wrapped in callback:%s", body)
		}
	})
}
