package main

import (
	"encoding/json"
	"testing"
)

func TestProcess(t *testing.T) {
	inputJSON := `{
		"b": "val",
		"a": "val",
		"c": null,
		"d": "",
		"list": [
			{"x": 1, "y": 2},
			{"y": 2, "x": 1},
			{"z": 3}
		],
		"nested": {
			"e": null,
			"f": ""
		}
	}`

	var data interface{}
	if err := json.Unmarshal([]byte(inputJSON), &data); err != nil {
		t.Fatalf("Failed to unmarshal input: %v", err)
	}

	processed := Process(data)

	// Marshal to check the final JSON string (which verifies sorting and structure)
	resultBytes, err := json.Marshal(processed)
	if err != nil {
		t.Fatalf("Failed to marshal result: %v", err)
	}
	resultStr := string(resultBytes)

	// Expected JSON structure
	// Keys should be sorted.
	// "a", "b", "list", "nested"
	// list: [{"x":1,"y":2},{"z":3}] (x, y sorted)
	// nested: {}
	
	expectedStr := `{"a":"val","b":"val","list":[{"x":1,"y":2},{"z":3}],"nested":{}}`
	
	if resultStr != expectedStr {
		t.Errorf("Expected %s, got %s", expectedStr, resultStr)
	}
}
