package dto

import (
	"strings"

	"github.com/jonathanwamsley/banking/errs"
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
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
