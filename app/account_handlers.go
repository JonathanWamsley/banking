package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
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
		return
	}

	result, err := ah.service.CreateAccount(accountRequest)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
		return
	}
	writeResponse(w, http.StatusCreated, result)
}

// GetAccount returns account information for a customer
func (ah *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["customer_id"]

	result, err := ah.service.GetAccount(id)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
		return
	}
	writeResponse(w, http.StatusOK, result)
}

// DeleteAccount uses a account type query and a customer_id to delete an account
func (ah *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["customer_id"]
	accountType := r.URL.Query().Get("account_type")

	if badAccountType(accountType) {
		err := errs.NewNotFoundError("invalid query parameter for account_type")
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	err := ah.service.DeleteAccount(id, accountType)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
		return
	}
	writeResponse(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func badAccountType(accountType string) bool {
	return accountType != "checking" && accountType != "saving"
}
