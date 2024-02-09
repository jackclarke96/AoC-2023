package main

func isValidCombination(springLengths []int, springString []string) int {
	hashLengths := getContiguousHashLengths(springString)

	// Check  we have placed all groups of #s
	if len(hashLengths) != len(springLengths) {
		return 0
	}

	// Check that # springs we have placed match desired lengths
	if !validateExactHashGroupLengths(hashLengths, springLengths) {
		return 0
	}
	return 1
}
