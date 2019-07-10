package utils

import (
	"encoding/json"
	"../structs"
)

func GetResultTest(obj []structs.TestStruct) string {
	res, err := json.Marshal(obj)

	if err != nil {
		panic("Json error")
	}

	return string(res)
}