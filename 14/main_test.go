package main

import (
	"strings"
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

func generateGrid(stringGrid string) grid {
	lines := strings.Split(stringGrid, "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		chars := strings.Split(line, "")
		grid[i] = chars
	}
	return grid
}

func compareColumnJ(grid1, grid2 grid, j int) bool {
	for i := range grid1 {
		// fmt.Println("grid1[i][j]", grid1[i][j], "grid2[i][j]", grid2[i][j])
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

func TestOrchestrateTilt(t *testing.T) {
	testGrid := generateGrid(testGridString)

	t.Run("Gets correct score for test input", func(t *testing.T) {
		expected := 136
		result := testGrid.performTilt()
		if result != expected {
			t.Errorf("Failed Gets correct score for test input: expected %v, got %v", expected, result)
		}
	})
}
