package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

type cardNumbers map[int]bool

type cardStruct struct {
	winningNumbers cardNumbers
	playerNumbers  cardNumbers
}

type cardSlice []cardStruct

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal("Failed to load input file with given path")
	}
	fmt.Println(executeMain(string(input)))
}

func executeMain(input string) int {
	cards := parseInputIntoCardSlice(input)
	return processCards(cards)
}

// Loop through each card and sum scores
func processCards(cards cardSlice) int {
	total := 0
	for _, card := range cards {
		total += card.getScore()
	}
	return total
}

// Function to get score for a given card. This will automatically handle 0 matches as 2^-1 = 0.5 will become 0 when coverted to int
func (c cardStruct) getScore() int {
	return int(math.Pow(2, float64(c.getNumberMatches()-1)))
}

// Compare keys of playerNumbers map and winningNumbers map to get number of matches.
func (c cardStruct) getNumberMatches() int {
	matches := 0
	for key := range c.playerNumbers {
		if c.winningNumbers[key] {
			matches++
		}
	}
	return matches
}
