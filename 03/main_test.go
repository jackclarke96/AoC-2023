package main

import "testing"

func TestIterateThroughEngine(t *testing.T) {
	testCases := []struct {
		name     string
		input    EngineSlice
		problem  string
		expected int
	}{
		// problem 1 tests
		{
			name: "Part 1 - Basic Test",
			input: EngineSlice{
				{"1", ".", "3", "4", "5", "", "7", ".", "9"},
				{"*", ".", ".", ".", ".", "#", ".", ".", "*"},
			},
			problem:  "1",
			expected: 362,
		},
		// problem 2
		{
			name: "Part 2; Single Gear; 1 adjacent number",
			input: EngineSlice{
				{"1", "2", ".", "4", "5", ".", ".", "8", "9"},
				{".", ".", ".", ".", ".", "*", ".", ".", "."},
			},
			problem:  "2",
			expected: 0,
		},
		{
			name: "Part 2; Single Gear; 2 adjacent numbers",
			input: EngineSlice{
				{"1", "2", ".", "4", "5", ".", "7", "8", "9"},
				{".", ".", ".", ".", ".", "*", ".", ".", "."},
			},
			problem:  "2",
			expected: 35505,
		},
		{
			name: "Part 2 - Single Gear - 3 adjacent numbers",
			input: EngineSlice{
				{"1", "2", ".", "4", "5", ".", "7", "8", "9"},
				{".", ".", ".", ".", ".", "*", ".", ".", "."},
				{"1", "2", ".", "4", "5", ".", ".", "8", "9"},
			},
			problem:  "2",
			expected: 0,
		},
		{
			name: "Part 2 - Single Gear - 4 adjacent numbers",
			input: EngineSlice{
				{"1", "2", ".", "4", "5", ".", "7", "8", "9"},
				{".", ".", ".", ".", ".", "*", ".", ".", "."},
				{"1", "2", ".", "4", "5", ".", "7", "8", "9"},
			},
			problem:  "2",
			expected: 0,
		},
		{
			name: "Part 2 - Single Gear - diagonally adjacent numbers",
			input: EngineSlice{
				{"1", "2", ".", "4", "5", ".", ".", "8", "9"},
				{".", ".", ".", ".", ".", "*", ".", ".", "."},
				{"1", "2", ".", "4", ".", ".", "7", "8", "9"},
			},
			problem:  "2",
			expected: 45 * 789,
		},
		{
			name: "Part 2 - Single Gear - Above and Beneath",
			input: EngineSlice{
				{"1", "2", ".", "4", "5", "6", ".", "8", "9"},
				{".", ".", ".", ".", ".", "*", ".", ".", "."},
				{"1", "2", ".", "4", ".", "6", "7", "8", "9"},
			},
			problem:  "2",
			expected: 456 * 6789,
		},

		{
			name: "Part 2 - Single Gear - Left and Right",
			input: EngineSlice{
				{"1", "2", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", "4", "5", "*", "7", "8", "."},
				{"1", "2", ".", ".", ".", ".", ".", ".", "9"},
			},
			problem:  "2",
			expected: 45 * 78,
		},

		{
			name: "Part 2 - Single Gear - Multiple Gears",
			input: EngineSlice{
				{"1", "2", ".", ".", ".", ".", ".", ".", "."},
				{".", "*", ".", "4", "5", "*", "7", "8", "."},
				{"1", "2", ".", ".", ".", ".", ".", ".", "9"},
			},
			problem:  "2",
			expected: 12*12 + 45*78,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.iterateThroughEngine(tc.problem)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
