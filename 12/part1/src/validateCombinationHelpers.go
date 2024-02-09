package main

func getContiguousHashLengths(slice []string) []int {
	var lengths []int
	count := 0 // Count consecutive '#'

	for _, element := range slice {
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

func validateHashGroupLengthsWithFinalCharCheck(hashLengths, springLengths []int, finalChar string) bool {
	if len(hashLengths) > len(springLengths) {
		return false
	}

	for i := 0; i < len(hashLengths); i++ {
		// For the last hash group, if finalChar is "#", it's okay for it to be less or equal
		if finalChar == "#" && i == len(hashLengths)-1 {
			if hashLengths[i] > springLengths[i] {
				return false
			}
		} else if hashLengths[i] != springLengths[i] {
			// All other cases must match exactly
			return false
		}
	}
	return true
}

func validateExactHashGroupLengths(hashLengths, springLengths []int) bool {
	for i := 0; i < len(hashLengths); i++ {
		if hashLengths[i] != springLengths[i] {
			return false
		}
	}
	return true
}

func checkStillRoomForRemainingSprings(remainingLengths []int, remainingSpace int) bool {
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
