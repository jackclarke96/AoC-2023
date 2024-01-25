package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type singleMap [][]int
type combinedMaps []singleMap

func main() {
	file, err := os.ReadFile("./files/input.txt")
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
	min := float64(0)
	max := float64(len(fm) - 1)
	for max >= min {
		guess := int(math.Floor(0.5 * (max + min)))
		if fm[guess][1] <= input && input < fm[guess][1]+fm[guess][2] {
			return fm[guess]
		} else if fm[guess][1] > input {
			max = float64(guess - 1)
		} else if fm[guess][1] <= input {
			min = float64(guess + 1)
		}
	}
	return nil
}

func applyMap(input int, mapping []int) int {
	return input + mapping[0] - mapping[1]
}
