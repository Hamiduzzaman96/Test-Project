package handlers

import (
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

func GetProductbyID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product ID", 404)
		return
	}
	product := database.Get(pId)
	util.SendData(w, product, 200)
}
