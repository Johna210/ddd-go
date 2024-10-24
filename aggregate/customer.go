// package aggreagate holds our aggregates that combines many entites that combines
// into a full object
package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/johna210/ddd-go/entity"
	valueobject "github.com/johna210/ddd-go/value_object"
)

var (
	ErrInvalidPerson = errors.New("a customer needs to have a valid name")
)

type Customer struct {
	// person is the root entity of the customer
	// which means person.ID is the main indentifier for the customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new customer aggregate
// it will validate the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}
