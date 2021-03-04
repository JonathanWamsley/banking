package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/service"
)

// Start helps decouples from running the whole entire application
// it connects the handlers, starts the server, and any other configuration setup
func Start() {
	router := mux.NewRouter()
	
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}