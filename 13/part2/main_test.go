package main

import (
	"testing"
)

var testGridsString = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

var testGrid1 = grid{
	[]int{1, 0, 1, 1, 0, 0, 1, 1, 0},
	[]int{0, 0, 1, 0, 1, 1, 0, 1, 0},
	[]int{1, 1, 0, 0, 0, 0, 0, 0, 1},
	[]int{1, 1, 0, 0, 0, 0, 0, 0, 1},
	[]int{0, 0, 1, 0, 1, 1, 0, 1, 0},
	[]int{0, 0, 1, 1, 0, 0, 1, 1, 0},
	[]int{1, 0, 1, 0, 1, 1, 0, 1, 0},
}
var testGrid2 = grid{
	[]int{1, 0, 0, 0, 1, 1, 0, 0, 1},
	[]int{1, 0, 0, 0, 0, 1, 0, 0, 1},
	[]int{0, 0, 1, 1, 0, 0, 1, 1, 1},
	[]int{1, 1, 1, 1, 1, 0, 1, 1, 0},
	[]int{1, 1, 1, 1, 1, 0, 1, 1, 0},
	[]int{0, 0, 1, 1, 0, 0, 1, 1, 1},
	[]int{1, 0, 0, 0, 0, 1, 0, 0, 1},
}

func TestExecuteMain(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Transposes the baseGrid",
			input:    testGridsString,
			expected: 400,
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

func TestGetGridScore(t *testing.T) {

	testCases := []struct {
		name             string
		receiver         grid
		horizontalWeight int
		verticalWeight   int
		expected         int
	}{
		{
			name:             "Gets correct grid score for grid 1",
			receiver:         testGrid1,
			verticalWeight:   1,
			horizontalWeight: 100,
			expected:         300,
		},
		{
			name:             "Gets correct grid score for grid 2",
			receiver:         testGrid2,
			verticalWeight:   1,
			horizontalWeight: 100,
			expected:         100,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := (tc.receiver).getGridScore(tc.horizontalWeight, tc.verticalWeight)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestGetHorizontalReflectionLength(t *testing.T) {

	testCases := []struct {
		name     string
		receiver grid
		expected int
	}{
		{
			name:     "Gets correct horizontal reflection lengthfor grid 1",
			receiver: testGrid1,
			expected: 3,
		},
		{
			name:     "Gets correct horizontal reflection lengthfor grid 2",
			receiver: testGrid2,
			expected: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := (tc.receiver).getHorizontalReflectionLength()
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestGetVerticalReflectionLength(t *testing.T) {

	testCases := []struct {
		name     string
		receiver grid
		expected int
	}{
		{
			name:     "Gets correct vertical reflection lengthfor grid 1",
			receiver: testGrid1,
			expected: 0,
		},
		{
			name:     "Gets correct vertical reflection lengthfor grid 2",
			receiver: testGrid2,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := (tc.receiver).getVerticalReflectionLength()
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
