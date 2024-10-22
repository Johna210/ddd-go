package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/johna210/ddd-go/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type CustomerRespository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
