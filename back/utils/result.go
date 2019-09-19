package utils

import (
	"../structs"
	"encoding/json"
)

func SendStatus(status bool) string {
	if status == true {
		return GetResultTemplate("true")
	} else {
		return GetResultTemplate("false")
	}
}

func GetResultTemplate(str string) string {
	return "{\"result\":" + str + "}"
}

func GetResultTest(obj []structs.TestStruct) string {
	res, err := json.Marshal(obj)

	if err != nil {
		panic("Json error")
	}

	return string(res)
}

func GetResultSession(obj structs.Session) string {
	res, err := json.Marshal(obj)

	if err != nil {
		panic("Json error")
	}

	return GetResultTemplate(string(res))
}

func GetResultArtists(obj []*structs.Artist) string {
	res, err := json.Marshal(obj)

	if err != nil {
		panic("Json error")
	}

	return GetResultTemplate(string(res))
}