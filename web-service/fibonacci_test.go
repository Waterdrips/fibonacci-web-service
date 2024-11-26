package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestFibonacciSequence(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    []int
		expectedErr error
	}{
		{"negative number", -1, []int{}, nil},
		{"zero", 0, []int{}, nil},
		{"one", 1, []int{0}, nil},
		{"two", 2, []int{0, 1}, nil},
		{"five", 5, []int{0, 1, 1, 2, 3}, nil},
		{"ten", 10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, nil},
		{"massive", 9999999999999999, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, errors.New("array index out of bounds")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := fibonacciSequence(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("fibonacciSequence(%d) = %v; want %v", tt.input, result, tt.expected)
			}

		})
	}
}
