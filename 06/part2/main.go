package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	TString := strings.Join(strings.Fields(Tvals[0]), "")
	DString := strings.Join(strings.Fields(Dvals[0]), "")

	getMarginOfError(TString, DString)
}

func convertStringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error parsing input string, %v into number", s)
	}
	return num
}
func getMarginOfError(Tval, Dval string) {
	numberWins, err := getRangeBinary(convertStringToInt(Tval), convertStringToInt(Dval))
	if err != nil {
		log.Println(err)
	}

	fmt.Println(numberWins)
}
func getRange(T, D int) (int, error) {
	t := 1
	for t <= T/2 {
		if performCalculation(t, T, D) {
			return T - 2*t + 1, nil
		}
		t += 1
	}
	return 0, errors.New("It is not possible to guarantee a win of this race")
}

func performCalculation(t, T, D int) bool {
	return t*(T-t) > D
}

func getRangeBinary(T, D int) (int, error) {
	tmin := 0
	tmax := T / 2
	for tmin <= tmax {
		guess := (tmax + tmin) / 2
		if performCalculation(guess, T, D) {
			if !performCalculation(guess-1, T, D) {
				return T - 2*guess + 1, nil
			}
			tmax = guess - 1
		} else {
			tmin = guess + 1
		}
	}
	return 0, errors.New("It is not possible to guarantee a win of this race")
}
