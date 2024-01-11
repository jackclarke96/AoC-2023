package main

import (
	"log"
	"strings"
)

// Parse string input into 2d slice of DirectionChanger interface representing the grid of pipes
func parseInputToStructGrid(s string) [][]DirectionChanger {

	// initialise grid
	ss := strings.Split(s, "\n")
	height := len(ss)
	width := len(strings.TrimSpace(ss[0]))
	grid := make([][]DirectionChanger, height)
	for i := range grid {
		grid[i] = make([]DirectionChanger, width)
	}

	// fill grid with Pipe Structs
	for i := range grid {
		str := strings.TrimSpace(ss[i])
		for j := range grid[i] {
			pt := PipeType(str[j])
			p, err := generatePipeStruct(pt, i, j)
			if err != nil {
				log.Fatalf("Could not parse input into matrix: %v", err)
			}
			grid[i][j] = p
		}
	}
	return grid
}

// Get start co ordinates for traversal of Grid by finding S co ords
func getStartCoords(grid [][]DirectionChanger) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell != nil && cell.GetPipe().Type == START {
				return i, j
			}
		}
	}
	return -1, -1
}
