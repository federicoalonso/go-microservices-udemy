package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomers).Methods(http.MethodGet)

	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
