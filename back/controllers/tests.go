package controllers

import (
	"../structs"
	"../utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetTests() string {
	db, err := sql.Open("mysql", "hello:root@/checkout")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM test LIMIT 10")

	if err != nil{
		panic(err)
	}

	var texts []structs.TestStruct

	for result.Next() {
		t := structs.TestStruct{}
		err := result.Scan(&t.Id, &t.Text)

		if err != nil {
			fmt.Println("Error")
			continue
		}

		texts = append(texts, t)
	}

	return utils.GetResultTest(texts)
}