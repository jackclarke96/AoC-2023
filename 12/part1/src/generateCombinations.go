package main

func generateCombinations(springString []string, springLengths []int) int {
	questionMarkIndices := findQuestionMarkIndices(springString)
	iMaxStart := calculateIMax(springString, springLengths)

	total := 0
	var closure func(depth int)

	closure = func(depth int) {
		if depth == len(questionMarkIndices) {
			total += isValidCombination(springLengths, springString, iMaxStart)
			return
		}
		for _, elem := range []string{".", "#"} {
			springString[questionMarkIndices[depth]] = elem
			if !isValidPartial(springLengths, springString[:depth+1], iMaxStart) {
				return
			}

			closure(depth + 1)
		}
	}
	closure(0)
	return total
}
