package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	executeMain(string(input))
}

func executeMain(input string) {
	sum := 0
	stringSlice := strings.Split(input, "\n")
	for _, slice := range stringSlice {
		fields := strings.Fields(slice)
		sum += makePrediction(fields)
	}
	fmt.Println(sum)
}

func makePrediction(inputSlice []string) int {
	intInput, err := stringsToIntegers(inputSlice)
	if err != nil {
		log.Fatalf("Failed to convert string slice input into integer slice")
	}
	processed := moveBackwards(moveForwards(intInput))
	return processed[0][len(processed[0])-1]
}

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func moveForwards(firstRow []int) [][]int {
	currentRow := append([]int(nil), firstRow...)
	currentRow = append(currentRow, 0)
	my2DSlice := [][]int{currentRow}
	i := 0
	for allSameInts(my2DSlice[i][0:len(my2DSlice[i])-1]) == false {
		my2DSlice = append(my2DSlice, make([]int, len(my2DSlice[i])-1))
		for j := 0; j < len(my2DSlice[i])-2; j++ {
			my2DSlice[i+1][j] = my2DSlice[i][j+1] - my2DSlice[i][j]
		}
		i++
	}
	return my2DSlice
}

func moveBackwards(graph [][]int) [][]int {
	i := len(graph) - 1
	j := len(graph[i]) - 1
	graph[i][j] = graph[i][j-1]
	for i > 0 {
		j = len(graph[i]) - 1
		graph[i-1][j+1] = graph[i][j] + graph[i-1][j]
		i -= 1
	}
	return graph
}

func allSameInts(s []int) bool {
	for _, v := range s {
		if v != s[0] {
			return false
		}
	}
	return true
}
