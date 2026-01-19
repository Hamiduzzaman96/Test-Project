package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hamiduzzaman96/Test-Project/config"
	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	err := json.NewDecoder(r.Body).Decode(&req) //decode from json
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz provide valid json", http.StatusBadRequest)
		return
	}
	usr := database.Find(req.Email, req.Password)
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, accessToken, http.StatusCreated)

}
