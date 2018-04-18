package resolver

import (
	"time"

	"../schema/types"
	"github.com/graph-gophers/graphql-go"
)

var mockUsers = []*types.User{
	{
		ID: "1",
		Name: types.Name{
			First: "Daniel",
			Last:  "Sharkov",
		},
		Gender: "Male",
		Birthdate: types.Date{
			Month: "9",
			Day:   "9",
			Year:  "1999",
		},
		Phone: "017692582954",
		Email: "scharktv@gmail.com",
		Jobs: []string{
			"Programmer",
		},
		Created: time.Now(),
		Edited: []time.Time{
			time.Now(),
		},
	},
	{
		ID: "4",
		Name: types.Name{
			First: "Paul",
			Last:  "Streetz",
		},
		Gender: "Male",
		Birthdate: types.Date{
			Month: "6",
			Day:   "22",
			Year:  "1997",
		},
		Phone: "017692539635",
		Email: "streezy@somewhere.com",
		Jobs: []string{
			"Programmer",
		},
		Created: time.Now(),
		Edited: []time.Time{
			time.Now(),
			time.Now(),
			time.Now(),
			time.Now(),
			time.Now(),
		},
	},
	{
		ID: "3",
		Name: types.Name{
			First: "Alice",
			Last:  "Wondalan",
		},
		Gender: "Female",
		Birthdate: types.Date{
			Month: "9",
			Day:   "28",
			Year:  "1991",
		},
		Phone: "017682952063",
		Email: "alice@wonderland.rich",
		Jobs: []string{
			"Graphic Designer",
		},
		Created: time.Now(),
		Edited: []time.Time{
			time.Now(),
		},
	},
}
var mockUsersData = make(map[graphql.ID]*types.User)

func init() {
	for _, user := range mockUsers {
		mockUsersData[user.ID] = user
	}
}
