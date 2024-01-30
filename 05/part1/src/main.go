package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type singleMap [][]int
type combinedMaps []singleMap

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file at given path: %v", err)
	}
	fmt.Println(executeMain(string(file)))
}

func executeMain(fileStr string) int {
	fileSlice := strings.Split(fileStr, "\n\n")
	inputString, mappings := fileSlice[0], fileSlice[1:]

	inputSlice, errInput := parseInput(inputString)
	mapsSlice, errMappings := parseMappings(mappings)

	if errInput != nil || errMappings != nil {
		log.Fatalf("could not load input or maps")
	}

	mapOutput := mapsSlice.passSeedsThroughMaps(inputSlice)
	return slices.Min(mapOutput)
}

func (m combinedMaps) passSeedsThroughMaps(inputSlice []int) []int {
	outputSlice := make([]int, len(inputSlice))
	for i, input := range inputSlice {
		for _, mapping := range m {
			transformation := performBinarySearch(input, mapping)
			if transformation != nil {
				input = applyMap(input, transformation)
			}
		}
		outputSlice[i] = input
	}
	return outputSlice

}

/*
Let min = 0 and max = n-1.
If max < min, then stop: target is not present in array. Return -1.
Compute guess as the average of max and min, rounded down (so that it is an integer).
If array[guess] equals target, then stop. You found it! Return guess.
If the guess was too low, that is, array[guess] < target, then set min = guess + 1.
Otherwise, the guess was too high. Set max = guess - 1.
Go back to step 2.
*/

func performBinarySearch(input int, fm singleMap) []int {
	// Initialise min as first slice entry and max as final entry
	min, max := 0, len(fm)-1

	for max >= min {
		// Guess that our input falls in the range of the middle slice entry
		guess := (max + min) / 2

		if fm[guess][1] <= input && input < fm[guess][1]+fm[guess][2] {
			// We guessed correctly
			return fm[guess]

		} else if fm[guess][1] > input {
			// We guessed too high. New max is the entry before our guess
			max = guess - 1

		} else if fm[guess][1] <= input {
			// We guessed too low. New min is the entry after our guess
			min = guess + 1
		}
	}
	return nil
}

func applyMap(input int, mapping []int) int {
	return input + mapping[0] - mapping[1]
}
