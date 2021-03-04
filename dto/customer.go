package dto

import (
	"strings"

	"github.com/jonathanwamsley/banking/errs"
)

// CustomerResponse returns the expected json response after a customer query is requested
type CustomerResponse struct {
	ID          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

// CustomerRequest holds the expected input fields from a request
type CustomerRequest struct {
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
}

// Validate makes sure all fields are not empty
func (c CustomerRequest) Validate() *errs.AppError {
	c.Name = strings.TrimSpace(c.Name)
	c.City = strings.TrimSpace(c.City)
	c.Zipcode = strings.TrimSpace(c.Zipcode)
	c.DateofBirth = strings.TrimSpace(c.DateofBirth)

	if c.Name == "" {
		return errs.NewValidationError("invalid name")
	}
	if c.City == "" {
		return errs.NewValidationError("invalid city")
	}
	if c.Zipcode == "" {
		return errs.NewValidationError("invalid zipcode")
	}
	if c.DateofBirth == "" {
		return errs.NewValidationError("invalid date of birth")
	}
	return nil
}
