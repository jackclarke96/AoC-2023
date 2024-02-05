package main

func isValidPartial(springLengths springLengths, partial springCombination) bool {

	i := 0
	j := 0

	for i < len(partial) {
		if partial[i] == "#" {

			// if we have encountered more groups of contiguous springs than allowed
			if j == len(springLengths) {
				return false
			}

			// We either want to check up to the end of the group of the jth contiguous group of springs or until the end of the partial
			iEnd := min(i+springLengths[j], len(partial)-1)

			// if the entire group of springs fits before the end of the partial
			if iEnd == i+springLengths[j] {

				// check all characters until end of the supposed group of springs are hashes
				if !partial[i:iEnd].checkAllCharactersHash() {
					return false
				}

				// If we are placing anything other than the final set of #s, check for dot following it
				if j < len(springLengths)-1 && !checkCharacterIsDot(partial[iEnd]) {
					return false
				}

				// skip i ahead to to final hash in the spring. i++ invoked beneath moves it to the then correct next index.
				i += springLengths[j] - 1
				j++

			} else {
				// Else, the entire key does not fit before the end of the partial. Therefore remaining chars should all be #
				return partial[i:iEnd].checkAllCharactersHash()
			}
		}
		i++
	}
	return true
}
