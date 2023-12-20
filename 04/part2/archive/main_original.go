package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gamesStringSlice []string

type gameNumbers map[int]bool

type game struct {
	winningNumbers gameNumbers
	ourNumbers     gameNumbers
	numberOfCards  int
}

type gamesSlice []game

func main() {
	input, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatal("Failed to load input file with given path")
	}
	fmt.Println(executeMain(string(input)))
}

func executeMain(input string) int {
	gamesStringSlice := strings.Split(input, "\n")
	gs := initialiseGameSlice(gamesStringSlice)
	overallScore := gs.processGames(gamesStringSlice)

	return overallScore
}

func initialiseGameSlice(gss gamesStringSlice) gamesSlice {
	gs := gamesSlice{}
	for _, gameString := range gss {
		gs.appendGameStruct(gameString)
	}
	return gs
}

// we need to use a pointer below as we are changing the slice itself by appending to it
// as opposed to changing anything in the underlying array.

func (gs *gamesSlice) appendGameStruct(gameString string) {
	numbers := strings.Split(strings.Split(gameString, ":")[1], "|")
	winningNumberMap, ourNumberMap := make(gameNumbers), make(gameNumbers)
	winningNumberMap.storeNumbers(numbers[0])
	ourNumberMap.storeNumbers(numbers[1])
	*gs = append(*gs, game{winningNumberMap, ourNumberMap, 1})
}

func (gs gamesSlice) processGames(gss gamesStringSlice) int {
	total := len(gss)
	index := 0
	for index < len(gs) {
		gameOutcome := gs[index].getGameOutcome()
		cardsToAdd := gs[index].numberOfCards
		maxCardIndex := min(index+gameOutcome, len(gs)-1)
		for i := index + 1; i <= maxCardIndex; i++ {
			gs[i].numberOfCards += cardsToAdd
			total += cardsToAdd
		}
		index++
	}
	return total
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
	for key, _ := range g.ourNumbers {
		if g.winningNumbers[key] {
			wins += 1
		}
	}
	return wins
}
