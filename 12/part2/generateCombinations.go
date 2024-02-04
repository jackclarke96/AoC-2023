package main

// func generateCombinations(springString []string, springLengths []int) int {
// 	// keysPlaced := 0
// 	questionMarkIndices := findQuestionMarkIndices(springString)
// 	iMaxStart := calculateIMax(springString, springLengths)
// 	// memo := make(map[MemoKey]int)
// 	total := 0

// 	var closure func(depth int, keysPlaced int, previousChar string) int
// 	closure = func(depth int, keysPlaced int, previousChar string) int {
// 		// Generate the key based on current state
// 		// key := StateKey{Depth: depth, NumHashes: numHashes}

// 		// Check if we have already computed this state
// 		// if val, found := memo[key]; found {
// 		// 	return val // Return the memoized value
// 		// }

// 		// Base case: if all question marks have been processed
// 		if depth == len(questionMarkIndices) {
// 			result := isValidCombination(springLengths, springString, iMaxStart)
// 			// if result == 1 {
// 			// 	memo[key] = result // Store the result in the memoization map
// 			// }
// 			// memo[key] = result // Store the result in the memoization map
// 			return result
// 		}

// 		localTotal := 0 // Local total for combinations from this state forward
// 		for _, elem := range []string{"#", "."} {
// 			springString[questionMarkIndices[depth]] = elem
// 			// updatedNumHashes := numHashes
// 			// if elem == "#" {
// 			// 	updatedNumHashes++
// 			// }
// 			if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
// 				fmt.Println(springString[:depth+1])
// 				fmt.Println(keysPlaced)
// 				localTotal += closure(depth+1, keysPlaced, elem)
// 			}
// 		}

// 		// memo[key] = localTotal // Memoize the computed total for this state
// 		return localTotal
// 	}

// 	// Initial call to the closure with starting values
// 	total = closure(0, 0, "") // Assuming "" as the initial previous character
// 	return total
// }

// func generateCombinationsWithoutMemoization(springString []string, springLengths []int) int {
// 	// keysPlaced := 0
// 	questionMarkIndices := findQuestionMarkIndicesMap(springString)
// 	iMaxStart := calculateIMax(springString, springLengths)
// 	// memo := make(map[MemoKey]int)
// 	total := 0

// 	var closure func(depth int, keysPlaced int) int

// 	closure = func(depth int, keysPlaced int) int {
// 		if depth == len(springString) {
// 			result := isValidCombination(springLengths, springString, iMaxStart)
// 			return result
// 		}

// 		localTotal := 0 // Local total for combinations from this state forward
// 		if questionMarkIndices[depth] {
// 			for _, elem := range []string{"#", "."} {
// 				springString[depth] = elem
// 				if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
// 					// fmt.Println(keysPlaced)

// 					localTotal += closure(depth+1, keysPlaced)
// 				}
// 			}
// 		} else {
// 			if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
// 				fmt.Println(keysPlaced)
// 				localTotal += closure(depth+1, keysPlaced)
// 			}
// 		}

// 		return localTotal
// 	}

// 	// Initial call to the closure with starting values
// 	total = closure(0, 0) // Assuming "" as the initial previous character
// 	return total
// }

// type MemoKey struct {
// 	Index   int
// 	NumKeys int
// }

// func generateCombinations(springString []string, springLengths []int) int {
// 	questionMarkIndices := findQuestionMarkIndicesMap(springString)
// 	iMaxStart := calculateIMax(springString, springLengths)
// 	memo := make(map[MemoKey]int)
// 	total := 0

// 	var closure func(depth int, keysPlaced int, previousChar string) int

// 	// The value of each node is equal to the sum of the values of its child nodes.
// 	// If we are in a state in which:
// 	//    * index is the same
// 	//    * number of keys placed is the same
// 	//    * the current char is the same at the index
// 	// then the value of the children will be the same as the value of the children from the first exploration
// 	// so we can just sum it on
// 	// and the value  of an unexplored will just be the total so far
// 	closure = func(depth int, keysPlaced int, previousChar string) int {

// 		memoKey := MemoKey{depth, keysPlaced}
// 		if val, found := memo[memoKey]; found {
// 			// the value of the node is equal to the sum of the values of the child nodes.
// 			// if the node has been explored already,
// 			fmt.Println(springString[:depth])
// 			fmt.Println("Been here before!:", memoKey, "Value:", val)
// 		}

// 		// base case: We have a valid combination
// 		if depth == len(springString) {
// 			result, _ := isValidCombination(springLengths, springString, iMaxStart)
// 			return result
// 		}

// 		localTotal := 0 // Local total for combinations from this state forward
// 		if questionMarkIndices[depth] {
// 			for _, elem := range []string{"#", "."} {
// 				springString[depth] = elem
// 				if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
// 					localTotal += closure(depth+1, keysPlaced, elem)
// 				}
// 				if elem == "." {
// 					memoKey := MemoKey{depth, keysPlaced}
// 					memo[memoKey] += localTotal
// 				}
// 			}
// 		} else {
// 			if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
// 				elem := springString[depth]
// 				localTotal += closure(depth+1, keysPlaced, elem)
// 				if elem == "." {
// 					memoKey := MemoKey{depth, keysPlaced}
// 					memo[memoKey] += localTotal
// 				}
// 			}
// 		}

// 		return localTotal
// 	}

// func generateCombinations(springString []string, springLengths []int) int {
// 	questionMarkIndices := findQuestionMarkIndicesMap(springString)
// 	iMaxStart := calculateIMax(springString, springLengths)
// 	memo := make(map[MemoKey]int)

// 	var closure func(depth int, keysPlaced int, previousChar string) int

// 	// The value of each node is equal to the sum of the values of its child nodes.
// 	// If we are in a state in which:
// 	//    * index is the same
// 	//    * number of keys placed is the same
// 	//    * the current char is the same at the index
// 	// then the value of the children will be the same as the value of the children from the first exploration
// 	// so we can just sum it on

// 	closure = func(depth int, keysPlaced int, previousChar string) int {
// 		// memoKey := MemoKey{depth, keysPlaced}
// 		// if previousChar == "." {
// 		// 	if val, found := memo[memoKey]; found {
// 		// 		return val
// 		// 	}
// 		// }

// 		// base case: We have a valid combination. Result = 1
// 		if depth == len(springString) {
// 			result, _ := isValidCombination(springLengths, springString, iMaxStart)
// 			return result
// 		}

// 		// Local total for combinations from this state forward. This will hold value of child nodes
// 		nodeValue := 0
// 		if questionMarkIndices[depth] {
// 			for _, elem := range []string{"#", "."} {

// 				springString[depth] = elem
// 				// check combination is valid
// 				if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {

// 					// the value of the node is equal to the sum of the value of its children
// 					nodeValue += closure(depth+1, keysPlaced, elem)
// 					// if elem == "." {
// 					// 	memoKey := MemoKey{depth, keysPlaced}
// 					// 	memo[memoKey] = nodeValue
// 					// }
// 				}

// 			}
// 		} else {
// 			if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
// 				elem := springString[depth]
// 				nodeValue += closure(depth+1, keysPlaced, elem)
// 				// if elem == "." {
// 				// 	memoKey := MemoKey{depth, keysPlaced}
// 				// 	memo[memoKey] = nodeValue
// 				// }
// 			}
// 		}
// 		return nodeValue
// 	}

// 	// Initial call to the closure with starting values
// 	result := closure(0, 0, "") // Assuming "" as the initial previous character
// 	fmt.Println(memo)
// 	// fmt.Println(memo[MemoKey{0, 0}])
// 	return result

// }

// ITS RETURNING 0 A LOT BECAUSE THE SUM OF THE NEXT DEPTH IS 0 IF WE DONT EVALUATE ALL DEPTHS. NEED TO HAVE EACH DEPTH CONTRIBUTE

type MemoKey struct {
	Index                    int
	NumHashes                int
	CurrentConsecutiveHashes int
	CurrentChar              string
}

func generateCombinations(springString []string, springLengths []int) int {
	questionMarkIndices := findQuestionMarkIndicesMap(springString)
	iMaxStart := calculateIMax(springString, springLengths)
	memo := make(map[MemoKey]int)

	var closure func(depth int, keysPlaced int, previousChar string) int

	// The value of each node is equal to the sum of the values of its child nodes.
	// If we are in a state in which:
	//    * index is the same
	//    * number of keys placed is the same
	//    * the current char is the same at the index
	// then the value of the children will be the same as the value of the children from the first exploration
	// so we can just sum it on

	closure = func(depth int, keysPlaced int, previousChar string) int {
		totalHashes, currentHashes := countHashesAndFinalConsecutiveHashes(springString[:depth])

		memoKey := MemoKey{depth, totalHashes, currentHashes, previousChar}
		if val, found := memo[memoKey]; found {
			return val
		}

		// base case: We have a valid combination. Result = 1
		if depth == len(springString) {
			result, _ := isValidCombination(springLengths, springString, iMaxStart)
			return result
		}

		// Local total for combinations from this state forward. This will hold value of child nodes
		nodeValue := 0
		if questionMarkIndices[depth] {
			for _, elem := range []string{"#", "."} {

				springString[depth] = elem
				// check combination is valid
				if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
					nodeValue += closure(depth+1, keysPlaced, elem)
				}

			}
		} else {
			if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
				elem := springString[depth]
				nodeValue += closure(depth+1, keysPlaced, elem)
			}
		}

		memo[memoKey] = nodeValue
		return nodeValue
	}

	// Initial call to the closure with starting values
	result := closure(0, 0, "") // Assuming "" as the initial previous character
	// fmt.Println(memo[MemoKey{0, 0, 0, ""}])
	// fmt.Println(memo[MemoKey{0, 0}])
	return result

}

func countHashesAndFinalConsecutiveHashes(stringSlice []string) (int, int) {
	totalHashes := 0
	finalConsecutiveHashes := 0
	consecutiveFound := false

	for _, char := range stringSlice {
		if char == "#" {
			totalHashes++
			consecutiveFound = true
		}
	}

	for i := len(stringSlice) - 1; i >= 0; i-- {
		if stringSlice[i] == "#" && consecutiveFound {
			finalConsecutiveHashes++
		} else if stringSlice[i] != "#" {
			break
		}
	}

	return totalHashes, finalConsecutiveHashes
}
