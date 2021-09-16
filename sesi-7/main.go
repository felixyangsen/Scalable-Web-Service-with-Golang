package main

import (
	"fmt"
	"net/http"

	"myapp/service"

	"github.com/gorilla/mux"
)

var (
	port = ":8080"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", service.GetOrdersHandler).Methods("GET")
	router.HandleFunc("/create-order", service.CreateOrderHandler).Methods("POST")
	// router.HandleFunc("/update-order", ).Methods("PUT")
	// router.HandleFunc("/delete-order", ).Methods("DEL")

	fmt.Println("running at http://localhost" + port)
	http.ListenAndServe(port, router)
}
