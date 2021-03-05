package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInvalidName(t *testing.T) {
	cr := CustomerRequest{"", "", "", ""}
	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, 422, err.Code)
	assert.EqualValues(t, "invalid name", err.Message)
}

func TestValidateInvalidCity(t *testing.T) {
	cr := CustomerRequest{"name", "", "", ""}
	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, 422, err.Code)
	assert.EqualValues(t, "invalid city", err.Message)
}

func TestValidateInvalidZipcode(t *testing.T) {
	cr := CustomerRequest{"name", "city", "", ""}
	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, 422, err.Code)
	assert.EqualValues(t, "invalid zipcode", err.Message)
}

func TestValidateInvalidDateOfBirth(t *testing.T) {
	cr := CustomerRequest{"name", "city", "32123", ""}
	err := cr.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, 422, err.Code)
	assert.EqualValues(t, "invalid date of birth", err.Message)
}

func TestValidateNoError(t *testing.T) {
	cr := CustomerRequest{"name", "city", "321123", "01/01/2000"}
	err := cr.Validate()
	assert.Nil(t, err)
}
