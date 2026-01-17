package main

import (
	"fmt"
	"net/http"

	"github.com/Hamiduzzaman96/Test-Project/handlers"
)

func main() {
	mux := http.NewServeMux()                                                          //Router
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProducts))            // Route
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))                //Route
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductbyID)) //Route
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
