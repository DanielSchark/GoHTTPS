package types

import (
	"time"

	"github.com/graph-gophers/graphql-go"
)

// Represents the type of user
type User struct {
	ID        graphql.ID  `json:"id,omitempty"`
	Name      Name        `json:"name,omitempty"`
	Gender    string      `json:"gender,omitempty"`
	Birthdate Date        `json:"birthdate,omitempty"`
	Phone     string      `json:"phone,omitempty"`
	Email     string      `json:"email,omitempty"`
	Jobs      []string    `json:"jobs,omitempty"`
	Created   time.Time   `json:"created,omitempty"`
	Edited    []time.Time `json:"edited,omitempty"`
}

// Represents a slice of type user
type Users []User
