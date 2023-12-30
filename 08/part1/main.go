package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}
	executeMain(string(input))
}

func executeMain(input string) {
	inputSlice, nodeMap := extractInputAndMap(input)
	fmt.Println(traverseGraph(inputSlice, nodeMap))
}

func traverseGraph(ss []string, nm NodeMap) int {
	steps := 0
	graphLocation := "AAA"
	i, maxI := 0, len(ss)-1
	for graphLocation != "ZZZ" {
		node := nm[graphLocation]
		steps++
		switch ss[i] {
		case "L":
			graphLocation = node.L
		case "R":
			graphLocation = node.R
		}
		if i == maxI {
			i = 0
		} else {
			i++
		}
	}
	return steps
}
