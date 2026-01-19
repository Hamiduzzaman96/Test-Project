package handlers

import (
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Test-Project/database"
	"github.com/Hamiduzzaman96/Test-Project/util"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product ID", 404)
		return
	}

	database.Delete(pId)

	util.SendData(w, "Successfully deleted product", 201)
}
