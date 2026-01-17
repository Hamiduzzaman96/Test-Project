package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createProducts(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&newProduct) //decode from json
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz provide valid json", 400)
	}

	newProduct.ID = len(productList) + 1          // automatic generated id increasing by 1
	productList = append(productList, newProduct) // append new product in product list
	w.WriteHeader(201)                            // http server status for created

	json.NewEncoder(w).Encode(newProduct) // send client or frontend
}
