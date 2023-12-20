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
	games := parseGames(input)

	return processGames(games)
}

func parseGames(input string) gamesSlice {
	games := gamesSlice{}
	cards := strings.Split(input, "\n")
	for _, card := range cards {
		games = append(games, parseGame(card))
	}
	return games
}

func parseGame(card string) game {
	numbers := strings.Split(card, ":")[1]
	numberSlice := strings.Split(numbers, "|")
	winningNumbers := populateNumberMap(numberSlice[0])
	ourNumbers := populateNumberMap(numberSlice[1])
	return game{winningNumbers, ourNumbers, 1}
}

func processGames(games gamesSlice) int {
	total := len(games)
	for index := range games {
		gameOutcome := games[index].getGameOutcome()
		cardsToAdd := games[index].numberOfCards
		maxCardIndex := min(index+gameOutcome, len(games)-1)
		for i := index + 1; i <= maxCardIndex; i++ {
			games[i].numberOfCards += cardsToAdd
			total += cardsToAdd
		}
		index++
	}
	return total
}

func populateNumberMap(s string) gameNumbers {
	numbers := make(gameNumbers)
	fields := strings.Fields(s)
	for _, numStr := range fields {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		numbers[num] = true
	}
	return numbers
}

func (g game) getGameOutcome() int {
	wins := 0
	for key, _ := range g.ourNumbers {
		if g.winningNumbers[key] {
			wins++
		}
	}
	return wins
}
