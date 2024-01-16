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

var tiltedNorthString = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

func compareColumnJ(grid1, grid2 grid, j int) bool {
	for i := range grid1 {
		if grid1[i][j] != grid2[i][j] {
			return false
		}
	}
	return true
}

func TestExecuteMain(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Gives correct result for test grid",
			input:    testGridString,
			expected: 136,
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

func TestTiltColumnNorth(t *testing.T) {
	testGrid := generateGrid(testGridString)
	tiltedGrid := generateGrid(tiltedNorthString)

	t.Run("Selected column is tilted north", func(t *testing.T) {
		testGrid.tiltColumnNorth(0)
		if !compareColumnJ(testGrid, tiltedGrid, 0) {
			t.Errorf("Column 0 was not correctly tilted north")
		}
	})

	t.Run("Other columns remain unaffected", func(t *testing.T) {
		for j := 1; j < len(testGrid[0]); j++ {
			if !compareColumnJ(testGrid, generateGrid(testGridString), j) {
				t.Errorf("Column %d was affected by tilting column 0", j)
			}
		}
	})
}

func TestPerformTilt(t *testing.T) {
	testGrid := generateGrid(testGridString)
	tilted := generateGrid(tiltedNorthString)

	t.Run("Tilts the entire grid", func(t *testing.T) {
		testGrid.performTilt()
		for j := 0; j < len(testGrid[0]); j++ {
			if !compareColumnJ(testGrid, tilted, j) {
				t.Error("Tilted grid did not match expected grid for column", j)
			}
		}
	})
}
