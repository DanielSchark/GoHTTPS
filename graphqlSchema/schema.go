package graphqlSchema

import (
	"github.com/graphql-go/graphql"
)

var data map[string]string

var (
	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"User": &graphql.Field{
				Type: UserSchema,
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
					"Name": &graphql.ArgumentConfig{
						Type: NameSchema,
					},
					"Birthdate": &graphql.ArgumentConfig{
						Type: DateSchema,
					},
					"Jobs": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Phone": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(
					params graphql.ResolveParams,
				) (interface{}, error) {
					idQurey, isOK := params.Args["id"].(string)
					if isOK {
						return data[idQurey], nil
					}
					return nil, nil
				},
			},
			"Test": &graphql.Field{
				Type: graphql.String,
				Resolve: func(
					params graphql.ResolveParams,
				) (interface{}, error) {
					idQurey, isOK := params.Args["id"].(string)
					if isOK {
						return data[idQurey], nil
					}
					return nil, nil
				},
			},
		},
	})
	Schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
)
