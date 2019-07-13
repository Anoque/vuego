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

		log.Println(user.Username, user.Password)

		query := "INSERT INTO users (username, password) VALUES ('" + user.Username + "', '" + user.Password + "')"

		status = utils.RunQuery(query)

	}


	return status
}