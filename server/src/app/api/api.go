package api

import (
	"github.com/go-martini/martini"
	"github.com/julianduniec/martini-jsonp"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"net/http"
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
	Middleware that checks that the user
	is authenticated, and passes user to the endpoint
*/
func (this *Api) authorized() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		//TODO: Email should be userId, DAMNIT
		isValid, email := this.authService.ValidateToken(r.Header.Get("Authorization"))
		if isValid == false {
			w.WriteHeader(403)
		} else {
			//TODO: Handle email
			user, _ := this.userService.FindFromEmail(email)
			c.Map(user)
		}
	}
}

/*
	Setup routing
*/
func (this *Api) route(m *martini.ClassicMartini) {

	/*
		Authentication
	*/

	m.Post("/auth/login", binding.Json(AuthLoginRequest{}), this.auth_login)

	m.Post("/auth/register", binding.Json(AuthRegisterRequest{}), this.auth_register)

	m.Get("/auth/token/validate", this.authorized(), this.auth_token_validate)

}
