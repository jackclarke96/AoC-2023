package main

func checkRemainingSpringsInBounds(i, iMax, numberSpringsRemaining, totalSpringsLength int) bool {
	minimumLengthWithGaps := totalSpringsLength + numberSpringsRemaining - 1
	if i+minimumLengthWithGaps < iMax {
		return true
	}
	return false
}

func checkProposedSpringInBounds(j, jMax, i, iMax, springLength int) bool {
	if j < jMax {
		// check that both the hashes and the dot following it will fit. No -1 needed. First # already accounted for but also have . at end so evens out
		return i+springLength <= iMax
	}
	// check that just the hashes will fit. -1 because we are already at first hash
	return i+springLength-1 <= iMax
}
