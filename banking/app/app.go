package app

import (
	"log"
	"net/http"

	"github.com/federicoalonso/go-microservices-udemy/banking/domain"
	"github.com/federicoalonso/go-microservices-udemy/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	// wiring
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
