package dto

import "github.com/jonathanwamsley/banking/errs"

// transaction types
const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

// MakeTransactionRequest fields to store a transaction
type MakeTransactionRequest struct {
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerID      string  `json:"-"`
}

// IsTransactionTypeWithdrawal checks for withdrawal type
func (r MakeTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

// IsTransactionTypeDeposit checks for deposit type
func (r MakeTransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

// Validate make sure transaction type is correct and amount withdrawal does is not great than account balance
func (r MakeTransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

// MakeTransactionResponse dto requirements
type MakeTransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
