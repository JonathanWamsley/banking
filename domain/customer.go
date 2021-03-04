package domain

import (
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

// Customer hold locality information are are owners of accounts
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

// statusAsText converts numeral string 0/1 to inactive/active
func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

// ToDTO converts a customer object to the appropriate response to be passed from the service to the handler to the caller
func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

// CustomerRepository implements:
//
// FindAll: returns all the customers or an error
type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
}
