package main

// import "fmt"

// func generateCombinationsCorrect(springString []string, springLengths []int) int {
// 	questionMarkIndices := findQuestionMarkIndicesMap(springString)
// 	iMaxStart := calculateIMax(springString, springLengths)
// 	memo := make(map[MemoKey]int)
// 	total := 0

// 	var closure func(depth int, keysPlaced int, previousChar string) int

// 	closure = func(depth int, keysPlaced int, previousChar string) int {

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
// 					if elem == "." {
// 						memoKey := MemoKey{depth, keysPlaced}
// 						memo[memoKey] = localTotal
// 					}
// 				}

// 			}
// 		} else {
// 			// if previousChar == "." {
// 			// 	memoKey := MemoKey{depth, keysPlaced}
// 			// 	if val, ok := memo[memoKey]; ok {
// 			// 		// fmt.Println("returning val!")
// 			// 		return val
// 			// 	}
// 			// }
// 			if ok, keysPlaced := isValidPartial(springLengths, springString[:depth+1], iMaxStart); ok {
// 				elem := springString[depth]
// 				localTotal += closure(depth+1, keysPlaced, elem)
// 				if elem == "." {
// 					memoKey := MemoKey{depth, keysPlaced}
// 					memo[memoKey] = localTotal
// 					// fmt.Println(memo)
// 				}
// 			}
// 		}

// 		return localTotal
// 	}

// 	// Initial call to the closure with starting values
// 	total = closure(0, 0, "") // Assuming "" as the initial previous character
// 	// fmt.Println(memo)
// 	fmt.Println(memo[MemoKey{0, 0}])
// 	return total
// }
