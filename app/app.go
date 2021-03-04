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
	"github.com/jonathanwamsley/banking/service"
)

var (
	// Client holds the db connection.
	Client *sqlx.DB
)

// loads the config variables and makes sure a database connection can be established
func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file found")
	}
	conf := config.NewConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		conf.MySQL.Username,
		conf.MySQL.Password,
		conf.MySQL.Host,
		conf.MySQL.Schema,
	)

	var err error
	Client, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	Client.SetConnMaxLifetime(time.Minute * 3)
	Client.SetMaxOpenConns(10)
	Client.SetMaxIdleConns(10)
	log.Println("database successfully configured")
}

// Start helps decouples from running the whole entire application
// it connects the handlers, starts the server, and any other configuration setup
func Start() {
	router := mux.NewRouter()

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
