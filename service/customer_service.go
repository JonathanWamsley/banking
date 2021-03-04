package service

import (
	"github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

// CustomerService is an interface that implements
//
// GetAllCustomer: returns all the customers or an error
// CreateCustomer: inserts a new customer into the db
// GetCustomer: returns a customer by id
// DeleteCustomer: Removes a customer by id from the db
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	CreateCustomer(dto.CustomerRequest) (*dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
	DeleteCustomer(string) *errs.AppError
}

// DefaultCustomerService has methods that call upon dto and domain
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// NewCustomerService is the entry point to the service to create a DefaultCustomerService struct
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

// GetAllCustomers returns all the customers as dto response
//
// if unsuccessfull, an AppError is sent with the error code and message
func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDTO())
	}
	return response, nil
}

// CreateCustomer validates customer, creates a customer and returns the customer information back with an customer id
func (s DefaultCustomerService) CreateCustomer(c dto.CustomerRequest) (*dto.CustomerResponse, *errs.AppError) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	customer := domain.NewCustomer(c)
	newCustomer, err := s.repo.Save(customer)
	if err != nil {
		return nil, err
	}

	response := newCustomer.ToDTO()

	return &response, nil
}

// GetCustomer returns a single customer by id
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDTO()
	return &response, nil
}

// DeleteCustomer first makes sure an id exists, then it removes the customer by id
func (s DefaultCustomerService) DeleteCustomer(id string) *errs.AppError {
	_, err := s.repo.ByID(id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
