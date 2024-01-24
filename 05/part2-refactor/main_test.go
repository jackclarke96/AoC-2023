package main

import (
	"fmt"
	"testing"
)

var s2f = `soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15`

func TestEvaluateScore(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		problem  string
		expected int
	}{
		{
			name: "Part 1 - 0 winning cards",
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
			problem:  "1",
			expected: 46,
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

func TestCombineLinearEquations(t *testing.T) {
	var nPlusOne []LinearEquation = []LinearEquation{
		{minInput: 0, maxInput: 13, transform: 21},
		{minInput: 14, maxInput: 14, transform: 28},
		{minInput: 15, maxInput: 21, transform: 20},
		{minInput: 22, maxInput: 25, transform: 64},
		{minInput: 26, maxInput: 43, transform: -26},
		{minInput: 44, maxInput: 49, transform: 12},
		{minInput: 50, maxInput: 51, transform: -31},
		{minInput: 52, maxInput: 53, transform: -9},
		{minInput: 54, maxInput: 58, transform: 27},
		{minInput: 59, maxInput: 68, transform: 31},
		{minInput: 69, maxInput: 81, transform: -1},
		{minInput: 82, maxInput: 92, transform: -37},
		{minInput: 93, maxInput: 97, transform: -30},
		{minInput: 98, maxInput: 98, transform: -36},
		{minInput: 99, maxInput: 99, transform: -81},
		{minInput: 100, maxInput: 9223372036854775807, transform: 0},
	}

	var pws []LinearEquation = []LinearEquation{
		{minInput: 0, maxInput: 68, transform: 1},
		{minInput: 69, maxInput: 69, transform: -69},
		{minInput: 70, maxInput: 9223372036854775807, transform: 0},
	}
	result := composeMapping(pws, nPlusOne)
	fmt.Println("result = ")
	fmt.Println(result)
}

// test safeSubtract
