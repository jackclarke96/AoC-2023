package main

func isValidCombination(springLengths springLengths, combination springCombination) int {

	i := 0
	j := 0

	for i < len(combination) {
		if combination[i] == "#" {
			// We  want to check up to the end of the group of the jth contiguous group of broken springs.
			// First check this wont go out of bounds. If it does, the combination is invalid.
			if !checkProposedSpringFits(springLengths[j], len(combination), i, j == len(springLengths)-1) {
				return 0
			}

			// Check all chars for the required length of contiguous springs are hashes
			if !combination[i : i+springLengths[j]].checkAllCharactersHash() {
				return 0
			}

			// If we are placing anything other than the final set of #s, check for dot following it.
			// This is required to separate it from the next set
			if j < len(springLengths)-1 && !checkCharacterIsDot(combination[i+springLengths[j]]) {
				return 0
			}

			// Skip i ahead to to final hash in the spring. i++ invoked beneath moves it to the then correct next index.
			i += springLengths[j] - 1
			j++

			if j == len(springLengths) {
				// We have reached end of our combination. Check remaining strings are all "."
				if combination.areAllRemainingDots(i) {
					return 1
				}
				return 0
			}
		}
		i++
	}
	return 0
}
