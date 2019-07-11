package main

import (
	"./controllers"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
	}))

	m.Get("/", func() string {
		return controllers.GetTests()
	})

	m.Group("/users", func(r martini.Router) {
		r.Post("/add", func(params martini.Params, req *http.Request, res http.ResponseWriter) string {
			controllers.CreateUser(req)

			res.WriteHeader(200)

			return "OK"
		})
	})

	m.Run()
}
