package domain

import (
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

// a time layout that the db is using
const dbTSLayout = "2006-01-02 15:04:05"

// Account holds banking information for an account
type Account struct {
	AccountID   string `db:"account_id"`
	CustomerID  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

// AccountRepository implements:
//
// Save: creates a new account for a customer, and returns customer account id
// ById: searches for accounts by a user_id
// Delete: deletes an account using a customer id and account type
// SaveTransaction: makes a transaction in a bank account and returns new account total
// FindBy: finds a specific account information
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	ByID(customerID string) ([]Account, *errs.AppError)
	Delete(id string, accountType string) *errs.AppError
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountID string) (*Account, *errs.AppError)
}

// ToCreateAccountResponseDTO converts account from database to account response for user
func (a Account) ToCreateAccountResponseDTO() *dto.CreateAccountResponse {
	return &dto.CreateAccountResponse{AccountID: a.AccountID}
}

// ToGetAccountResponseDTO converts a account to a account response for the user
func (a Account) ToGetAccountResponseDTO() dto.GetAccountResponse {
	return dto.GetAccountResponse{
		AccountID:   a.AccountID,
		CustomerID:  a.CustomerID,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
	}
}

// NewAccount converts account request from user to an account to be processed by db
func NewAccount(a dto.CreateAccountRequest) Account {
	return Account{
		CustomerID:  a.CustomerID,
		OpeningDate: dbTSLayout,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      "1",
	}
}

// CanWithdraw checks if an transaction can be made
func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
