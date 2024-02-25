package handlers

import (
	"encoding/json"
	"fmt"
	"gomod/models"
	"gomod/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerCreateTeam(w http.ResponseWriter, r *http.Request) {
	var newTeam models.Team
	json.NewDecoder(r.Body).Decode(&newTeam)

	teamInformation := services.CreateTeam(newTeam)

	data, _ := json.Marshal(teamInformation)
	fmt.Fprintf(w, string(data))
}
func HandlerGetTeams(w http.ResponseWriter, r *http.Request) {
	var teams []models.Team

	teamsInformations := services.GetTeams(teams)

	data, _ := json.Marshal(teamsInformations)
	fmt.Fprintf(w, string(data))
}
func HandlerGetTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	variables := mux.Vars(r)
	teamID, _ := strconv.Atoi(variables["id"])

	team.ID = uint(teamID)

	teamInformation := services.GetTeam(team)

	data, _ := json.Marshal(teamInformation)
	fmt.Fprintf(w, string(data))
}
func HandlerUpdateTeam(w http.ResponseWriter, r *http.Request) {
	var updatedTeam models.Team
	variables := mux.Vars(r)
	teamID, _ := strconv.Atoi(variables["id"])
	updatedTeam.ID = uint(teamID)

	json.NewDecoder(r.Body).Decode(&updatedTeam)
	teamInformation := services.UpdateTeam(updatedTeam)

	data, _ := json.Marshal(teamInformation)
	fmt.Fprintf(w, string(data))
}
func HandlerDeleteTeam(w http.ResponseWriter, r *http.Request) {
	var deletedTeam models.Team
	variables := mux.Vars(r)
	teamID, _ := strconv.Atoi(variables["id"])
	deletedTeam.ID = uint(teamID)

	services.DeleteTeam(deletedTeam)

	responseMessage := models.Information{"Data has deleted"}
	data, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	fmt.Fprintf(w, string(data))
}
