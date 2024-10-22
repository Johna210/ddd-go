// Package entites holds all the entities that are shared accross subdomains
package entity

import (
	"github.com/google/uuid"
)

type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
