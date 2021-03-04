package domain

// CustomerRepositoryStub is used to test basic api functionality before implementing a db
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll returns all customers
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub creates the mock data
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
