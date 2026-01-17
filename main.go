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
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz provide valid json", 400)
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	w.WriteHeader(201)

	json.NewEncoder(w).Encode(newProduct)
}

type Products struct {
	ID          int
	Title       string
	Description string
	price       float64
	ImageURL    string
}

var newProduct Products

var productList []Products //empty slice

func main() {
	mux := http.NewServeMux() //Router

	mux.HandleFunc("/products", getProducts) //Route
	fmt.Println("Server running on :8080")

	mux.HandleFunc("/create-products", createProducts)

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
