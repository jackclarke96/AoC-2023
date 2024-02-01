package main

import (
	"os"
	"testing"
)

var s2f = `soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15`

func TestExecuteMain(t *testing.T) {
	testInput, err := os.ReadFile("../input.txt")
	if err != nil {
		t.Errorf("failed to read input.txt file: %v", err)
	}

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Part 2 - test input",
			input: `seeds: 79 14 55 13

			seed-to-soil map:
			50 98 2
			52 50 48

			soil-to-fertilizer map:
			0 15 37
			37 52 2
			39 0 15

			fertilizer-to-water map:
			49 53 8
			0 11 42
			42 0 7
			57 7 4

			water-to-light map:
			88 18 7
			18 25 70

			light-to-temperature map:
			45 77 23
			81 45 19
			68 64 13

			temperature-to-humidity map:
			0 69 1
			1 0 69

			humidity-to-location map:
			60 56 37
			56 93 4`,
			expected: 46,
		},
		{
			name:     "Part 2 - real input",
			input:    string(testInput),
			expected: 108956227,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := executeMain(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
