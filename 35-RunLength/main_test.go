package main

import "testing"

func TestRunLength(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"wwwggopp", "3w2g1o2p"},
		{"aabbcde", "2a2b1c1d1e"},
		{"wwwbbbw", "3w3b1w"},
		{"a", "1a"},
		{"", ""},
	}

	for _, test := range tests {
		result := RunLength(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %q but got %q", test.input, test.expected, result)
		}
	}
}
