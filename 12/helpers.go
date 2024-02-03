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

func modifySpringLengthsToAccountForGap(springLengths []int) []int {
	for i := 0; i < len(springLengths)-1; i++ {
		springLengths[i]++
	}
	return springLengths
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

func copySpringLengthSlice(slice []int) []int {
	copiedSlice := make([]int, 5*len(slice))
	for i := 0; i < len(copiedSlice); i++ {
		copiedSlice[i] = slice[i%(len(slice))]
	}
	return copiedSlice
}

func copySpringSlice(slice []string) []string {
	slice = append(slice, "?")
	copiedSlice := make([]string, 5*len(slice))
	for i := 0; i < len(copiedSlice); i++ {
		copiedSlice[i] = slice[i%(len(slice))]
	}
	return copiedSlice[:len(copiedSlice)-1]
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
