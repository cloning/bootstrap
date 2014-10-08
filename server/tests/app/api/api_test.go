package api

import (
	"../../../src/app/api"
	"regexp"
	"testing"
)

func TestJsonPShouldWrapResponse(t *testing.T) {
	WithApi(func(api *api.Api) {
		body, _, _ := Get("/?jsonp=callback")

		if matched, err := regexp.Match("callback(.*)", body); !matched || err != nil {
			t.Errorf("Not wrapped in callback:%s", body)
		}
	})
}
