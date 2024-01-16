package main

import (
	"testing"
)

var testGridString = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

var singleCycleString = `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`

var twoCyclesString = `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`

var threeCyclesString = `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`

func TestExecuteMain(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Gives correct result for test grid",
			input:    testGridString,
			expected: 64,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := executeMain(tc.input, 1000000000)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestPerformSpin(t *testing.T) {
	testGrid := generateGrid(testGridString)
	singleCycle := generateGrid(singleCycleString)
	secondCycle := generateGrid(twoCyclesString)
	thirdCycle := generateGrid(threeCyclesString)

	t.Run("Spins", func(t *testing.T) {
		testGrid.performSpinCycle()
		for j := 0; j < len(testGrid[0]); j++ {
			if !compareColumnJ(testGrid, singleCycle, j) {
				t.Error("rotated grid did not match expected grid for column", j)
			}
		}
		testGrid.performSpinCycle()
		for j := 0; j < len(testGrid[0]); j++ {
			if !compareColumnJ(testGrid, secondCycle, j) {
				t.Error("Tilted grid did not match expected grid for column", j)
			}
		}
		testGrid.performSpinCycle()
		for j := 0; j < len(testGrid[0]); j++ {
			if !compareColumnJ(testGrid, thirdCycle, j) {
				t.Error("Tilted grid did not match expected grid for column", j)
			}
		}
	})
}
