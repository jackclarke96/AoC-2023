package main

import (
	"log"
	"strconv"
	"strings"
)

// orchestrator for generating handStruct making use of functions beneath it
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

// map chars to their frequency in a hand. If a J appears, allow it to act as the currently most frequently appearing card in order to maximise the handScore.
func generateHandMap(s string) CountMap {
	countMap := make(CountMap)
	for _, char := range s {
		countMap[string(char)]++
	}

	highestNumberOfAppearances := 0
	var mostFrequentChar string
	if jFrequency, ok := countMap["J"]; ok {
		delete(countMap, "J")
		for key, val := range countMap {
			if val > highestNumberOfAppearances {
				highestNumberOfAppearances = val
				mostFrequentChar = key
			}
		}
		countMap[mostFrequentChar] += jFrequency
	}

	return countMap
}

// generate a cardScore slice with each cardScore in order
func generateCardScores(s string) []int {
	cardScores := make([]int, len(s))
	for i, char := range s {
		cardScores[i] = cardToScoreMap[string(char)]
	}
	return cardScores
}

// get the handScore of the hand
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

// helper to check whether a hand with 2 distinct card types is fourOfAKind or fullHouse and return its corresponding handScore
func (c CountMap) evaluateLengthTwoMaps() int {
	for _, value := range c {
		if value == 1 || value == 4 {
			return handScoreMap["FourOfAKind"]
		} else if value == 2 || value == 3 {
			return handScoreMap["FullHouse"]
		}
	}
	log.Fatalf("Error calculating handScores for map: %v", c)
	return 0
}

// helper to check whether a hand with 3 distinct card types is ThreeOfAKind or TwoPair and return its corresponding handScore
func (c CountMap) evaluateLengthThreeMaps() int {
	for _, value := range c {
		if value == 3 {
			return handScoreMap["ThreeOfAKind"]
		} else if value == 2 {
			return handScoreMap["TwoPair"]
		}
	}
	log.Fatalf("Error calculating handScores for map: %v", c)
	return 0
}
