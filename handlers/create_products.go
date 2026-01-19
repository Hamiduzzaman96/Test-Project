package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Products
	err := json.NewDecoder(r.Body).Decode(&newProduct) //decode from json
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz provide valid json", 400)
		return
	}
	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, 201)

}
