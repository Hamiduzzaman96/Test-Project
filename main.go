package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Plz provide get request", 400)
	}
	json.NewEncoder(w).Encode(productList)
}

func createProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Please provide valid post request", 400)
	}
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

type Products struct {
	ID          int
	Title       string
	Description string
	price       float64
	ImageURL    string
}

var newProduct Products // create a instance or object for products struct

var productList []Products //empty slice

func main() {
	mux := http.NewServeMux()                          //Router
	mux.HandleFunc("/create-products", createProducts) // Route
	mux.HandleFunc("/products", getProducts)           //Route
	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", mux) //Start the server
	if err != nil {
		fmt.Println("Error starting the sever", err)
	}

}

func init() {
	prd1 := Products{
		ID:          1,
		Title:       "Orange",
		Description: "This is orange and i love it",
		price:       20.5,
		ImageURL:    "https://www.istockphoto.com/photos/orange",
	}
	prd2 := Products{
		ID:          2,
		Title:       "Apple",
		Description: "This is apple and i love it",
		price:       100,
		ImageURL:    "https://www.istockphoto.com/photos/green-apple",
	}
	prd3 := Products{
		ID:          3,
		Title:       "Grape",
		Description: "This is Grape and i love it",
		price:       60,
		ImageURL:    "https://www.istockphoto.com/photos/red-grape",
	}
	prd4 := Products{
		ID:          4,
		Title:       "Banana",
		Description: "This is Banana and i love it",
		price:       70,
		ImageURL:    "https://stock.adobe.com/search?k=picture+of+banana",
	}

	productList = append(productList, prd1, prd2, prd3, prd4) //empty slice a append kora hoyece
}
