package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/jonathanwamsley/banking/errs"
	"github.com/jonathanwamsley/banking/logger"
)

// The query statements
const (
	CreateAccount = "insert into accounts(customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?);"
)

// AccountRepositoryDB holds the sql client connection
type AccountRepositoryDB struct {
	client *sqlx.DB
}

// NewAccountRepositoryDB creates a new AccountRepositoryDB to call sql methods
func NewAccountRepositoryDB(repo *sqlx.DB) AccountRepositoryDB{
	return AccountRepositoryDB{repo}
}

// Save creates a new account for a customer. The account id is returned
func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	result, err := d.client.Exec(CreateAccount, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
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