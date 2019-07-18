package controllers

import (
	"../structs"
	"../utils"
	"log"
	"strconv"
	"time"
)

func CreateSession(UserId int) structs.Session {
	SessionHash := utils.RandStringBytes(32)
	query := "INSERT INTO sessions (session, user_id) VALUES ('"+SessionHash+"', "+strconv.Itoa(UserId)+")"
	log.Println(query)
	status := utils.RunQuery(query)
	Session := structs.Session{}

	if status {
		currentTime := time.Now()
		Session = structs.Session{Id: 0, Session: SessionHash, UserId: UserId, DateCreated: currentTime.Format("2006-01-02 15:04:05")}
	} else {
		log.Println("Error in create of session")
	}

	return Session
}