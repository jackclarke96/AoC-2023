package main

func generateCombinations(combination springCombination, lengths springLengths) int {

	questionMarkIndices := combination.findQuestionMarkIndices()
	var closure func(depth int)
	combinationLength := len(combination)
	total := 0

	closure = func(depth int) {

		// We have a valid combination. Result = 1 if valid combination
		if depth == len(combination) {
			// fmt.Println(combination)
			total += isValidCombination(lengths, combination)
			return
		}

		// Local total for combinations from this state forward. This will hold value of child nodes
		if questionMarkIndices[depth] {
			for _, elem := range []string{"#", "."} {

				combination[depth] = elem
				// check combination is valid
				if isValidPartial(lengths, combination[:depth+1], combinationLength-(depth+1)) {
					closure(depth + 1)
				} else {
					// fmt.Println(combination[:depth+1])
				}

			}
		} else {
			if isValidPartial(lengths, combination[:depth+1], combinationLength-(depth+1)) {
				closure(depth + 1)
			} else {
				// fmt.Println(combination[:depth+1])
			}
		}

	}

	// Initial call to the closure with starting depth
	closure(0)
	return total

}
