package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.ProductList)
}
func GetProductbyID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")
	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product ID", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pID {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "Data not found", 404)
}
