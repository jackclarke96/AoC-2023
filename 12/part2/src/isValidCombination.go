package main

func isValidCombination(lengths springLengths, combination springCombination) int {
	hashLengths := combination.getContiguousHashLengths()

	// Check  we have placed all groups of #s
	if len(hashLengths) != len(lengths) {
		return 0
	}

	// Check that # springs we have placed match desired lengths
	if !validateExactHashGroupLengths(hashLengths, lengths) {
		return 0
	}
	return 1
}
