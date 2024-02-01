package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	result := executeMain(string(input))
	fmt.Println("result =", result)
	elapsed := time.Since(start)
	fmt.Printf("Total runtime %s", elapsed)
}

func executeMain(s string) int {

	// Parse string representation of grid into 2d slice of structs
	grid := parseInputToStructGrid(s)

	// Find co-ords of S in grid
	iStart, jStart := getStartCoords(grid)

	// Replace S with Pipe of correct directions
	grid[iStart][jStart] = convertPipeStart(iStart, jStart, grid)

	// Use replacement Pipe to calculate valid start direction.
	// Doesn't matter what input direction is here as a valid output direction will always be calculated
	startDirection := grid[iStart][jStart].nextDirection(west)

	// Mark start point as traversed
	grid[iStart][jStart].getPipe().markTraversed()

	// Initialise person at start point
	person := person{startDirection, iStart, jStart}

	// Traverse all the way back to the starting point to find all boundaries
	person.traverseMap(grid)

	// Mark untraversed parts of the grid as Nil values to allow simpler checks in Ray Casting
	markUntraversedAsNil(grid)

	// Ray cast to find number of closed spaces
	return rayCast(grid)
}
