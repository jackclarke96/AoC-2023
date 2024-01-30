package main

import (
	"fmt"
	"log"
	"os"
)

type cardNumbers map[int]bool

type cardStruct struct {
	winningNumbers cardNumbers
	playerNumbers  cardNumbers
	numberOfCopies int
}

type cardSlice []cardStruct

var cards = cardNumbers{
	41: true,
	48: true,
	83: true,
	86: true,
	17: true,
}

var cards2 = cardNumbers{
	83: true,
	86: true,
	6:  true,
	31: true,
	17: true,
	9:  true,
	48: true,
	53: true,
}

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

func processCards(cards cardSlice) int {
	// 1 copy of each card to start with
	total := len(cards)

	for index := range cards {
		// get matchesᵢ, the quantity of matching numbers on the card at the current index
		matches := cards[index].getNumberMatches()

		// get copiesᵢ, the number of copies of the card at the current index
		copiesToAdd := cards[index].numberOfCopies

		// We want to add copies to the next matchesᵢ cards unless we run out of cards to add copies to
		maxCardIndex := min(index+matches, len(cards)-1)
		for i := index + 1; i <= maxCardIndex; i++ {

			// update the numberOfCopies field on each of the cardStructs
			cards[i].numberOfCopies += copiesToAdd

			// also update the overall total number of scratch cards
			total += copiesToAdd
		}
	}
	return total
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
