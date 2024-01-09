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
			name:     "makePrediction - test example input 1",
			input:    []string{"0", "3", "6", "9", "12", "15"},
			expected: 18,
		},
		{
			name:     "makePrediction - test example input 2",
			input:    []string{"1", "3", "6", "10", "15", "21"},
			expected: 28,
		},
		{
			name:     "makePrediction - test example input 3",
			input:    []string{"10", "13", "16", "21", "30", "45"},
			expected: 68,
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
