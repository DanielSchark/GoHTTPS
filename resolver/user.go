package resolver

import (
	"../schema/types"
	"github.com/graph-gophers/graphql-go"
)

type userArgs struct {
	ID        graphql.ID
	Firstname *string
	Lastname  *string
	Email     *string
}

// Users resolver
func (resolver *Resolver) Users() []*userResolver {
	var userList []*userResolver
	for _, user := range mockUsersData {
		userList = append(userList, &userResolver{user})
	}
	return userList
}

type usersResolver struct {
	data *types.Users
}

// User resolver
func (resolver *Resolver) User(args userArgs) *userResolver {
	if user := mockUsersData[args.ID]; user != nil {
		return &userResolver{user}
	}
	return nil
}

type userResolver struct {
	data *types.User
}

func (resolver *userResolver) ID() graphql.ID {
	return resolver.data.ID
}
func (resolver *userResolver) Name() *nameResolver {
	return &nameResolver{&resolver.data.Name}
}
func (resolver *userResolver) Gender() string {
	return resolver.data.Gender
}
func (resolver *userResolver) Birthdate() *dateResolver {
	return &dateResolver{&resolver.data.Birthdate}
}
func (resolver *userResolver) Phone() *string {
	return &resolver.data.Phone
}
func (resolver *userResolver) Email() *string {
	return &resolver.data.Email
}
func (resolver *userResolver) Jobs() *[]string {
	list := make([]string, len(resolver.data.Jobs))
	for index, tmp := range resolver.data.Jobs {
		list[index] = tmp
	}
	return &list
}
func (resolver *userResolver) Created() graphql.Time {
	return graphql.Time{Time: resolver.data.Created}
}
func (resolver *userResolver) Edited() *[]graphql.Time {
	list := make([]graphql.Time, len(resolver.data.Edited))
	for index, tmp := range resolver.data.Edited {
		list[index] = graphql.Time{Time: tmp}
	}
	return &list
}
