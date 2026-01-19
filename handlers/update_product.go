package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Products
	productID := r.PathValue("productId")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product ID", 404)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&newProduct) //decode from json
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz provide valid json", 400)
		return
	}

	newProduct.ID = pId

	database.Update(newProduct)

	util.SendData(w, "Successfully updated product", 201)
}
