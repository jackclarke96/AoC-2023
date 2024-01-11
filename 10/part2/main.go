package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	result := executeMain(string(input))
	fmt.Println("result =", result)
}

func executeMain(s string) int {

	// Parse string representation of grid into 2d slice of structs
	grid := parseInputToStructGrid(s)

	// Find co-ords of S in grid
	iStart, jStart := getStartCoords(grid)

	// Replace S with Pipe of correct directions
	grid[iStart][jStart] = convertPipeStart(iStart, jStart, grid)

	// Use replacement Pipe to calculate valid start direction
	startDirection := grid[iStart][jStart].nextDirection(West)

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
