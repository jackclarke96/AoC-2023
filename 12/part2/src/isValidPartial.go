package main

func isValidPartial(lengths springLengths, partial springCombination, numRemainingChars int) bool {

	hashLengths := partial.getContiguousHashLengths()
	finalChar := partial[len(partial)-1]

	// Check if we have placed too many groups of #s
	if len(hashLengths) > len(lengths) {
		return false
	}

	// Check that # springs we have placed so far match desired lengths
	if !validateHashGroupLengthsWithFinalCharCheck(hashLengths, lengths, finalChar) {
		return false
	}

	// Check that enough space remains to fit all remaining springs. We only check this on "."
	// This is because if we are not on a "." we are in the process of adding hashes,
	// so can rely on the previous check, also at a dot, in which the remaining springs were judged to have been still able to fit.
	if finalChar == "." {
		remainingLengths := lengths[len(hashLengths):]
		return hasSufficientSpaceForSprings(remainingLengths, numRemainingChars)
	}

	return true
}
