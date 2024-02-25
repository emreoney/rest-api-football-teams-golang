package services

import (
	"gomod/database"
	"gomod/models"
)

func CreatePlayer(newPlayer models.Player) models.Player {
	database.DB.Create(&newPlayer)
	return *&newPlayer
}

func GetPlayers(players []models.Player) []models.Player {
	database.DB.Find(&players)
	return *&players
}

func GetPlayer(player models.Player) models.Player {
	database.DB.First(&player)
	return *&player
}

func DeletePlayer(deletedPlayer models.Player) {
	database.DB.Delete(&deletedPlayer)
}

func UpdatePlayer(updatedPlayer models.Player) models.Player {
	database.DB.Save(&updatedPlayer)
	return *&updatedPlayer
}
