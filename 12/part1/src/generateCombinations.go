package main

func generateCombinations(combination springCombination, lengths springLengths) int {

	questionMarkIndices := combination.findQuestionMarkIndices()
	combinationLength := len(combination)
	total := 0
	var closure func(depth int)

	closure = func(depth int) {
		numRemainingChars := combinationLength - (depth + 1)

		// We have a valid combination. Result = 1 if valid combination
		if depth == len(combination) {
			total += isValidCombination(lengths, combination)
			return
		}

		if questionMarkIndices[depth] {

			// Bifurcation point. Sub in both # and . for ? and test partial is valid.
			// If so, add to partial and recurse. Otherwise, prune
			for _, elem := range []string{"#", "."} {
				combination[depth] = elem
				if isValidPartial(lengths, combination[:depth+1], numRemainingChars) {
					closure(depth + 1)
				}
			}
		} else {
			// Test next partial is valid. If so, recurse. Otherwise, prune.
			if isValidPartial(lengths, combination[:depth+1], numRemainingChars) {
				closure(depth + 1)
			}
		}
	}

	// Initial call to the closure with starting depth
	closure(0)
	return total

}
