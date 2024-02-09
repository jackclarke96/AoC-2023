package main

import "strconv"

func (combination springCombination) findQuestionMarkIndices() map[int]bool {
	indices := map[int]bool{}
	for i, elem := range combination {
		if elem == "?" {
			indices[i] = true
		}
	}
	return indices
}

func (combination springCombination) countHashes() int {
	total := 0
	for _, char := range combination {
		if char == "#" {
			total++
		}
	}

	return total
}

func (combination springCombination) countClosingConsecutiveHashes() int {
	consecutive := 0

	for i := len(combination) - 1; i >= 0; i-- {
		if combination[i] == "#" {
			consecutive++
		} else if combination[i] != "#" {
			break
		}
	}

	return consecutive
}

func convertSliceStringToInt(s []string) (springLengths, error) {
	intSlice := make(springLengths, len(s))
	for i, str := range s {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
