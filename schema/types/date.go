package types

import "github.com/graph-gophers/graphql-go"

type Date struct {
	Month graphql.ID `json:"month,omitempty"`
	Day   graphql.ID `json:"day,omitempty"`
	Year  graphql.ID `json:"year,omitempty"`
}
