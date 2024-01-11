package main

import "testing"

func TestEvaluateScore(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "executeMain - test example input",
			input: `...#......
			.......#..
			#.........
			..........
			......#...
			.#........
			.........#
			..........
			.......#..
			#...#.....
			`,
			expected: 374,
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

func TestCalculateDistance(t *testing.T) {
	testCases := []struct {
		name     string
		input    calculationInput
		expected int
	}{
		{
			name: "calculateDistance - test example input 1 - galaxy 1 and galaxy 7",
			input: calculationInput{
				galaxyStruct{0, 3},
				galaxyStruct{8, 7},
				2,
				2,
			},
			expected: 15,
		},
		{
			name: "calculateDistance - test example input 2 - galaxy 3 and galaxy 6",
			input: calculationInput{
				galaxyStruct{2, 0},
				galaxyStruct{6, 9},
				1,
				3,
			},
			expected: 17,
		},
		{
			name: "calculateDistance - test example input 2 - galaxy 8 and galaxy 9",
			input: calculationInput{
				galaxyStruct{9, 0},
				galaxyStruct{9, 4},
				0,
				1,
			},
			expected: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculateDistance(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
