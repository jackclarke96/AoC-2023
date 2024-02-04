package main

func isValidCombination(springLengths []int, combination []string, iMax int) (int, int) {
	i := 0
	j := 0

	// fmt.Println("combination: ", combination)

	for i <= iMax {
		if combination[i] == "#" {

			if !checkAllCharactersHash(combination[i : i+springLengths[j]]) {
				return 0, j
			}
			// If we are placing anything other than the final set of #s, check for dot following it
			if j < len(springLengths)-1 && !checkCharacterIsDot(combination[i+springLengths[j]]) {
				return 0, j
			}

			iMax = recalculateIMax(iMax, j, springLengths)
			i += springLengths[j] - 1
			j++

			if j == len(springLengths) {
				// We have reached end of our combination. Check remaining strings are all "."
				if !checkNoMoreHashes(i, combination) {
					return 0, j
				}
				return 1, j
			}

		}
		i++
	}
	return 0, j
}

/*
combination:  [. . . .] 0 keysPlaced for combination prior to evaluation
combination:  [. . . #] 0 keysPlaced for combination prior to evaluation
combination:  [. . # .] 0 keysPlaced for combination prior to evaluation
combination:  [. . # #] 0 keysPlaced for combination prior to evaluation
combination:  [. # . .] 1 keysPlaced for combination prior to evaluation
combination:  [. # . #] 1 keysPlaced for combination prior to evaluation its valid!!!
combination:  [# . . .] 1 keysPlaced for combination prior to evaluation
combination:  [# . . #] 1 keysPlaced for combination prior to evaluation its valid!!!
combination:  [# . # .] 2 keysPlaced for combination prior to evaluation its valid!!!
*/

/* When i introduce memoization based on springs placed so far:*/
