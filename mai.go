package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Plz provide get request", 400)
	}
	json.NewEncoder(w).Encode(productList)
}

type Products struct {
	ID          int
	Title       string
	Description string
	price       float64
	ImageURL    string
}

var productList []Products

func main() {
	mux := http.NewServeMux() //Router

	mux.HandleFunc("/products", getProducts)
	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", mux)
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

	productList = append(productList, prd1, prd2, prd3, prd4)
}
