package main

func isValidPartial(springLengths []int, partial []string, remainingSpace int) bool {

	hashLengths := getContiguousHashLengths(partial)
	finalChar := partial[len(partial)-1]

	// Check if we have placed too many groups of #s
	if len(hashLengths) > len(springLengths) {
		return false
	}

	// Check that # springs we have placed so far match desired lengths
	if !validateHashGroupLengthsWithFinalCharCheck(hashLengths, springLengths, finalChar) {
		return false
	}

	// Check that enough space remains to fit all remaining springs. We only check this on "."
	// This is because if we are not on a "." we are in the process of adding hashes,
	// so can rely on the previous check, also at a dot, in which the remaining springs were judged to have been still able to fit.
	if finalChar == "." {
		remainingLengths := springLengths[len(hashLengths):]
		return checkStillRoomForRemainingSprings(remainingLengths, remainingSpace)
	}

	return true
}
