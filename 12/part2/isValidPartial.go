package main

import "fmt"

func isValidPartial(springLengths []int, partial []string, iMax int) (bool, int) {
	// fmt.Println(partial)

	// problemSlice := []string{".", "#", "#"}
	i := 0
	j := 0

	logOut := false
	// if isEqual(problemSlice, partial) {
	// 	logOut = true
	// }
	if logOut {
		fmt.Println("partial = ", partial)
	}
	for i <= len(partial)-1 {
		if logOut {
			fmt.Println("i = ", i)
		}
		if partial[i] == "#" {
			if logOut {
				fmt.Println("its a hash")
			}
			if j == len(springLengths) {
				// fmt.Println("false!")
				return false, j
			}

			iEnd := min(i+springLengths[j], len(partial)-1)

			// if i use i+springLengths[j]-1

			// if the entire key fits before the end of the partial
			if iEnd == i+springLengths[j] {
				// check all characters until end of key are hashes
				if !checkAllCharactersHash(partial[i:iEnd]) {
					return false, j
				}
				// If we are placing anything other than the final set of #s, check for dot following it
				if j < len(springLengths)-1 && !checkCharacterIsDot(partial[iEnd]) {
					return false, j
				}

				iMax = recalculateIMax(iMax, j, springLengths)
				i += springLengths[j] - 1
				j++

			} else {
				// if the key doesn't fit, check all up until the end are hashes
				if !checkAllCharactersHash(partial[i:iEnd]) {
					return false, j
				}
				return true, j
			}
		}
		if logOut {
			fmt.Println("Incrementing I")
		}
		i++
	}
	return true, j
}

func isEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}
