package main

import "testing"

func TestEvaluateScore(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		problem  string
		expected int
	}{
		{
			name:     "Part 1 - 0 winning cards",
			input:    "game 1: 1 2 3 4 | 5 6 7 8 9 10",
			problem:  "1",
			expected: 0,
		},
		{
			name:     "Part 1 - 1 winning card",
			input:    "game 2: 1 2 3 4 | 4 6 7 8 9 10",
			problem:  "1",
			expected: 1,
		},
		{
			name:     "Part 1 - Multiple winning cards",
			input:    "game 3: 1 2 3 4 | 5 6 1 2 3 4",
			problem:  "1",
			expected: 8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := processGame(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
