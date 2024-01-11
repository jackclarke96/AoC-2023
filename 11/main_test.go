package main

import (
	"testing"
)

var mainInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

var expansionSizeMainTest1 = 1
var expansionSizeMainTest2 = 9
var expansionSizeMainTest3 = 99

var horizontalMap, verticalMap, _ = expandGalaxy([]byte(mainInput), 10, 10, expansionSizeMainTest1)

func TestExecuteMain(t *testing.T) {
	testCases := []struct {
		name          string
		galaxyInput   string
		expansionSize int
		expected      int
	}{
		{
			name:          "executeMain - expansionSize 1",
			galaxyInput:   mainInput,
			expansionSize: expansionSizeMainTest1,
			expected:      374,
		},
		{
			name:          "executeMain - expansionSize 10",
			galaxyInput:   mainInput,
			expansionSize: expansionSizeMainTest2,
			expected:      1030,
		},
		{
			name:          "executeMain - expansionSize 100",
			galaxyInput:   mainInput,
			expansionSize: expansionSizeMainTest3,
			expected:      8410,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := executeMain([]byte(tc.galaxyInput), tc.expansionSize)
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
				horizontalMap,
				verticalMap,
			},
			expected: 15,
		},
		{
			name: "calculateDistance - test example input 2 - galaxy 3 and galaxy 6",
			input: calculationInput{
				galaxyStruct{2, 0},
				galaxyStruct{6, 9},
				horizontalMap,
				verticalMap,
			},
			expected: 17,
		},
		{
			name: "calculateDistance - test example input 2 - galaxy 8 and galaxy 9",
			input: calculationInput{
				galaxyStruct{9, 0},
				galaxyStruct{9, 4},
				horizontalMap,
				verticalMap,
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
