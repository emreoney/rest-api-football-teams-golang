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

func HandlerCreateLeague(w http.ResponseWriter, r *http.Request) {
	var newLeague models.League
	json.NewDecoder(r.Body).Decode(&newLeague)

	leagueInformation := services.CreateLeague(newLeague)

	data, err := json.Marshal(leagueInformation)
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	fmt.Fprintf(w, string(data))
}

func HandlerGetLeagues(w http.ResponseWriter, r *http.Request) {
	var leagues []models.League

	leaguesInformations := services.GetLeagues(leagues)

	data, err := json.Marshal(leaguesInformations)
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	fmt.Fprintf(w, string(data))
}
func HanglerGetLeauge(w http.ResponseWriter, r *http.Request) {
	var league models.League
	variables := mux.Vars(r)
	leagueID, err := strconv.Atoi(variables["id"])
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	league.ID = uint(leagueID)
	services.GetLeague(league)

	data, _ := json.Marshal(league)
	fmt.Fprintf(w, string(data))
}
func HandlerUpdateLeague(w http.ResponseWriter, r *http.Request) {
	var updatedLeague models.League
	variables := mux.Vars(r)
	leagueID, err := strconv.Atoi(variables["id"])
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	updatedLeague.ID = uint(leagueID)
	json.NewDecoder(r.Body).Decode(&updatedLeague)

	services.UpdateLeague(updatedLeague)

	data, err := json.Marshal(updatedLeague)
	fmt.Fprintf(w, string(data))
}
func HandlerDeleteLeague(w http.ResponseWriter, r *http.Request) {
	var deletedLeague models.League
	variables := mux.Vars(r)
	leagueID, _ := strconv.Atoi(variables["id"])
	deletedLeague.ID = uint(leagueID)

	services.DeleteLeague(deletedLeague)

	responseMessage := models.Information{"Data has deleted"}
	data, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	fmt.Fprintf(w, string(data))

}
