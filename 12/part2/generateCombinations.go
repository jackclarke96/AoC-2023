package main

type memoKey struct {
	index                    int
	numHashes                int
	currentConsecutiveHashes int
}

func generateCombinations(springString []string, springLengths []int) int {
	questionMarkIndices := findQuestionMarkIndicesMap(springString)
	// iMaxStart := calculateIMax(springString, springLengths)
	memo := make(map[memoKey]int)

	var closure func(depth int) int

	// The value of each node is equal to the sum of the values of its child nodes.
	// If we are in a state in which:
	//    * index is the same
	//    * number of keys placed is the same
	//    * the current char is the same at the index
	// then the value of the children will be the same as the value of the children from the first exploration
	// so we can just sum it on

	closure = func(depth int) int {
		totalHashes, currentHashes := countHashesAndFinalConsecutiveHashes(springString[:depth])

		memoKey := memoKey{depth, totalHashes, currentHashes}
		if val, found := memo[memoKey]; found {
			return val
		}

		// base case: We have a valid combination. Result = 1 if valid combination
		if depth == len(springString) {
			return isValidCombination(springLengths, springString)
		}

		// Local total for combinations from this state forward. This will hold value of child nodes
		nodeValue := 0
		if questionMarkIndices[depth] {
			for _, elem := range []string{"#", "."} {

				springString[depth] = elem
				// check combination is valid
				if isValidPartial(springLengths, springString[:depth+1]) {
					nodeValue += closure(depth + 1)
				}

			}
		} else {
			if isValidPartial(springLengths, springString[:depth+1]) {
				nodeValue += closure(depth + 1)
			}
		}

		memo[memoKey] = nodeValue
		return nodeValue
	}

	// Initial call to the closure with starting values
	result := closure(0) // Assuming "" as the initial previous character
	return result

}
