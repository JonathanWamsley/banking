package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/service"
)

// CustomerHandler connects routing parateters and the CustomerService being called
type CustomerHandler struct {
	service service.CustomerService
}

// GetAllCustomers returns all customers
func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customers)
}

// CreateCustomer returns the customers information back with an ID
func (ch *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerRequest dto.CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		writeResponse(w, http.StatusBadRequest, "invalid json")
		return
	}
	customer, err := ch.service.CreateCustomer(customerRequest)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customer)
}

// GetCustomer gets a customer by id
func (ch *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customer)
}

// DeleteCustomer returns a confirmation status deleted if success
func (ch *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["customer_id"]

	err := ch.service.DeleteCustomer(customerID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// writeResponse returns the header with an encoded json data as a response
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
