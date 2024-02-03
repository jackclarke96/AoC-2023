package main

func isValidPartial(springLengths []int, partial []string, iMax int) bool {

	i := 0
	j := 0

	for i <= min(iMax, len(partial)-1) {
		if partial[i] == "#" {
			if j == len(springLengths) {
				return false
			}
			// extract elements starting at first # and finishing
			iEnd := min(i+springLengths[j], len(partial)-1)
			if iEnd == i+springLengths[j] {
				if !checkAllCharactersHash(partial[i:iEnd]) {
					return false
				}
				// If we are placing anything other than the final set of #s, check for dot following it
				if j < len(springLengths)-1 && !checkCharacterIsDot(partial[iEnd]) {
					return false
				}
				iMax = recalculateIMax(iMax, j, springLengths)
				i += springLengths[j] - 1
				j++
			} else {
				if !checkAllCharactersHash(partial[i:iEnd]) {
					return false
				}
				return true
			}
		}
		i++
	}
	return true
}
