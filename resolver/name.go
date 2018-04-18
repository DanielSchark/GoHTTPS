package resolver

import "../schema/types"

type nameResolver struct {
	data *types.Name
}

func (resolver *nameResolver) First() string {
	return resolver.data.First
}

func (resolver *nameResolver) Last() *string {
	return &resolver.data.Last
}
