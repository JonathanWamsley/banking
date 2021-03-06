package service

import (
	"fmt"
	"time"

	"github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

// AccountService is an interface that implements
//
// CreateAccount: creates a new account for a given customer and returns account id back on success
// GetAccount: gets the user checking and savings account
// DeleteAccount: deletes a user account
// MakeTransaction: a customer creates a transation into an account and receive the new balance
type AccountService interface {
	CreateAccount(dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errs.AppError)
	GetAccount(id string) ([]dto.GetAccountResponse, *errs.AppError)
	DeleteAccount(id string, accountType string) *errs.AppError
	MakeTransaction(request dto.MakeTransactionRequest) (*dto.MakeTransactionResponse, *errs.AppError)
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

// MakeTransaction makes a withdrawal or deposit to an account. It then returns the updated balance for the account
func (s DefaultAccountService) MakeTransaction(req dto.MakeTransactionRequest) (*dto.MakeTransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountID)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDTO()
	return &response, nil
}
