package main

import (
	"fmt"
	"net/http"
)

var newProduct Products // create a instance or object for products struct

var productList []Products //empty slice

func main() {
	mux := http.NewServeMux()                                             //Router
	mux.Handle("POST /create-products", http.HandlerFunc(createProducts)) // Route
	mux.Handle("GET /products", http.HandlerFunc(getProducts))            //Route
	globalRouter(mux)
	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", mux) //Start the server
	if err != nil {
		fmt.Println("Error starting the sever", err)
	}

}

func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)

}
