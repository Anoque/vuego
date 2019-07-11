package main

import (
	"./controllers"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return controllers.GetTests()
	})

	m.Group("/users", func(r martini.Router) {
		r.Post("/add", func(params martini.Params, req *http.Request) {
			controllers.CreateUser(req)
		})
	})

	m.Run()
}
