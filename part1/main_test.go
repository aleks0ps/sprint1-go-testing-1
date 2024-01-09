package main

import (
	"testing"
)

// -3, 3, -2.000001, -0.000000003
func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Test for 3",
			value: float64(3),
			want:  float64(3),
		},
		{
			name:  "Test for -3",
			value: float64(-3),
			want:  float64(3),
		},
		{
			name:  "Test for -2.000001",
			value: -2.000001,
			want:  2.000001,
		},
		{
			name:  "Test for -0.000000003",
			value: -0.000000003,
			want:  0.000000003,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if abs := Abs(test.value); abs != test.want {
				t.Errorf("Abs() = %f, want %f", abs, test.want)
			}
		})
	}
}
