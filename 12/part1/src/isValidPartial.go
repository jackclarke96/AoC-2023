package main

func isValidPartial(springLengths []int, partial []string, iMax int) bool {

	i := 0
	j := 0

	for i <= min(iMax, len(partial)-1) {
		if partial[i] == "#" {

			if j == len(springLengths) {
				// we have placed too many hashes
				return false
			}

			// We either want to check up to the end of the group of the jth contiguous group of springs or until the end of the partial
			iEnd := min(i+springLengths[j], len(partial)-1)

			// We either want to check up to the end of the group of the jth contiguous group of springs or until the end of the partial
			if iEnd == i+springLengths[j] {

				// check all characters until end of the supposed group of springs are hashes
				if !checkAllCharactersHash(partial[i:iEnd]) {
					return false
				}
				// If we are placing anything other than the final set of #s, check for dot following it
				if j < len(springLengths)-1 && !checkCharacterIsDot(partial[iEnd]) {
					return false
				}

				// skip i ahead to to final hash in the spring. i++ invoked beneath moves it to the then correct next index.
				iMax = recalculateIMax(iMax, j, springLengths)
				i += springLengths[j] - 1
				j++
			} else {
				// Else, the entire key does not fit before the end of the partial. Therefore remaining chars should all be #
				return checkAllCharactersHash(partial[i:iEnd])
			}
		}
		i++
	}
	return true
}