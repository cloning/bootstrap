martini-jsonp
=============

Middleware for github.com/go-martini/martini that wraps responses with JSONP callback if the query is present.  

Example  

          m := martini.Classic()

          m.Use(jsonp.JSONP(jsonp.Options{
                ParameterName: "jsonp",
          }))
