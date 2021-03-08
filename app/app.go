package app

import (
	"fmt"
	"log"

	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/jonathanwamsley/banking/config"
	"github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/logger"
	"github.com/jonathanwamsley/banking/service"
)

// getDbClient loads and returns db connection. The db makes connection is confirmed via Ping
func getDbClient(connectionInfo string) *sqlx.DB {

	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", connectionInfo)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(); err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

// Start helps decouples from running the whole entire application
// it connects the handlers, starts the server, and any other configuration setup
func Start() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("no .env file found")
		panic(err)
	}
	config := config.NewConfig()
	dbClient := getDbClient(config.GetMySQLInfo())
	serverInfo := config.GetServerInfo()

	router := mux.NewRouter()
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient))}
	ah := AccountHandler{service.NewAccountService(domain.NewAccountRepositoryDB(dbClient))}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer", ch.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.DeleteCustomer).Methods(http.MethodDelete)

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.GetAccount).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.DeleteAccount).Methods(http.MethodDelete)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	logger.Info(fmt.Sprintf("Starting server on %s ...", serverInfo))
	log.Fatal(http.ListenAndServe(serverInfo, router))
}
