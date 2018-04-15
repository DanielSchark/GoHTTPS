package graphqlSchema

import "github.com/graphql-go/graphql"

type User struct {
	Id        int `json:"id,omitempty"`
	Name      `json:"name,omitempty"`
	Birthdate Date     `json:"birthdate,omitempty"`
	Jobs      []string `json:"jobs,omitempty"`
	Phone     string   `json:"phone,omitempty"`
}

type Name struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}

type Date struct {
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
	Year  int `json:"year,omitempty"`
}

var (
	UserSchema = graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "User",
			Description: "Representing an user",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.ID,
				},
				"Name": &graphql.Field{
					Type: NameSchema,
				},
				"Birthdate": &graphql.Field{
					Type: DateSchema,
				},
				"Jobs": &graphql.Field{
					Type: graphql.String,
				},
				"Phone": &graphql.Field{
					Type: graphql.String,
				},
				"Email": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	DateSchema = graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Date",
			Description: "Date",
			Fields: graphql.Fields{
				"Month": &graphql.Field{
					Type: graphql.Int,
				},
				"Day": &graphql.Field{
					Type: graphql.Int,
				},
				"Year": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
	NameSchema = graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Name",
			Description: "Name",
			Fields: graphql.Fields{
				"First": &graphql.Field{
					Type: graphql.String,
				},
				"Last": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)

var testUsers = []User{
	User{
		Id: 25,
		Name: Name{
			First: "Daniel",
			Last:  "Sharkov",
		},
		Birthdate: Date{
			Month: 3,
			Day:   16,
			Year:  2001,
		},
		Jobs: []string{
			"CEO",
			"Programmer",
		},
		Phone: "017662845232",
	},
	User{
		Id: 836,
		Name: Name{
			First: "Bob",
			Last:  "Ross",
		},
		Birthdate: Date{
			Month: 4,
			Day:   14,
			Year:  1984,
		},
		Jobs: []string{
			"Painter",
		},
		Phone: "017634672633",
	},
	User{
		Id: 9,
		Name: Name{
			First: "Steven",
			Last:  "Jobby",
		},
		Birthdate: Date{
			Month: 6,
			Day:   6,
			Year:  1936,
		},
		Phone: "01764628542",
	},
}
