package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	start := time.Now() // Start timing

	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	fmt.Println(executeMain(string(input)))

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Execution took %s\n", elapsed)
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
		// fmt.Println("handling row", i)
		total += handleRow(row)
	}
	return total
}

func handleRow(row string) int {
	slice := strings.Split(row, " ")
	springs, springLength := slice[0], slice[1]
	springSlice := strings.Split(springs, "")
	springSlice = copySpringSlice(springSlice)
	springLengths, _ := convertSliceStringToInt(strings.Split(springLength, ","))
	springLengths = copySpringLengthSlice(springLengths)
	return generateCombinations(springSlice, springLengths)
}
func calculateNumberOfHashes(springString []string) int {
	numHashes := 0
	for _, val := range springString {
		if val == "#" {
			numHashes++
		}
	}
	return numHashes
}

func recalculateIMax(currentMax, springLengthsIndex int, springLengths []int) int {
	// we will never have to deal with what comes after final placement so always use + 1
	return currentMax + springLengths[springLengthsIndex] + 1
}

func calculateIMax(springString []string, springLengths []int) int {
	return len(springString) - (sum(springLengths) + len(springLengths) - 1)
}

/* To add Memoization:
 * Store a table with number of springs placed and index as key.
 */
