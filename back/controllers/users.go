package controllers

import (
	"log"
	"io"
	"strings"
	"net/http"
	"encoding/json"
	"../structs"
	"../utils"
)

func CreateUser(req *http.Request) bool {
	err := req.ParseForm()
	if err != nil {
		log.Fatal("Parse form error")
	}

	v := req.Form
	status := false

	// {"{key_is_json_from_client":null}
	for key := range v {
		var user structs.User
		dec := json.NewDecoder(strings.NewReader(key))

		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		user.Password = utils.GetEncodeString(user.Password)
		query := "INSERT INTO users (username, password) VALUES ('" + user.Username + "', '" + user.Password + "')"
		status = utils.RunQuery(query)

		break
	}

	return status
}

func AuthUser(req *http.Request) structs.User {
	err := req.ParseForm()
	if err != nil {
		log.Fatal("Parse form error")
	}
	v := req.Form
	User := structs.User{}

	for key := range v {
		var user structs.User
		dec := json.NewDecoder(strings.NewReader(key))

		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		log.Println(user)

		status := utils.UserIsAuthenticated(user.Username, user.Password)

		if status {
			User = utils.GetUser(user.Username, user.Password)
			log.Println(User.Id)
		} else {
			log.Println("User not found")
		}

		break
	}

	return User
}

