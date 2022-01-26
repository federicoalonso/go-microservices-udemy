package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/federicoalonso/go-microservices-udemy/banking/service"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "John", City: "New York", Zipcode: "10001"},
	// 	{Name: "Mary", City: "New York", Zipcode: "10001"},
	// 	{Name: "Bob", City: "New York", Zipcode: "10001"},
	// }

	customers, _ := ch.service.GetAll()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
