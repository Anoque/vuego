package controllers

import (
	"log"
	"net/http"
)

func CreateUser(req *http.Request) {
	var username string
	var password string
	username = req.FormValue("username")
	password = req.FormValue("password")

	log.Println(username, password)

	/*if len(username) > 0 && len(password) > 0 {
		log.Println("Yes")
	} else {
		panic("Not all parameters")
	}*/

}