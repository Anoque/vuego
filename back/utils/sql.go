package utils

import (
	"database/sql"
	"log"
	"../structs"
	_ "github.com/go-sql-driver/mysql"
)

var db_user = "hello"
var db_password = "root"
var db_name = "vuego"

func RunQuery(query string) bool {
	db, err := sql.Open("mysql", db_user + ":" + db_password + "@/" + db_name)
	defer db.Close()

	if err != nil {
		log.Fatal("DB connection error")
		return false
	}

	result, err := db.Query(query)
	defer result.Close()

	if err != nil {
		log.Fatal("Query error: " + query)
		return false
	} else {
		return true
	}
}

func GetUser(name, pass string) structs.User {
	db, err := sql.Open("mysql", db_user + ":" + db_password + "@/" + db_name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT * FROM users WHERE username = '" + name + "' AND password = '" + GetEncodeString(pass) + "'"
	result := db.QueryRow(query)

	User := structs.User{}
	err = result.Scan(&User.Id, &User.Username, &User.Password, &User.DateCreated)

	if err != nil {
		log.Fatal(err)
	}

	return User
}

func UserIsAuthenticated(uname, pswd string) bool {
	var isAuthenticated bool
	db, err := sql.Open("mysql", db_user + ":" + db_password + "@/" + db_name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT IF(COUNT(*),'true','false') FROM users WHERE username = '"+uname+"' AND password = '"+GetEncodeString(pswd)+"'"

	err = db.QueryRow(query).Scan(&isAuthenticated)
	if err != nil {
		log.Fatal(err)
	}

	return isAuthenticated
}