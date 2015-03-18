package api

import (
	"github.com/go-martini/martini"
	"github.com/julianduniec/martini-jsonp"
	//"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	//"net/http"
)

/*
	Setup middleware
*/
func (this *Api) middleware(m *martini.ClassicMartini) {

	// Use render rendering-engine
	m.Use(render.Renderer(render.Options{
		Charset: "UTF-8",
	}))

	// Allow JSONP
	m.Use(jsonp.JSONP(jsonp.Options{
		ParameterName: "jsonp",
	}))

	// Allow cross origin post
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))
}

/*
	Setup routing
*/
func (this *Api) route(m *martini.ClassicMartini) {

}
