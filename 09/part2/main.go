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
	c := make(chan int)
	for i, _ := range stringSlice {
		go func(ind int) {
			fields := strings.Fields(stringSlice[ind])
			c <- makePrediction(fields)
		}(i)
	}

	for range stringSlice {
		sum += <-c
	}
	fmt.Println(sum)
}

func makePrediction(inputSlice []string) int {
	intInput, err := stringsToIntegers(inputSlice)
	if err != nil {
		log.Fatalf("Failed to convert string slice input into integer slice")
	}
	processed := moveBackwards(moveForwards(intInput))
	return processed[0][0]
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
	currentRow := []int{0}
	currentRow = append(currentRow, firstRow...)
	my2DSlice := [][]int{currentRow}
	i := 0
	for allSameInts(my2DSlice[i][1:len(my2DSlice[i])]) == false {
		my2DSlice = append(my2DSlice, make([]int, len(my2DSlice[i])-1))
		for j := 1; j < len(my2DSlice[i])-1; j++ {
			my2DSlice[i+1][j] = my2DSlice[i][j+1] - my2DSlice[i][j]
		}
		i++
	}
	return my2DSlice
}

func moveBackwards(graph [][]int) [][]int {
	i := len(graph) - 1
	graph[i][0] = graph[i][1]

	for i > 0 {
		graph[i-1][0] = graph[i-1][1] - graph[i][0]
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
