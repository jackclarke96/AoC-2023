package main

import "testing"

var s2f = `soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15`

func TestEvaluateScore(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "makePrediction - test example input 3",
			input:    []string{"10", "13", "16", "21", "30", "45"},
			expected: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := makePrediction(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
