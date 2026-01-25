package main

import "testing"

func TestDistinctList(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 2, 2, 3}, 2},
		{[]int{0, -2, -2, 5, 5, 5}, 3},
		{[]int{100, 2, 101, 4}, 0},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := DistinctList(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
