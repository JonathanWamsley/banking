package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/service"
)

// AccountHandler connects account routing options to account services
type AccountHandler struct {
	service service.AccountService
}

// CreateAccount returns a new account for a customer on success
func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["customer_id"]

	var accountRequest dto.CreateAccountRequest
	accountRequest.CustomerID = id
	if err := json.NewDecoder(r.Body).Decode(&accountRequest); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	result, err := ah.service.CreateAccount(accountRequest)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	}
	writeResponse(w, http.StatusCreated, result)
}

// GetAccount returns account information for a customer
func (ah *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["customer_id"]

	result, err := ah.service.GetAccount(id)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	}
	writeResponse(w, http.StatusOK, result)
}
