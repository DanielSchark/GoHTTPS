package schema

var Schema = `
scalar Time
scalar ID

type Name {
	first: String!
	last: String
}

type Date {
	month: ID!
	day: ID!
	year: ID!
}

type User {
	id: ID!
	name: Name!
	gender: String!
	birthdate: Date
	phone: String
	email: String
	jobs: [String!]
	created: Time!
	edited: [Time!]
}

type Query {
	user(
		id: ID!
		firstname: String
		lastname: String
		email: String
	): User
	users(): [User]!
}

schema {
	# mutation: Mutation
	query: Query
}
`
