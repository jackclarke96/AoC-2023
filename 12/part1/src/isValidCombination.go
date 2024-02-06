package main

func isValidCombination(springLengths []int, combination []string, iMax int) int {
	i := 0
	j := 0

	for i <= iMax {
		if combination[i] == "#" {

			// Check all chars for the required length of contiguous springs are hashes
			if !checkAllCharactersHash(combination[i : i+springLengths[j]]) {
				return 0
			}

			// If we are placing anything other than the final set of #s, check for dot following it.
			// This is required to separate it from the next set
			if j < len(springLengths)-1 && !checkCharacterIsDot(combination[i+springLengths[j]]) {
				return 0
			}

			// Skip i ahead to to final hash in the spring. i++ invoked beneath moves it to the then correct next index.
			iMax = recalculateIMax(iMax, j, springLengths)
			i += springLengths[j] - 1
			j++

			if j == len(springLengths) {
				// We have reached end of our combination. Check remaining strings are all "."
				if checkNoMoreHashes(i, combination) {
					return 1
				}
				return 0
			}

		}
		i++
	}
	return 0
}
