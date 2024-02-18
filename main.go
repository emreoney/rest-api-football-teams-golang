package main

import (
	"gomod/database"
	"gomod/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.Init()

	router := mux.NewRouter()

	//League APIs
	router.HandleFunc("/league", handlers.HandlerCreateLeague).Methods("POST")
	router.HandleFunc("/league", handlers.HandlerGetLeagues).Methods("GET")
	router.HandleFunc("/league/{id}", handlers.HanglerGetLeauge).Methods("GET")
	router.HandleFunc("/league/{id}", handlers.HandlerUpdateLeague).Methods("PUT")
	router.HandleFunc("/league/{id}", handlers.HandlerDeleteLeague).Methods("DELETE")

	//Teams APIs
	router.HandleFunc("/team", handlers.HandlerCreateTeam).Methods("POST")
	router.HandleFunc("/team", handlers.HandlerGetTeams).Methods("GET")
	router.HandleFunc("/team/{id}", handlers.HandlerGetTeam).Methods("GET")
	router.HandleFunc("/team/{id}", handlers.HandlerUpdateTeam).Methods("PUT")
	router.HandleFunc("/team/{id}", handlers.HandlerDeleteTeam).Methods("DELETE")

	//Player APIs
	router.HandleFunc("/player", handlers.HandlerCreatePlayer).Methods("POST")
	router.HandleFunc("/player", handlers.HandlerGetPlayers).Methods("GET")
	router.HandleFunc("/player/{id}", handlers.HandlerGetPlayer).Methods("GET")
	router.HandleFunc("/player/{id}", handlers.HandlerUpdatePlayer).Methods("PUT")
	router.HandleFunc("/player/{id}", handlers.HandlerDeletePlayer).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
