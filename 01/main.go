package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	problem := os.Args[1]
	if !(problem == "1" || problem == "2") {
		log.Fatalf("Please provide number indicating which problem is to be solved")
	}
	total := 0

	input, err := getInput()
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	stringSlice := strings.Split(input, "\n")

	for _, str := range stringSlice {
		numbersString := findFirstLastNum(str, problem)
		if numbersString != "" {
			firstLast, err := strconv.Atoi(numbersString)

			if err != nil {
				log.Printf("Failed to combine numbers for %s: %v", str, err)
				continue
			}
			total += firstLast
		}
	}
	fmt.Println(total)
}

func getInput() (string, error) {

	content, err := os.ReadFile("./files/input.txt")
	if err != nil {
		return "", err
	}

	return string(content), nil
}
