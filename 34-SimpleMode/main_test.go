package main

import "testing"

func TestSimpleMode(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{10, 4, 5, 2, 4}, 4},
		{[]int{5, 10, 10, 6, 5}, 5},
		{[]int{5, 5, 2, 2, 1}, 5},
		{[]int{3, 4, 1, 6, 10}, -1},
		{[]int{1, 2, 3}, -1},
		{[]int{1}, -1},
	}

	for _, test := range tests {
		result := SimpleMode(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
