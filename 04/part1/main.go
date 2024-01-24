package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type gameNumbers map[int]bool

type game struct {
	winningNumbers gameNumbers
	ourNumbers     gameNumbers
}

func main() {
	input, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatal("Failed to load input file with given path")
	}
	fmt.Println(executeMain(string(input)))
}

func executeMain(input string) int {
	overallScore := 0
	games := strings.Split(input, "\n")
	for _, game := range games {
		overallScore += processGame(game)
	}
	return overallScore
}

func processGame(gameString string) int {

	fmt.Println(gameString)

	numbers := strings.Split(strings.Split(gameString, ":")[1], "|")
	winningNumberMap, ourNumberMap := make(gameNumbers), make(gameNumbers)

	winningNumberMap.storeNumbers(numbers[0])
	ourNumberMap.storeNumbers(numbers[1])

	g := game{winningNumberMap, ourNumberMap}

	return g.getGameOutcome()
}

func (g gameNumbers) storeNumbers(s string) {
	numbers := strings.Fields(s)
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		g[num] = true
	}
}

func (g game) getGameOutcome() int {
	wins := 0
	for key := range g.ourNumbers {
		if g.winningNumbers[key] {
			wins += 1
		}
	}
	if wins == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(wins-1)))
	}
}
