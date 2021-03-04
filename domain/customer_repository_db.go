package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/jonathanwamsley/banking/errs"
)

// the query need
const (
	findAllCustomers = "select customer_id, name, city, zipcode, date_of_birth, status from customers;"
	insertCustomer   = "insert into customers(name, date_of_birth, city, zipcode, status) values(?, ?, ?, ?, ?);"
	getCustomer      = "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?;"
	deleteCustomer   = "delete from customers where customer_id=?;"
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
	err := d.client.Select(&customers, findAllCustomers)

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

// Save inserts a new customer and returns back the customer information with an id
func (d CustomerRepositoryDB) Save(c Customer) (*Customer, *errs.AppError) {
	result, err := d.client.Exec(insertCustomer, c.Name, c.DateofBirth, c.City, c.Zipcode, c.Status)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	c.ID = strconv.FormatInt(id, 10)

	return &c, nil
}

// ByID returns the customer by an id
func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	var c Customer
	err := d.client.Get(&c, getCustomer, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

// Delete requires a user id. No error is returned on success
func (d CustomerRepositoryDB) Delete(id string) *errs.AppError {
	_, err := d.client.Exec(deleteCustomer, id)
	if err != nil {
		return errs.NewUnexpectedError("Unexpected error from database")
	}
	return nil
}
