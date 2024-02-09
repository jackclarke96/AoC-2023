package main

func (combination springCombination) getContiguousHashLengths() []int {
	var lengths []int
	count := 0 // Count consecutive '#'

	for _, element := range combination {
		if element == "#" {
			count++
		} else if count > 0 { // we have come to end of our # chain
			lengths = append(lengths, count)
			count = 0 // Reset count for the next group
		}
	}

	// Check if the last element in the slice was part of a contiguous group of '#'
	if count > 0 {
		lengths = append(lengths, count)
	}

	return lengths
}

func validateHashGroupLengthsWithFinalCharCheck(actual, desired springLengths, finalChar string) bool {
	if len(actual) > len(desired) {
		return false
	}

	for i := 0; i < len(actual); i++ {
		// For the last hash group, if finalChar is "#", it's okay for it to be less or equal as the next iteration may
		if finalChar == "#" && i == len(actual)-1 {
			if actual[i] > desired[i] {
				return false
			}
		} else if actual[i] != desired[i] {
			// All other cases must match exactly
			return false
		}
	}
	return true
}

func validateExactHashGroupLengths(actual, desired []int) bool {
	for i := 0; i < len(actual); i++ {
		if actual[i] != desired[i] {
			return false
		}
	}
	return true
}

func hasSufficientSpaceForSprings(remainingLengths []int, remainingSpace int) bool {
	requiredSpace := sum(remainingLengths) + len(remainingLengths) - 1
	return remainingSpace >= requiredSpace
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
