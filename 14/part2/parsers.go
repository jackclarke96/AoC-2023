package main

import "strings"

func generateGrid(stringGrid string) grid {
	lines := strings.Split(stringGrid, "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		chars := strings.Split(line, "")
		grid[i] = chars
	}
	return grid
}
