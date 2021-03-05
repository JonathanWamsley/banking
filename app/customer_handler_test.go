package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/jonathanwamsley/banking/dto"
	"github.com/jonathanwamsley/banking/errs"
	"github.com/jonathanwamsley/banking/mocks/service"
)

var router *mux.Router
var ch CustomerHandler
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandler{mockService}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func TestGetCustomersNoError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
	}
	mockService.EXPECT().GetAllCustomers().Return(dummyCustomers, nil)
	router.HandleFunc("/customers", ch.GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestGetCustomersInternalError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetAllCustomers().Return(nil, errs.NewUnexpectedError("some database error"))
	router.HandleFunc("/customers", ch.GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func TestGetCustomerError(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetCustomer("").Return(nil, errs.NewUnexpectedError("some database error"))
	router.HandleFunc("/customer", ch.GetCustomer)
	request, _ := http.NewRequest(http.MethodGet, "/customer", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func TestGetCustomerNoError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyCustomers := &dto.CustomerResponse{
		"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1",
	}
	mockService.EXPECT().GetCustomer("").Return(dummyCustomers, nil)
	router.HandleFunc("/customer", ch.GetCustomer)
	request, _ := http.NewRequest(http.MethodGet, "/customer", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestDeleteCustomerError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().DeleteCustomer("").Return(errs.NewUnexpectedError("some database error"))
	router.HandleFunc("/customer", ch.DeleteCustomer)
	request, _ := http.NewRequest(http.MethodDelete, "/customer", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func TestDeleteCustomerNoError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().DeleteCustomer("").Return(nil)
	router.HandleFunc("/customer", ch.DeleteCustomer)
	request, _ := http.NewRequest(http.MethodDelete, "/customer", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestCreateCustomerNoError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().DeleteCustomer("").Return(nil)
	router.HandleFunc("/customer", ch.DeleteCustomer)
	request, _ := http.NewRequest(http.MethodDelete, "/customer", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}
