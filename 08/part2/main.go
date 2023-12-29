package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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
	startLocations, endLocations := nodeMap.getStartFinishLocations()

	initialIteration := executeTraversal(inputSlice, nodeMap, startLocations, endLocations)
	secondIteration := executeTraversal(inputSlice, nodeMap, endLocations, endLocations)
	fmt.Println(initialIteration)
	fmt.Println(secondIteration)
	steps := make([]int, 6)
	for i, result := range secondIteration {
		steps[i] = result.steps
	}
	fmt.Println(LCM(steps[0], steps[1], steps...))
}

func executeTraversal(ss []string, nm NodeMap, startLocations []string, endLocations []string) []TraverseGraphResult {
	results := []TraverseGraphResult{}
	c := make(chan TraverseGraphResult)
	for i, _ := range startLocations {
		go func(ind int) {
			c <- traverseGraph(ss, nm, startLocations[ind], endLocations)
		}(i)
	}
	for range startLocations {
		result := <-c
		results = append(results, result)
	}
	return results
}

// assuming we always start with 0th index of ss
func traverseGraph(ss []string, nm NodeMap, graphLocation string, endLocations []string) TraverseGraphResult {
	startLocation := graphLocation
	steps := 0
	i, maxI := 0, len(ss)-1
	for true {

		mapping := nm[graphLocation]
		steps++
		graphLocation = getNextNode(ss[i], mapping)

		if slices.Contains(endLocations, graphLocation) {
			break
		}

		if i == maxI {
			i = 0
		} else {
			i++
		}
	}
	fmt.Println(i)
	return TraverseGraphResult{startLocation, steps, graphLocation}
}

func getNextNode(s string, node LeftRight) string {
	if s == "L" {
		return node.L
	}
	return node.R
}
