package domain

import (
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
)

// a time layout that the db is using
const dbTSLayout = "2006-01-02 15:04:05"

// Account holds banking information for an account
type Account struct {
	AccountID string `db:"account_id"`
	CustomerID string `db:"customer_id"`
	OpeningDate	string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount float64
	Status string
}

// AccountRepository implements:
//
// Save: creates a new account for a customer, and returns customer account id
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

// ToCreateAccountResponseDTO converts account from database to account response for user
func (a Account) ToCreateAccountResponseDTO() *dto.CreateAccountResponse {
	return &dto.CreateAccountResponse{AccountID: a.AccountID}
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