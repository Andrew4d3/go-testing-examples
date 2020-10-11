package util

import (
	"testing"
)

// Traditional way
func Test_IsEven(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{"Even number", 4, true},
		{"Odd number", 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.number); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Alternative way
func Test_IsEven_Alternative(t *testing.T) {
	t.Run("Should return true if an even number is entered", func(t *testing.T) {
		if got := IsEven(4); got != true {
			t.Errorf("Got = %v, expected %v", got, true)
		}
	})

	t.Run("Should return false if an odd number is entered", func(t *testing.T) {
		if got := IsEven(3); got != false {
			t.Errorf("Got = %v, expected %v", got, false)
		}
	})
}
