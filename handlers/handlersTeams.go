package handlers

import (
	"encoding/json"
	"fmt"
	"gomod/database"
	"gomod/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerCreateTeam(w http.ResponseWriter, r *http.Request) {
	var newTeam models.Team
	json.NewDecoder(r.Body).Decode(&newTeam)

	database.DB.Create(&newTeam)

	data, _ := json.Marshal(newTeam)
	fmt.Fprintf(w, string(data))
}
func HandlerGetTeams(w http.ResponseWriter, r *http.Request) {
	var teams []models.Team
	database.DB.Preload("Players").Find(&teams)

	data, _ := json.Marshal(teams)
	fmt.Fprintf(w, string(data))
}
func HandlerGetTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	variables := mux.Vars(r)
	teamID, _ := strconv.Atoi(variables["id"])

	team.ID = uint(teamID)
	database.DB.Preload("Players").Find(&team)

	data, _ := json.Marshal(team)
	fmt.Fprintf(w, string(data))
}
func HandlerUpdateTeam(w http.ResponseWriter, r *http.Request) {
	var updatedTeam models.Team
	variables := mux.Vars(r)
	teamID, _ := strconv.Atoi(variables["id"])
	updatedTeam.ID = uint(teamID)

	json.NewDecoder(r.Body).Decode(&updatedTeam)
	database.DB.Save(&updatedTeam)

	data, _ := json.Marshal(updatedTeam)
	fmt.Fprintf(w, string(data))
}
func HandlerDeleteTeam(w http.ResponseWriter, r *http.Request) {
	var deletedTeam models.Team
	variables := mux.Vars(r)
	teamID, _ := strconv.Atoi(variables["id"])
	deletedTeam.ID = uint(teamID)

	database.DB.Delete(&deletedTeam)

	responseMessage := models.Information{"Data has deleted"}
	data, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	fmt.Fprintf(w, string(data))
}
