package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type formattedMap [][]int

func main() {
	file, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Bla bla bla")
	}

	fileStr := string(file)
	fmt.Println(slices.Min(executeMain(fileStr)))

}

func executeMain(fileStr string) []int {
	fileSlice := strings.Split(fileStr, "\n\n")
	inputString, mappings := fileSlice[0], fileSlice[1:]

	inputSlice, errInput := parseInput(inputString)

	fmt.Println(inputSlice)
	mapsSlice, errMappings := parseMappings(mappings)

	if errInput != nil || errMappings != nil {
		log.Fatalf("could not load input or maps")
	}

	outputSlice := []int{}

	for _, input := range inputSlice {
		for _, mapping := range mapsSlice {
			transformation := performBinarySearch(input, mapping)
			fmt.Println("transformSlice = ", transformation)
			if len(transformation) == 3 {
				fmt.Println("input transformed from", input)
				input = applyMap(input, transformation)
				fmt.Println("to", input)
			}
		}
		outputSlice = append(outputSlice, input)
	}
	fmt.Println(outputSlice)
	return outputSlice
}

func parseInput(inputString string) ([]int, error) {
	stringSlice := strings.Fields(strings.Split(inputString, ":")[1])
	return ConvertSliceStringToInt(stringSlice)
}

func ConvertSliceStringToInt(slice []string) ([]int, error) {
	intSlice := make([]int, len(slice))
	for i, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intSlice[i] = num
	}
	return intSlice, nil
}

// convert to numbers
func parseMappings(unformatted []string) ([]formattedMap, error) {
	fms := []formattedMap{}
	for _, mapping := range unformatted {
		formatted := formattedMap{}
		splitMap := strings.Split(mapping, "\n")[1:]
		for _, split := range splitMap {
			strSlice := strings.Fields(split)
			intSlice, err := ConvertSliceStringToInt(strSlice)
			if err != nil {
				fmt.Println("err!!", err)
				return nil, err
			}
			formatted = append(formatted, intSlice)
		}
		sort.Slice(formatted, func(i, j int) bool {
			return formatted[i][1] < formatted[j][1]
		})
		fms = append(fms, formatted)
	}
	return fms, nil
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

/*
if (input) falls in range:

	(col 2) to  (col 2 + ccol 3 - 1),
	then (ouput) = (input + (col 1 - col 2))
	otherwise, output = input
*/
func performBinarySearch(input int, fm formattedMap) []int {
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
	return []int{}
}

func applyMap(input int, mapping []int) int {
	return input + mapping[0] - mapping[1]
}
