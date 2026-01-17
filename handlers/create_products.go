package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&database.NewProduct) //decode from json
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz provide valid json", 400)
	}

	database.NewProduct.ID = len(database.ProductList) + 1                   // automatic generated id increasing by 1
	database.ProductList = append(database.ProductList, database.NewProduct) // append new product in product list
	util.SendData(w, database.NewProduct, 201)

}
