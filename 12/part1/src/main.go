package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	input, err := os.ReadFile("../input.txt")
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
