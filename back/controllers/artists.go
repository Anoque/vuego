package controllers

import (
	"../utils"
	"log"
)

func GetArtists(limit, offset rune) string {
	log.Print("test")
	artists := utils.GetArtists(100, 0)

	return utils.GetResultArtists(artists)
}