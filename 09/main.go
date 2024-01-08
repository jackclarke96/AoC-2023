package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	fmt.Println(executeMain(string(input)))
}

func executeMain(input string) string {
	return ""
}

func makePrediction(inputSlice []string) int {
	return 0
}
