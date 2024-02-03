package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	fmt.Println(executeMain(string(input)))
}

func executeMain(s string) int {
	rows := strings.Split(s, "\n")

	total := 0
	numCores := runtime.NumCPU()
	numRowsPerCore := len(rows) / numCores
	extraRows := len(rows) % numCores

	totalChan := make(chan int, numCores) // Ensure the channel has enough buffer for all goroutines

	startRow := 0
	for i := 0; i < numCores; i++ {
		endRow := startRow + numRowsPerCore
		if i < extraRows { // Distribute any extra rows among the first few goroutines
			endRow++
		}

		go func(start, end int) {
			totalChan <- handleRows(rows[start:end])
		}(startRow, endRow)

		startRow = endRow
	}

	for i := 0; i < numCores; i++ {
		total += <-totalChan
	}
	return total
}

func handleRows(rows []string) int {
	total := 0
	for _, row := range rows {
		total += handleRow(row)
	}
	return total
}

func handleRow(row string) int {
	slice := strings.Split(row, " ")
	springs, springLength := slice[0], slice[1]
	springSlice := strings.Split(springs, "")
	springLengths, _ := convertSliceStringToInt(strings.Split(springLength, ","))
	return generateCombinations(springSlice, springLengths)
}

func generateCombinations(springString []string, springLengths []int) int {
	questionMarkIndices := findQuestionMarkIndices(springString)
	iMaxStart := calculateIMax(springString, springLengths)

	total := 0
	var closure func(depth int)

	closure = func(depth int) {
		if depth == len(questionMarkIndices) {
			total += isValidCombination(springLengths, springString, iMaxStart)
			return
		}
		for _, elem := range []string{".", "#"} {
			springString[questionMarkIndices[depth]] = elem
			if !isValidPartial(springLengths, springString[:depth+1], iMaxStart) {
				return
			}

			closure(depth + 1)
		}
	}
	closure(0)
	return total
}

func isValidCombination(springLengths []int, combination []string, iMax int) int {
	i := 0
	j := 0

	for i <= iMax {
		if combination[i] == "#" {

			if !checkAllCharactersHash(combination[i : i+springLengths[j]]) {
				return 0
			}
			// If we are placing anything other than the final set of #s, check for dot following it
			if j < len(springLengths)-1 && !checkCharacterIsDot(combination[i+springLengths[j]]) {
				return 0
			}

			iMax = recalculateIMax(iMax, j, springLengths)
			i += springLengths[j] - 1
			j++

			if j == len(springLengths) {
				// We have reached end of our combination. Check remaining strings are all "."
				if !checkNoMoreHashes(i, combination) {
					return 0
				}
				return 1
			}

		}
		i++
	}
	return 0
}

func recalculateIMax(currentMax, springLengthsIndex int, springLengths []int) int {
	// we will never have to deal with what comes after final placement so always use + 1
	return currentMax + springLengths[springLengthsIndex] + 1
}

func calculateIMax(springString []string, springLengths []int) int {
	return len(springString) - (sum(springLengths) + len(springLengths) - 1)
}
