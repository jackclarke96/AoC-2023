package main

func isValidCombination(springLengths []int, combination []string) int {

	i := 0
	j := 0

	for i < len(combination) {
		if combination[i] == "#" {

			// Check wont go out of bounds by placing spring
			if !checkProposedSpringFits(springLengths[j], len(combination), i, j == len(springLengths)-1) {
				return 0
			}

			if !checkAllCharactersHash(combination[i : i+springLengths[j]]) {
				return 0
			}

			// If we are placing anything other than the final set of #s, check for dot following it
			if j < len(springLengths)-1 && !checkCharacterIsDot(combination[i+springLengths[j]]) {
				return 0
			}

			// iMax = recalculateIMax(iMax, j, springLengths)
			i += springLengths[j] - 1
			j++

			if j == len(springLengths) {
				// We have reached end of our combination. Check remaining strings are all "."
				if !checkNoMoreHashes(i, combination) {
					return 0
				}
				return 1
			}

		}
		i++
	}
	return 0
}
