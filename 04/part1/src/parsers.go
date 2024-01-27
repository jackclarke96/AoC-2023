package main

import (
	"log"
	"strconv"
	"strings"
)

func parseInputIntoCardSlice(input string) cardSlice {
	cardSlice := cardSlice{}
	cardsString := strings.Split(input, "\n")
	for _, card := range cardsString {
		cardSlice = append(cardSlice, parseIntoCardStruct(card))
	}
	return cardSlice
}

func parseIntoCardStruct(cardString string) cardStruct {
	numbers := strings.Split(cardString, ":")[1]
	numberSlice := strings.Split(numbers, "|")
	winningNumbers, playerNumbers := populateCardNumbers(numberSlice[0]), populateCardNumbers(numberSlice[1])
	return cardStruct{winningNumbers, playerNumbers}
}

// Create a cardNumbers map. Any number that appears in the provided string will be added to the set of keys.
func populateCardNumbers(s string) cardNumbers {

	numbers := make(cardNumbers)
	fields := strings.Fields(s)

	for _, numStr := range fields {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("Error converting string to int %v:", err)
		}
		numbers[num] = true
	}
	return numbers
}
