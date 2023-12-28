package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Result struct {
	Number int
	Err    error
}

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	executeMain(string(input))
}

func executeMain(input string) {
	inputSlice := strings.Split(input, "\n")
	Tvals, Dvals := strings.Split(inputSlice[0], ":")[1:], strings.Split(inputSlice[1], ":")[1:]
	getMarginOfError(strings.Fields(Tvals[0]), strings.Fields(Dvals[0]))
}

func convertStringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error parsing input string, %v into number", s)
	}
	return num
}

func getMarginOfError(Tvals, Dvals []string) {
	total := 0
	c := make(chan Result)

	// for i, T := range Tvals {
	// 	go func(index int, T, D string) {
	// 		c <- getRange(convertStringToInt(T), convertStringToInt(D), c)
	// 	}(i, T, Dvals[i])

	// }

	for i := 0; i < len(Tvals); i++ {
		go func(idx int) {
			c <- getRange(convertStringToInt(Tvals[idx]), convertStringToInt(Dvals[idx]), c)
		}(i)
	}
	for range Tvals {
		result := <-c
		fmt.Println(result)
		if result.Err != nil {
			fmt.Println(result.Err)
		} else if total == 0 {
			total = result.Number
		} else {
			total *= result.Number
		}
	}
	fmt.Println(total)
}
func getRange(T, D int, c chan Result) Result {
	t := 1
	for t <= T/2 {
		if performCalculation(t, T, D) {
			return Result{T - 2*t + 1, nil}
		}
		t += 1
	}
	return Result{0, errors.New("It is not possible to guarantee a win of this race")}
}

func performCalculation(t, T, D int) bool {
	return t*(T-t) > D
}
