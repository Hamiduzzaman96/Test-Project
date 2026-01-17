package main

import "net/http"

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Plz provide get request", 400)
	}
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
}

func init() {

}
