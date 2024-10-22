// Package memory is an in-memory implementation of the Customer repository
package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/johna210/ddd-go/aggregate"
	"github.com/johna210/ddd-go/domain/customer"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure the customer is already in repo
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer alread exists :%w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer doesnt exists: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}
