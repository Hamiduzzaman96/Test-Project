package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	err := json.NewDecoder(r.Body).Decode(&newUser) //decode from json
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz provide valid json", http.StatusBadRequest)
		return
	}
	createdUser := newUser.Store()
	util.SendData(w, createdUser, http.StatusCreated)

}
