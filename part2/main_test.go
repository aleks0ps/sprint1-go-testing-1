package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullUsername(t *testing.T) {
	tests := []struct {
		name     string
		fullname User
		want     string
	}{
		{
			name: "Have FirstName and LastName",
			fullname: User{
				FirstName: "Alex",
				LastName:  "Unknown",
			},
			want: "Alex Unknown",
		},
		{
			name: "Have only FirstName",
			fullname: User{
				FirstName: "Alex",
				LastName:  "",
			},
			want: "Alex",
		},
		{
			name: "Empty name",
			fullname: User{
				FirstName: "",
				LastName:  "",
			},
			want: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.fullname.FullName(), test.want)
		})
	}
}
