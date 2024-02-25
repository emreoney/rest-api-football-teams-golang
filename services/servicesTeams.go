package services

import (
	"gomod/database"
	"gomod/models"
)

func CreateTeam(newTeam models.Team) models.Team {
	database.DB.Create(&newTeam)
	return *&newTeam
}

func GetTeams(teams []models.Team) []models.Team {
	database.DB.Preload("Players").Find(&teams)
	return *&teams
}

func GetTeam(team models.Team) models.Team {
	database.DB.Preload("Players").Find(&team)
	return *&team
}

func UpdateTeam(updatedTeam models.Team) models.Team {
	database.DB.Save(&updatedTeam)
	return *&updatedTeam
}

func DeleteTeam(deletedTeam models.Team) {
	database.DB.Delete(&deletedTeam)
}
