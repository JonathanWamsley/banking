package service

import (
	"github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

// CustomerService is an interface that implements
//
// GetAllCustomer: returns all the customers or an error
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	CreateCustomer(dto.CustomerRequest) (*dto.CustomerResponse, *errs.AppError)
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
