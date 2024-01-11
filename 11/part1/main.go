package main

import (
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	executeMain(string(input))
}

func executeMain(s string) int {
	return 5
}

func calculateDistance(input calculationInput) int {
	return 5
}
