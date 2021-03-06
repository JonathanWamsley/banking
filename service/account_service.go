package service

import (
	"fmt"

	"github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

// AccountService is an interface that implements
//
// CreateAccount: creates a new account for a given customer and returns account id back on success
// GetAccount: gets the user checking and savings account
// DeleteAccount: deletes a user account
type AccountService interface {
	CreateAccount(dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errs.AppError)
	GetAccount(id string) ([]dto.GetAccountResponse, *errs.AppError)
	DeleteAccount(id string, accountType string) *errs.AppError
}

// DefaultAccountService has methods that call dto and the domain
type DefaultAccountService struct {
	repo domain.AccountRepository
}

// NewAccountService  is the entry point to the service to create a DefaultAccountService struct
func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}

// CreateAccount manages the account dto and database interaction
func (s DefaultAccountService) CreateAccount(req dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errs.AppError) {

	accounts, _ := s.repo.ByID(req.CustomerID)
	for _, a := range accounts {
		if a.AccountType == req.AccountType {
			return nil, errs.NewValidationError(fmt.Sprintf("Error, only one %s type is allow", req.AccountType))
		}
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	account := domain.NewAccount(req)
	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}
	return newAccount.ToCreateAccountResponseDTO(), nil
}

// GetAccount manages the account dto and database interaction
func (s DefaultAccountService) GetAccount(id string) ([]dto.GetAccountResponse, *errs.AppError) {
	accounts, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}

	response := make([]dto.GetAccountResponse, 0)
	for _, a := range accounts {
		response = append(response, a.ToGetAccountResponseDTO())
	}
	return response, nil
}

// DeleteAccount deletes an account using a customer_id and account_type
func (s DefaultAccountService) DeleteAccount(id string, accountType string) *errs.AppError {
	accounts, err := s.repo.ByID(id)
	if err != nil {
		return err
	}
	for _, a := range accounts {
		if a.AccountType == accountType && a.Amount != 0 {
			return errs.NewValidationError("Account must have a balance of 0 to close")
		}
	}

	err = s.repo.Delete(id, accountType)
	if err != nil {
		return err
	}
	return nil
}
