package main

import "strconv"

func checkAllCharactersHash(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "#" {
			return false
		}
	}
	return true
}

func checkCharacterIsDot(s string) bool {
	return s == "."
}

func checkNoMoreHashes(i int, s []string) bool {
	if (i + 1) < len(s) {
		for j := i + 1; j < len(s); j++ {
			if s[j] == "#" {
				return false
			}
		}
	}
	return true
}

func convertSliceStringToInt(slice []string) ([]int, error) {
	intSlice := make([]int, len(slice))
	for i, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intSlice[i] = num
	}
	return intSlice, nil
}

func findQuestionMarkIndices(slice []string) []int {
	var indices []int
	for i, elem := range slice {
		if elem == "?" {
			indices = append(indices, i)
		}
	}
	return indices
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func recalculateIMax(currentMax, springLengthsIndex int, springLengths []int) int {
	// we will never have to deal with what comes after final placement so always use + 1
	return currentMax + springLengths[springLengthsIndex] + 1
}

func calculateIMax(springString []string, springLengths []int) int {
	return len(springString) - (sum(springLengths) + len(springLengths) - 1)
}
