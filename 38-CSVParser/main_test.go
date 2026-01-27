package main

import (
	"math"
	"testing"
)

func TestSummarizeCSV(t *testing.T) {
	t.Run("valid csv", func(t *testing.T) {
		csvData := `city,jan,feb,mar
Berlin,1.2,2.3,5
Paris,3.4,4.5,6.7
Milan,2,3,4`

		expected := []summary{
			{Column: "city", Sum: 0},
			{Column: "jan", Sum: 6.6},
			{Column: "feb", Sum: 9.8},
			{Column: "mar", Sum: 15.7},
		}

		result, err := SummarizeCSV([]byte(csvData))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		compareSummaries(t, expected, result)
	})

	t.Run("invalid csv with non-numeric value", func(t *testing.T) {
		csvData := `city,jan,feb,mar
Berlin,1.2,2.3,5
Paris,3.4,4.5,6.7
Madrid,5.O,6.1,7.2
Milan,2,3,4`

		// Expected sums:
		// jan: 1.2 + 3.4 + (skip 5.O) + 2 = 6.6
		// feb: 2.3 + 4.5 + 6.1 + 3 = 15.9
		// mar: 5 + 6.7 + 7.2 + 4 = 22.9
		expected := []summary{
			{Column: "city", Sum: 0},
			{Column: "jan", Sum: 6.6},
			{Column: "feb", Sum: 15.9},
			{Column: "mar", Sum: 22.9},
		}

		result, err := SummarizeCSV([]byte(csvData))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		compareSummaries(t, expected, result)
	})
}

func compareSummaries(t *testing.T, expected, got []summary) {
	t.Helper()
	if len(expected) != len(got) {
		t.Fatalf("expected length %d, got %d", len(expected), len(got))
	}

	const epsilon = 1e-9

	for i := range expected {
		if expected[i].Column != got[i].Column {
			t.Errorf("index %d: expected column %q, got %q", i, expected[i].Column, got[i].Column)
		}
		if math.Abs(expected[i].Sum-got[i].Sum) > epsilon {
			t.Errorf("index %d (%s): expected sum %v, got %v", i, expected[i].Column, expected[i].Sum, got[i].Sum)
		}
	}
}
