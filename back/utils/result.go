package utils

import (
	"encoding/json"
	"../structs"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db_user = "hello"
var db_password = "root"
var db_name = "vuego"

func SendStatus(status bool) string {
	if status == true {
		return "{\"result\":true}"
	} else {
		return "{\"result\":false}"
	}
}

func GetResultTest(obj []structs.TestStruct) string {
	res, err := json.Marshal(obj)

	if err != nil {
		panic("Json error")
	}

	return string(res)
}

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