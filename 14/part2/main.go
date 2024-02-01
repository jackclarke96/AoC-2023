package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	fmt.Println(executeMain(string(input), 1000000000))
}

func executeMain(s string, numberSpins int) int {
	gridsEncountered := []grid{}
	g := generateGrid(s)
	index := 0
	for index < numberSpins {
		g.performSpinCycle()

		for gridIndex, encounteredGrid := range gridsEncountered {
			if compareGrid(g, encounteredGrid) { // cycle detected
				cycleLength := index - gridIndex
				// offset + (some larger number)*cycleLength = 1000000000.
				// Each grid at offset + i*cycleLength will match grid at 1000000000 so long as it is after the cycle has begun
				offset := numberSpins % cycleLength
				// use offset, cycle length and start and end index of cycle to get the correct grid after 1000000000 iterations
				for cycleIndex := gridIndex; cycleIndex < index; cycleIndex++ {
					if (cycleIndex+1)%cycleLength == offset {
						return gridsEncountered[cycleIndex].getScore()
					}
				}
			}
		}
		gridsEncountered = append(gridsEncountered, g.copyGrid())
		index++
	}
	return g.getScore()
}

func (g grid) getScore() int {
	total := 0
	gridHeight := len(g)
	for i, row := range g {
		for _, element := range row {
			if element == circle {
				total += gridHeight - i
			}
		}
	}
	return total
}

func compareColumnJ(grid1, grid2 grid, j int) bool {
	for i := range grid1 {
		if grid1[i][j] != grid2[i][j] {
			return false
		}
	}
	return true
}

func compareGrid(grid1, grid2 grid) bool {
	for j := 0; j < len(grid1[0]); j++ {
		if !compareColumnJ(grid1, grid2, j) {
			return false
		}
	}
	return true
}

func (g *grid) copyGrid() grid {
	gridCopy := make([][]string, len(*g))
	for i, row := range *g {
		rowCopy := make([]string, len(row))
		copy(rowCopy, row)
		gridCopy[i] = rowCopy
	}
	return gridCopy
}

func (g *grid) performSpinCycle() {
	g.tiltGridNorth()
	g.tiltGridWest()
	g.tiltGridSouth()
	g.tiltGridEast()
}
