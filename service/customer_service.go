package service

import "github.com/jonathanwamsley/banking/domain"

// CustomerService is an interface that implements
//
// GetAllCustomer: returns all the customers or an error
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// DefaultCustomerService has methods that call upon dto and domain
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// NewCustomerService is the entry point to the service to create a DefaultCustomerService struct
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

// GetAllCustomers returns all the customers or an error
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}
