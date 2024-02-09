package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type springCombination []string
type springLengths []int

func main() {
	start := time.Now()

	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	fmt.Println(executeMain(string(input)))

	elapsed := time.Since(start)

	fmt.Printf("Execution took %s\n", elapsed)
}

func executeMain(s string) int {
	rows := strings.Split(s, "\n")

	total := 0
	numCores := runtime.NumCPU()
	numRowsPerCore := len(rows) / numCores
	extraRows := len(rows) % numCores

	// Ensure the channel has enough buffer for all goroutines
	totalChan := make(chan int, numCores)

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

// Loop through each row and gets number of combinations for each of them.
func handleRows(rows []string) int {
	total := 0
	for _, row := range rows {
		total += handleRow(row)
	}
	return total
}

// function that parses individual row containing spring combinations and lengths of contiguous springs and returns number of combinations for that row.
func handleRow(row string) int {
	slice := strings.Split(row, " ")
	springs, springLength := slice[0], slice[1]

	// Format the spring combination into slice of individual characters and unfold
	combination := springCombination(strings.Split(springs, ""))
	combination = combination.unfoldSpringArrangement()

	// Format the spring combination into slice of ints representing lengths of contiguous broken springs and unfold
	springLengths, _ := convertSliceStringToInt(strings.Split(springLength, ","))
	springLengths = springLengths.unfoldSpringLengths()

	// invoke algorithm to determine number of valid combinations
	return generateCombinations(combination, springLengths)
}
