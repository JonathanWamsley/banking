package dto

import (
	"strings"

	"github.com/jonathanwamsley/banking/errs"
)

const (
	MINIMUM_FUNDS = 5000.0
	SAVING = "saving"
	CHECKING = "checking"
)

// CreateAccountRequest must follow this format to create a new account
type CreateAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

// CreateAccountResponse must follow this format to return a new account created
type CreateAccountResponse struct {
	AccountID string `json:"account_id"`
}

// GetAccountResponse must follow this format to return a an account
type GetAccountResponse struct {
	AccountID   string  `json:"account_id"`
	CustomerID  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

// Validate checks that an a new account being created has
// a minimum amount of 5000
// a saving or checking type
func (r CreateAccountRequest) Validate() *errs.AppError {
	if minimumFunds(r.Amount) {
		return errs.NewValidationError("Minimum account amount is not met")
	}
	if validAccountType(r.AccountType) {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}

func minimumFunds(amount float64) bool {
	return amount < MINIMUM_FUNDS
}

func validAccountType(accountType string) bool {
	return strings.ToLower(accountType) != SAVING && strings.ToLower(accountType) != CHECKING
}