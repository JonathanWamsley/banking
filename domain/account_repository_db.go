package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/jonathanwamsley/banking/errs"
	"github.com/jonathanwamsley/banking/logger"
)

// The query statements
const (
	createAccount   = "insert into accounts(customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?);"
	getAccounts     = "select account_id, customer_id, opening_date, account_type, amount from accounts where customer_id = ?;"
	deleteAccount   = "delete from accounts where customer_id = ? and account_type = ?;"
	getAccount      = "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?;"
	makeTransaction = "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?);"
)

// AccountRepositoryDB holds the sql client connection
type AccountRepositoryDB struct {
	client *sqlx.DB
}

// NewAccountRepositoryDB creates a new AccountRepositoryDB to call sql methods
func NewAccountRepositoryDB(repo *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{repo}
}

// Save creates a new account for a customer. The account id is returned
func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	result, err := d.client.Exec(createAccount, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("error while creating new account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last id from the new account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

// ByID returns all the accounts of a customers id from the database
func (d AccountRepositoryDB) ByID(id string) ([]Account, *errs.AppError) {
	accounts := make([]Account, 0)
	err := d.client.Select(&accounts, getAccounts, id)

	if err != nil {
		if err == sql.ErrNoRows {
			// no need to log queries about missing customers
			return nil, errs.NewNotFoundError("Customer has no accounts")
		}
		logger.Error("Error while querying account table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return accounts, nil
}

// Delete returns nil for a customer of a given account type
func (d AccountRepositoryDB) Delete(id string, accountType string) *errs.AppError {
	result, err := d.client.Exec(deleteAccount, id, accountType)
	if err != nil {
		logger.Error("Error while trying to delete account " + err.Error())
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	rowsChanged, _ := result.RowsAffected()
	if rowsChanged == 0 {
		return errs.NewNotFoundError("Account not found")
	}
	return nil
}

// SaveTransaction completes a withdrawal or deposit in a bank account. A new total will be returned.
func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(makeTransaction, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountID)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionID, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountID)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionID = strconv.FormatInt(transactionID, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

// FindBy returns a specific account information given the account id
func (d AccountRepositoryDB) FindBy(accountID string) (*Account, *errs.AppError) {
	var account Account
	err := d.client.Get(&account, getAccount, accountID)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}
