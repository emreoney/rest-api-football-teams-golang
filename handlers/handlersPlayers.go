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

func HandlerCreatePlayer(w http.ResponseWriter, r *http.Request) {
	var newPlayer models.Player
	json.NewDecoder(r.Body).Decode(&newPlayer)

	playerInformations := services.CreatePlayer(newPlayer)

	data, err := json.Marshal(playerInformations)
	if err != nil {
		fmt.Println("HATA: ", err.Error())
	}
	fmt.Fprintf(w, string(data))
}
func HandlerGetPlayers(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	//	database.DB.Find(&players)

	playersInformaitons := services.GetPlayers(players)

	data, _ := json.Marshal(playersInformaitons)
	fmt.Fprintf(w, string(data))
}
func HandlerGetPlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	variables := mux.Vars(r)
	playerID, _ := strconv.Atoi(variables["id"])

	player.ID = uint(playerID)
	//database.DB.First(&player)
	playerInformations := services.GetPlayer(player)

	data, _ := json.Marshal(playerInformations)
	fmt.Fprintf(w, string(data))
}
func HandlerUpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var updatedPlayer models.Player
	variables := mux.Vars(r)
	playerID, _ := strconv.Atoi(variables["id"])
	updatedPlayer.ID = uint(playerID)
	json.NewDecoder(r.Body).Decode(&updatedPlayer)

	//database.DB.Save(&updatedPlayer)
	playerInformations := services.UpdatePlayer(updatedPlayer)
	data, _ := json.Marshal(playerInformations)
	fmt.Fprintf(w, string(data))
}
func HandlerDeletePlayer(w http.ResponseWriter, r *http.Request) {
	var deletedPlayer models.Player
	variables := mux.Vars(r)
	playerID, _ := strconv.Atoi(variables["id"])
	deletedPlayer.ID = uint(playerID)

	// database.DB.Delete(&deletedPlayer)
	services.DeletePlayer(deletedPlayer)
	responseMessage := models.Information{"Data has deleted"}
	data, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	fmt.Fprintf(w, string(data))
}
