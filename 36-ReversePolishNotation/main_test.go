package main

import "testing"

func TestReversePolishNotation(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1 1 + 1 + 1 +", 4},
		{"4 5 + 2 1 + *", 27},
		{"2 12 + 7 /", 2},
		{"1 2 + 3 *", 9},
		{"10 2 /", 5},
	}

	for _, test := range tests {
		result := ReversePolishNotation(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
