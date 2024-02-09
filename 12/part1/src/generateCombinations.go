package main

import "fmt"

func generateCombinations(springString []string, springLengths []int) int {

	questionMarkIndices := findQuestionMarkIndices(springString)
	var closure func(depth int)
	total := 0
	combinationLength := len(springString)

	closure = func(depth int) {

		// base case: We have a valid combination. Result = 1 if valid combination
		if depth == len(springString) {
			fmt.Println(springString)
			total += isValidCombination(springLengths, springString)
			return
		}

		if questionMarkIndices[depth] {
			for _, elem := range []string{"#", "."} {
				springString[depth] = elem

				if isValidPartial(springLengths, springString[:depth+1], combinationLength-(depth+1)) {
					closure(depth + 1)
				} else {
					fmt.Println(springString[:depth+1])
				}

			}
		} else {
			if isValidPartial(springLengths, springString[:depth+1], combinationLength-(depth+1)) {
				closure(depth + 1)
			} else {
				fmt.Println(springString[:depth+1])
			}
		}

	}

	// Initial call to the closure with starting depth
	closure(0)
	return total
}
