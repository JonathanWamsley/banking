package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumAmount(t *testing.T) {
	assert.True(t, minimumFunds(MINIMUM_FUNDS-1))
	assert.False(t, minimumFunds(MINIMUM_FUNDS))
	assert.False(t, minimumFunds(MINIMUM_FUNDS+1))
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, 5000.0, MINIMUM_FUNDS)
	assert.EqualValues(t, "saving", SAVING)
	assert.EqualValues(t, "checking", CHECKING)
}

func TestValidAccountType(t *testing.T) {
	assert.False(t, validAccountType("saving"))
	assert.False(t, validAccountType("SAVING"))
	assert.False(t, validAccountType("checking"))
	assert.False(t, validAccountType("CHECKING"))
	assert.True(t, validAccountType("checkings"))
}

func TestValidateNoErrors(t *testing.T) {
	a := CreateAccountRequest{Amount: MINIMUM_FUNDS, AccountType: CHECKING}
	err := a.Validate()
	assert.Nil(t, err)
}

func TestValidateBadAmount(t *testing.T) {
	a := CreateAccountRequest{Amount: MINIMUM_FUNDS - 1}
	err := a.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, 422, err.Code)
	assert.EqualValues(t, "Minimum account amount is not met", err.Message)
}

func TestValidateBadAccountType(t *testing.T) {
	a := CreateAccountRequest{Amount: MINIMUM_FUNDS, AccountType: ""}
	err := a.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, 422, err.Code)
	assert.EqualValues(t, "Account type should be checking or saving", err.Message)
}
