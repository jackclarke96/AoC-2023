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
type gameStruct struct {
	winningNumbers gameNumbers
	playerNumbers  gameNumbers
	numberOfCards  int
}
type gamesSlice []gameStruct

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

func parseGame(card string) gameStruct {
	numbers := strings.Split(card, ":")[1]
	numberSlice := strings.Split(numbers, "|")
	winningNumbers := populateNumberMap(numberSlice[0])
	ourNumbers := populateNumberMap(numberSlice[1])
	return gameStruct{winningNumbers, ourNumbers, 1}
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
	// game := gameStruct{
	// 	map[int]bool{41: true, 48: true, 83: true, 86: true, 17: true},
	// 	map[int]bool{83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true},
	// 	1,
	// }
	// fmt.Println(game.getGameOutcome()) // 4
	// return 2
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

func (g gameStruct) getGameOutcome() int {
	wins := 0
	for key, _ := range g.playerNumbers {
		if g.winningNumbers[key] {
			wins++
		}
	}
	return wins
}
