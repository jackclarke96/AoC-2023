package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	executeMain(string(input))
}

// orchestrator
func executeMain(inputString string) {
	sortedHandsChannel := make(chan SortedHandsResult)
	handStrings := strings.Split(inputString, "\n")

	// Parse input into desired Hand struct format
	handsStructSlice := make([]Hand, len(handStrings)) // We know the length of the slice in advance. More efficient to do this
	for i, handString := range handStrings {
		handsStructSlice[i] = generateHandStruct(handString)
	}

	// irst group by hand score.
	orderedByHand := groupByHandScore(handsStructSlice)

	// Then order each division by card scores concurrently.
	orderedByScore := make([][]Hand, 7)
	for i, _ := range orderedByHand {
		go func(ind int) {
			sortedHandsChannel <- SortedHandsResult{ind, sortByCardScores(orderedByHand[ind])}
		}(i)
	}

	// wait for responses of each goroutine.
	for range orderedByScore {
		result := <-sortedHandsChannel
		orderedByScore[result.index] = result.hands
	}
	fmt.Println(calculateTotalWinnings(orderedByHand))
}

// group by hand
func groupByHandScore(hands []Hand) [][]Hand {
	orderedByHand := make([][]Hand, 7)
	for i := 0; i < len(hands); i++ {
		orderedByHand[hands[i].handScore-1] = append(orderedByHand[hands[i].handScore-1], hands[i])
	}
	return orderedByHand
}

// iterate through hands with the same HandScore and order based on cardScores
func sortByCardScores(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		return compareCardScores(hands[i].cardsScores, hands[j].cardsScores)
	})
	return hands
}

// compare hand A and hand B and order based on cardScores
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

// loop through fully ordered hands and get total winnings
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
