package main

type memoKey struct {
	index                    int
	numHashes                int
	currentConsecutiveHashes int
}

// If we define the value of a node to be the number of valid combinations possible from that point onwards,
// then the value of a node is equal to the sum of the value of its children
// If we define the state to be a combination of:
//    * The index of the spring currently being evaluated
//    * number of hashes placed up to and including the current index
//    * The number of consecutive hashes leading up to the current index
// then the number of combinations from that point onwards is the same regardless of how the previous hashes are ordered.
// This means that the value of the children will be the same as the value of the children from the first exploration

func generateCombinations(combination springCombination, lengths springLengths) int {

	questionMarkIndices := combination.findQuestionMarkIndices()
	combinationLength := len(combination)
	memo := make(map[memoKey]int)
	var closure func(depth int) int

	// define closure
	closure = func(depth int) int {
		numRemainingChars := combinationLength - (depth + 1)
		totalHashes, currentHashes := combination[:depth].countHashes(), combination[:depth].countClosingConsecutiveHashes()

		// check memoKey to see if state already explored
		memoKey := memoKey{depth, totalHashes, currentHashes}
		if val, found := memo[memoKey]; found {
			return val
		}

		// base case: We have a valid combination. Result = 1 if valid combination
		if depth == len(combination) {
			return isValidCombination(lengths, combination)
		}

		// Local total for combinations from this state forward. This will hold value of child nodes
		nodeValue := 0
		if questionMarkIndices[depth] {
			for _, elem := range []string{"#", "."} {
				combination[depth] = elem
				// check combination is valid
				if isValidPartial(lengths, combination[:depth+1], numRemainingChars) {
					nodeValue += closure(depth + 1)
				}
			}
		} else {
			if isValidPartial(lengths, combination[:depth+1], numRemainingChars) {
				nodeValue += closure(depth + 1)
			}
		}

		memo[memoKey] = nodeValue
		return nodeValue
	}

	// Initial call to the closure with starting depth
	return closure(0)

}
