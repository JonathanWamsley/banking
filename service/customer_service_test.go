package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	realdomain "github.com/jonathanwamsley/banking/domain"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
	"github.com/jonathanwamsley/banking/mocks/domain"
	"github.com/stretchr/testify/assert"
)

var mockRepo *domain.MockCustomerRepository
var service CustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockCustomerRepository(ctrl)
	service = NewCustomerService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func TestCreateCustomerValidationError(t *testing.T) {
	request := dto.CustomerRequest{
		Name:        "",
		City:        "",
		Zipcode:     "",
		DateofBirth: "",
	}

	service := NewCustomerService(nil)

	resp, err := service.CreateCustomer(request)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

func TestCreateCustomerServerError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.CustomerRequest{
		Name:        "jon doe",
		City:        "city",
		Zipcode:     "123321",
		DateofBirth: "11/11/2000",
	}
	customer := realdomain.Customer{
		ID:          "",
		Name:        req.Name,
		City:        req.City,
		Zipcode:     req.Zipcode,
		DateofBirth: req.DateofBirth,
		Status:      "1",
	}

	mockRepo.EXPECT().Save(customer).Return(nil, errs.NewUnexpectedError("unexpected database error"))

	customerResponse, err := service.CreateCustomer(req)
	assert.Nil(t, customerResponse)
	assert.NotNil(t, err)
}

func TestCreateCustomerNoError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.CustomerRequest{
		Name:        "jon doe",
		City:        "city",
		Zipcode:     "123321",
		DateofBirth: "11/11/2000",
	}
	customer := realdomain.Customer{
		ID:          "",
		Name:        req.Name,
		City:        req.City,
		Zipcode:     req.Zipcode,
		DateofBirth: req.DateofBirth,
		Status:      "1",
	}

	customerOut := customer
	customerOut.ID = "1234"

	mockRepo.EXPECT().Save(customer).Return(&customerOut, nil)

	customerResponse, err := service.CreateCustomer(req)
	assert.Nil(t, err)
	assert.NotNil(t, customerResponse)
	assert.EqualValues(t, "1234", customerResponse.ID)
	assert.EqualValues(t, "jon doe", customerResponse.Name)
	assert.EqualValues(t, "city", customerResponse.City)
	assert.EqualValues(t, "123321", customerResponse.Zipcode)
	assert.EqualValues(t, "11/11/2000", customerResponse.DateofBirth)
	assert.EqualValues(t, "active", customerResponse.Status) // To Dto changes 1 to active
}

func TestGetAllCustomersNoError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// resp := []dto.CustomerRequest{
	// 	{
	// 		Name: "jon doe",
	// 		City: "city",
	// 		Zipcode: "123321",
	// 		DateofBirth: "11/11/2000",
	// 	},
	// 	{
	// 		Name: "jon smith",
	// 		City: "city2",
	// 		Zipcode: "123456",
	// 		DateofBirth: "12/12/2012",
	// 	},
	// }

	customers := []realdomain.Customer{
		{
			ID:          "1234",
			Name:        "jon doe",
			City:        "city",
			Zipcode:     "123321",
			DateofBirth: "11/11/2000",
			Status:      "1",
		},
		{
			ID:          "1235",
			Name:        "jon smith",
			City:        "city2",
			Zipcode:     "123456",
			DateofBirth: "12/12/2012",
			Status:      "1",
		},
	}

	mockRepo.EXPECT().FindAll().Return(customers, nil)

	customersResponse, err := service.GetAllCustomers()
	assert.Nil(t, err)
	assert.NotNil(t, customersResponse)
	assert.EqualValues(t, 2, len(customersResponse))

	assert.EqualValues(t, customersResponse[0].ID, "1234")
	assert.EqualValues(t, customersResponse[0].Name, "jon doe")
	assert.EqualValues(t, customersResponse[0].City, "city")
	assert.EqualValues(t, customersResponse[0].Zipcode, "123321")
	assert.EqualValues(t, customersResponse[0].DateofBirth, "11/11/2000")
	assert.EqualValues(t, customersResponse[0].Status, "active")

	assert.EqualValues(t, customersResponse[1].ID, "1235")
	assert.EqualValues(t, customersResponse[1].Name, "jon smith")
	assert.EqualValues(t, customersResponse[1].City, "city2")
	assert.EqualValues(t, customersResponse[1].Zipcode, "123456")
	assert.EqualValues(t, customersResponse[1].DateofBirth, "12/12/2012")
	assert.EqualValues(t, customersResponse[1].Status, "active")
}

func TestGetAllCustomersServerError(t *testing.T) {

	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().FindAll().Return(nil, errs.NewUnexpectedError("unexpected database error"))

	customerResponse, err := service.GetAllCustomers()
	assert.Nil(t, customerResponse)
	assert.NotNil(t, err)
}

func TestGetCustomerNoError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	customer := &realdomain.Customer{
		ID:          "1234",
		Name:        "jon doe",
		City:        "city",
		Zipcode:     "123321",
		DateofBirth: "11/11/2000",
		Status:      "1",
	}

	mockRepo.EXPECT().ByID("1234").Return(customer, nil)

	c, err := service.GetCustomer("1234")
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.EqualValues(t, "1234", c.ID)
	assert.EqualValues(t, "jon doe", c.Name)
	assert.EqualValues(t, "city", c.City)
	assert.EqualValues(t, "123321", c.Zipcode)
	assert.EqualValues(t, "11/11/2000", c.DateofBirth)
	assert.EqualValues(t, "active", c.Status) // dto transformation
}

func TestGetCustomerSeverError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().ByID("1234").Return(nil, errs.NewUnexpectedError("unexpected database error"))

	customerResponse, err := service.GetCustomer("1234")
	assert.Nil(t, customerResponse)
	assert.NotNil(t, err)
}

func TestGetCustomerNotFoundError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().ByID("").Return(nil, errs.NewNotFoundError("unexpected database error"))

	customerResponse, err := service.GetCustomer("")
	assert.Nil(t, customerResponse)
	assert.NotNil(t, err)
}

func TestDeleteCustomerServerError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	customer := &realdomain.Customer{
		ID:          "1234",
		Name:        "jon doe",
		City:        "city",
		Zipcode:     "123321",
		DateofBirth: "11/11/2000",
		Status:      "1",
	}

	mockRepo.EXPECT().ByID("1234").Return(customer, nil)
	mockRepo.EXPECT().Delete("1234").Return(errs.NewUnexpectedError("unexpected database error"))

	err := service.DeleteCustomer("1234")
	assert.NotNil(t, err)
}

func TestDeleteCustomerNoError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	customer := &realdomain.Customer{
		ID:          "1234",
		Name:        "jon doe",
		City:        "city",
		Zipcode:     "123321",
		DateofBirth: "11/11/2000",
		Status:      "1",
	}

	mockRepo.EXPECT().ByID("1234").Return(customer, nil)
	mockRepo.EXPECT().Delete("1234").Return(nil)

	err := service.DeleteCustomer("1234")
	assert.Nil(t, err)
}

func TestNewCustomerService(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	customerService := NewCustomerService(mockRepo)
	assert.NotNil(t, customerService)
}
