package resolver

import (
	"../schema/types"
	"github.com/graph-gophers/graphql-go"
)

type dateResolver struct {
	data *types.Date
}

func (resolver *dateResolver) Month() graphql.ID {
	return resolver.data.Month
}

func (resolver *dateResolver) Day() graphql.ID {
	return resolver.data.Day
}

func (resolver *dateResolver) Year() graphql.ID {
	return resolver.data.Year
}
