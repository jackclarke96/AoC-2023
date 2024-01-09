package main

import "testing"

func TestEvaluateScore(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "executeMain - test example input 1 - unused pipes not included",
			input: `.....
			.F-7.
			.|.|.
			.L-J.
			.....`,
			expected: 4,
		},
		{
			name: "executeMain - test example input 1 - unused pipes included",
			input: `.-L|F7
			7S-7|
			L|7||
			-L-J|
			L|-JF`,
			expected: 4,
		},
		{
			name: "executeMain - test example input 2 - unused pipes included",
			input: `..F7.
			.FJ|.
			SJ.L7
			|F--J
			LJ...`,
			expected: 8,
		},
		{
			name: "executeMain - test example input 2 - unused pipes included",
			input: `7-F7-
			.FJ|7
			SJLL7
			|F--J
			LJ.LJ`,
			expected: 8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := executeMain(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
