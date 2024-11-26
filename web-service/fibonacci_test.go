package main

import (
	"reflect"
	"testing"
)

func TestFibonacciSequence(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{"negative number", -1, []int{}},
		{"zero", 0, []int{}},
		{"one", 1, []int{0}},
		{"two", 2, []int{0, 1}},
		{"five", 5, []int{0, 1, 1, 2, 3}},
		{"ten", 10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fibonacciSequence(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("fibonacciSequence(%d) = %v; want %v", tt.input, result, tt.expected)
			}

		})
	}
}
