package utils

import (
	"encoding/json"
	"../structs"
)

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