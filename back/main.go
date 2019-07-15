package main

import (
	"./controllers"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"net/http"
	"./utils"
	"./structs"
)

func main() {
	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT", "PATCH"},
		AllowHeaders: []string{"X-Auth-Key", "X-Auth-Secret", "Content-Type"},
	}))

	m.Get("/", func() string {
		return controllers.GetTests()
	})

	m.Group("/users", func(r martini.Router) {
		r.Post("/add", func(params martini.Params, req *http.Request, res http.ResponseWriter) string {
			var status bool
			status = controllers.CreateUser(req)

			return utils.SendStatus(status)
		})
		r.Post("/auth", func(params martini.Params, req *http.Request, res http.ResponseWriter) string {
			var user structs.User
			user = controllers.AuthUser(req)

			if user.Id > 0 {
				return "{\"result\":\"" + user.Id + "\"}"
			} else {
				return utils.SendStatus(false)
			}

		})
	})

	m.Run()
}
