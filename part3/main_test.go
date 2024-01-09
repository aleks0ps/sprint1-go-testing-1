package main

import "testing"

func TestAddNew(t *testing.T) {
	family := Family{}
	two_fathers := []struct {
		name   string
		rel    Relationship
		person Person
	}{
		{
			name: "Father 1",
			rel:  Father,
			person: Person{
				FirstName: "David",
				LastName:  "Furnish",
				Age:       61,
			},
		},
		{
			name: "Father 2",
			rel:  Father,
			person: Person{
				FirstName: "Elton",
				LastName:  "John",
				Age:       76,
			},
		},
	}

	for _, test := range two_fathers {
		t.Run(test.name, func(t *testing.T) {
			if err := family.AddNew(test.rel, test.person); err != nil {
				t.Errorf("%v", ErrRelationshipAlreadyExists)
			}
		})
	}
}
