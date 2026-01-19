package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Hamiduzzaman96/Test-Project/config"
	"github.com/Hamiduzzaman96/Test-Project/handlers"
)

func GetRoutes(cnf config.Config) {
	mux := http.NewServeMux()                                                            //Router
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProducts))              // Route
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))                  //Route
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductbyID))   //Route
	mux.Handle("PUT /products/{productId}", http.HandlerFunc(handlers.UpdateProduct))    //Route
	mux.Handle("DELETE /products/{productId}", http.HandlerFunc(handlers.DeleteProduct)) //Route
	mux.Handle("POST /users", http.HandlerFunc(handlers.CreateUser))                     //Route
	mux.Handle("POST /users/login", http.HandlerFunc(handlers.Login))                    //Route
	globalRouter(mux)

	port := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port:", port)

	err := http.ListenAndServe(port, mux) //Start the server
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
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
