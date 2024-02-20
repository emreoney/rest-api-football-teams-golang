package handlers

import (
	"encoding/json"
	"fmt"
	"gomod/database"
	"gomod/helpers"
	"gomod/models"
	"net/http"
)

func HandlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount models.Account
	json.NewDecoder(r.Body).Decode(&newAccount)

	database.DB.Create(&newAccount)

	data, err := json.Marshal(newAccount)
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
	fmt.Fprintf(w, string(data))
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)

	var accountDb models.Account
	database.DB.Where("username = ?", account.Username).Find(&accountDb)

	if account.Username == accountDb.Username && account.Password == accountDb.Password {
		token, _ := helpers.CreateToken()
		account.Token = token
		data, _ := json.Marshal(account)
		fmt.Fprintf(w, string(data))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		messsage := models.Information{"Wrong username or password!"}
		data, _ := json.Marshal(messsage)
		fmt.Fprintf(w, string(data))
	}

}
