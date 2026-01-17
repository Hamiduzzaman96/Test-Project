package main

import (
	"encoding/json"
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(productList)
}
