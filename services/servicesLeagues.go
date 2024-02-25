package services

import (
	"gomod/database"
	"gomod/models"
)

func CreateLeague(newLeague models.League) models.League {
	database.DB.Create(&newLeague)
	return *&newLeague
}
func GetLeagues(leagues []models.League) []models.League {
	database.DB.Preload("Teams").Find(&leagues)
	return *&leagues
}
func GetLeague(league models.League) models.League {
	database.DB.Preload("Teams").First(&league)
	return *&league
}
func UpdateLeague(updatedLeague models.League) models.League {
	database.DB.Save(&updatedLeague)
	return *&updatedLeague
}
func DeleteLeague(deletedLeague models.League) {
	database.DB.Delete(&deletedLeague)
}
