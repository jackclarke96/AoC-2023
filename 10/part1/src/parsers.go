package main

import (
	"strings"
)

// Parse string input into 2d slice of DirectionChanger interface representing the grid of pipes
func parseInputToStructGrid(s string) [][]directionChanger {

	// initialise grid
	ss := strings.Split(s, "\n")
	height := len(ss)
	width := len(strings.TrimSpace(ss[0]))
	grid := make([][]directionChanger, height)
	for i := range grid {
		grid[i] = make([]directionChanger, width)
	}

	// fill grid with Pipe Structs
	for i := range grid {
		str := strings.TrimSpace(ss[i])
		for j := range grid[i] {
			pt := pipeType(str[j])
			p := generatePipeStruct(pt)
			grid[i][j] = p
		}
	}
	return grid
}

// Get start co ordinates for traversal of Grid by finding S coords
func getStartCoords(grid [][]directionChanger) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell != nil && cell.getPipe().pipeType == START {
				return i, j
			}
		}
	}
	panic("START cell not found in grid")
}
