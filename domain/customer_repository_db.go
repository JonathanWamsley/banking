package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/jonathanwamsley/banking/errs"
)

// the query need
const (
	findAllSQL = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
)

// CustomerRepositoryDB holds the sql client connection
type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// NewCustomerRepositoryDB creates a new CustomerRepositoryDB to call sql methods
func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client}
}

// FindAll returns all the customers from the database
func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	err := d.client.Select(&customers, findAllSQL)

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}
