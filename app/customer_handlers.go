package app

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanwamsley/banking/service"
)

// CustomerHandler connects routing parateters and the CustomerService being called
type CustomerHandler struct {
	service service.CustomerService
}

// GetAllCustomers returns all customers
//
func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		writeResponse(w, http.StatusNotFound, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

// writeResponse returns the header with an encoded json data as a response
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
