package jsonp

import (
	"github.com/go-martini/martini"
	"net/http"
)

type Options struct {
	ParameterName string
}

func JSONP(options Options) martini.Handler {

	if options.ParameterName == "" {
		options.ParameterName = "jsonp"
	}

	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {

		callback := r.URL.Query().Get(options.ParameterName)

		hasJsonp := callback != ""

		if hasJsonp {
			w.Write([]byte(callback + "("))
		}

		c.Next()

		if hasJsonp {
			w.Write([]byte(")"))
		}
	}
}
