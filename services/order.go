package services

import (
	"github.com/johna210/ddd-go/domain/customer"
	"github.com/johna210/ddd-go/domain/customer/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRespository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Loop through all the cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRespository) OrderConfiguration {
	// Return a function that matches the orderconfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}
