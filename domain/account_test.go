package domain

import (
	"testing"

	"github.com/jonathanwamsley/banking/dto"
	"github.com/stretchr/testify/assert"
)

func TestToCreateAccountResponseDTO(t *testing.T) {
	a := Account{AccountID: "123"}
	resp := a.ToCreateAccountResponseDTO()
	assert.Equal(t, "123", resp.AccountID)
}

func TestToGetAccountResponseDTO(t *testing.T) {
	a := Account{
		AccountID: "123",
		CustomerID: "456",
		OpeningDate: "11-11-1111",
		AccountType: "checking",
		Amount: 1000,
	}
	resp := a.ToGetAccountResponseDTO()
	assert.Equal(t, "123", resp.AccountID)
	assert.Equal(t, "456", resp.CustomerID)
	assert.Equal(t, "11-11-1111", resp.OpeningDate)
	assert.Equal(t, "checking", resp.AccountType)
	assert.Equal(t, 1000.0, resp.Amount)
}

func TestNewAccount(t *testing.T) {
	a := dto.CreateAccountRequest{
		CustomerID: "456",
		AccountType: "checking",
		Amount: 1000,
	}
	resp := NewAccount(a)
	assert.Equal(t, "456", resp.CustomerID)
	assert.NotNil(t, resp.OpeningDate)
	assert.Equal(t, "checking", resp.AccountType)
	assert.Equal(t, 1000.0, resp.Amount)
	assert.Equal(t, "1", resp.Status)
}

func TestCanWithdraw(t *testing.T) {
	a := Account{Amount: 1000.0}
	assert.True(t, a.CanWithdraw(1000.0))
	assert.True(t, a.CanWithdraw(900.0))
	assert.False(t, a.CanWithdraw(2000.0))

}