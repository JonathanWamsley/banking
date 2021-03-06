package domain

import "github.com/jonathanwamsley/banking/dto"

// WITHDRAWAL is an transaction type
const WITHDRAWAL = "withdrawal"

// Transaction holds requirements to do a bank transaction
type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

// IsWithdrawal checks transaction type
func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

// ToDTO converts transaction to the transaction response for the user
func (t Transaction) ToDTO() dto.MakeTransactionResponse {
	return dto.MakeTransactionResponse{
		TransactionID:   t.TransactionID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
