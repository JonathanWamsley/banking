package service

import (
	"github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

// AccountService is an interface that implements
//
// CreateAccount: creates a new account for a given customer and returns account id back on success
type AccountService interface {
	CreateAccount(dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errs.AppError)
}

// DefaultAccountService has methods that call dto and the domain
type DefaultAccountService struct{
	repo domain.AccountRepository
}

// NewAccountService  is the entry point to the service to create a DefaultAccountService struct
func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}

// CreateAccount manages the account dto and database interaction
func (s DefaultAccountService) CreateAccount(req dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errs.AppError) {
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