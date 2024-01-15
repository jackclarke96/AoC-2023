package main

import (
	"strings"
)

func parseStringIntoGridsSlice(s string) []grid {
	unparsedGrids := strings.Split(strings.TrimSpace(s), "\n\n")
	gridsSlice := make([]grid, len(unparsedGrids))
	for i, upg := range unparsedGrids {
		gridsSlice[i] = parseStringIntoGrid(upg)
	}
	return gridsSlice
}

func parseStringIntoGrid(s string) grid {
	rows := strings.Split(s, "\n")
	h, w := len(rows), len(rows[0])
	grid := make(grid, h)
	for i := range grid {
		row := make([]int, w)
		for j := range row {
			row[j] = mapCharacterToInt[string(rows[i][j])]
		}
		grid[i] = row
	}
	return grid
}

var mapCharacterToInt = map[string]int{
	"#": 1,
	".": 0,
}
