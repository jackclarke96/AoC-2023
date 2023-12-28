package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardToScoreMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var handScoreMap = map[string]int{
	"FiveOfAKind":  7,
	"FourOfAKind":  6,
	"FullHouse":    5,
	"ThreeOfAKind": 4,
	"TwoPair":      3,
	"OnePair":      2,
	"HighCard":     1,
}

type SortedHandsResult struct {
	index int
	hands []Hand
}
type CountMap map[string]int

type Hand struct {
	cards       string
	cardsScores []int
	handScore   int
	bid         int
}

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	executeMain(string(input))
}

func executeMain(inputString string) {
	handStrings := strings.Split(inputString, "\n")
	handsStructSlice := make([]Hand, len(handStrings))
	for i, handString := range handStrings {
		handsStructSlice[i] = generateHandStruct(handString)
	}
	orderedByHand := sortByHandScore(handsStructSlice)

	sortedHandsChannel := make(chan SortedHandsResult)
	orderedByScore := make([][]Hand, 7)
	for i, _ := range orderedByHand {
		go func(ind int) {
			sortedHandsChannel <- SortedHandsResult{ind, sortByCardScores(orderedByHand[ind])}
		}(i)
	}

	for range orderedByScore {
		result := <-sortedHandsChannel
		orderedByScore[result.index] = result.hands
	}
	fmt.Println(calculateTotalWinnings(orderedByHand))
}

func calculateTotalWinnings(orderedHands [][]Hand) int {
	ranking := 1
	totalWinnings := 0
	for _, orderedByScore := range orderedHands {
		for _, hand := range orderedByScore {
			totalWinnings += hand.bid * ranking
			ranking += 1
		}
	}
	return totalWinnings
}

func generateHandStruct(s string) Hand {
	parts := strings.Split(s, " ")
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Error parsing bid")
	}
	handString := parts[0]
	countMap := generateHandMap(handString)

	return Hand{
		handString,
		generateCardScores(handString),
		generateHandScore(countMap),
		bid,
	}
}

func generateHandScore(c CountMap) int {
	mapLength := len(c)
	switch mapLength {
	case 1:
		return handScoreMap["FiveOfAKind"]
	case 2:
		return c.evaluateLengthTwoMaps()
	case 3:
		return c.evaluateLengthThreeMaps()
	case 4:
		return handScoreMap["OnePair"]
	}
	return handScoreMap["HighCard"]
}

func (c CountMap) evaluateLengthTwoMaps() int {
	for _, value := range c {
		if value == 1 || value == 4 {
			return handScoreMap["FourOfAKind"]
		}
	}
	return handScoreMap["FullHouse"]
}

func (c CountMap) evaluateLengthThreeMaps() int {
	for _, value := range c {
		if value == 3 {
			return handScoreMap["ThreeOfAKind"]
		} else if value == 2 {
			return handScoreMap["TwoPair"]
		}
	}
	return handScoreMap["FullHouse"]
}

func generateHandMap(s string) CountMap {
	countMap := make(CountMap)
	for _, char := range s {
		countMap[string(char)]++
	}
	return countMap
}

func generateCardScores(s string) []int {
	cardScores := make([]int, len(s))
	for i, char := range s {
		cardScores[i] = cardToScoreMap[string(char)]
	}
	return cardScores
}

func sortByHandScore(hands []Hand) [][]Hand {
	orderedByHand := make([][]Hand, 7)
	for i := 0; i < len(hands); i++ {
		orderedByHand[hands[i].handScore-1] = append(orderedByHand[hands[i].handScore-1], hands[i])
	}
	return orderedByHand
}

func sortByCardScores(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		return compareCardScores(hands[i].cardsScores, hands[j].cardsScores)
	})
	return hands
}

func compareCardScores(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		} else if b[i] > a[i] {
			return true
		}
	}

	return true
}
