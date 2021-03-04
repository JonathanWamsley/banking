package domain

// Customer hold locality information are are owners of accounts
type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// CustomerRepository implements:
//
// FindAll: returns all the customers or an error
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}