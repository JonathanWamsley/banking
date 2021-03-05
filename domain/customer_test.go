package domain

import (
	"testing"

	"github.com/jonathanwamsley/banking/dto"
	"github.com/stretchr/testify/assert"
)

func TestStatusAsTextActive(t *testing.T) {
	c := Customer{Status: "1"}
	status := c.statusAsText()
	assert.EqualValues(t, "active", status)
}

func TestStatusAsTextInactive(t *testing.T) {
	c := Customer{Status: "0"}
	status := c.statusAsText()
	assert.EqualValues(t, "inactive", status)
}

func TestToDTO(t *testing.T) {
	c := Customer{
		ID:          "1234",
		Name:        "jon doe",
		City:        "city",
		Zipcode:     "123321",
		DateofBirth: "11/11/2000",
		Status:      "1",
	}
	cr := c.ToDTO()
	assert.EqualValues(t, "1234", cr.ID)
	assert.EqualValues(t, "jon doe", cr.Name)
	assert.EqualValues(t, "city", cr.City)
	assert.EqualValues(t, "123321", cr.Zipcode)
	assert.EqualValues(t, "11/11/2000", cr.DateofBirth)
	assert.EqualValues(t, "active", cr.Status)
}

func TestNewCustomer(t *testing.T) {
	cr := dto.CustomerRequest{
		Name:        "jon doe",
		City:        "city",
		Zipcode:     "123321",
		DateofBirth: "11/11/2000",
	}
	c := NewCustomer(cr)
	assert.EqualValues(t, "", c.ID)
	assert.EqualValues(t, "jon doe", c.Name)
	assert.EqualValues(t, "city", c.City)
	assert.EqualValues(t, "123321", c.Zipcode)
	assert.EqualValues(t, "11/11/2000", c.DateofBirth)
	assert.EqualValues(t, "1", c.Status)
}
