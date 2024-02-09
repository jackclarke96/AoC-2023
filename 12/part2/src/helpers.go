package main

import "strconv"

func (s springLengths) unfoldSpringLengths() springLengths {
	copiedSlice := make([]int, 5*len(s))
	for i := 0; i < len(copiedSlice); i++ {
		copiedSlice[i] = s[i%(len(s))]
	}
	return copiedSlice
}

func (s springCombination) unfoldSpringArrangement() []string {
	s = append(s, "?")
	copiedSlice := make([]string, 5*len(s))
	for i := 0; i < len(copiedSlice); i++ {
		copiedSlice[i] = s[i%(len(s))]
	}
	return copiedSlice[:len(copiedSlice)-1]
}

func (s springCombination) findQuestionMarkIndices() map[int]bool {
	indices := map[int]bool{}
	for i, elem := range s {
		if elem == "?" {
			indices[i] = true
		}
	}
	return indices
}

func (s springCombination) countHashes() int {
	total := 0
	for _, char := range s {
		if char == "#" {
			total++
		}
	}

	return total
}

func (s springCombination) countClosingConsecutiveHashes() int {
	consecutive := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == "#" {
			consecutive++
		} else if s[i] != "#" {
			break
		}
	}

	return consecutive
}

func convertSliceStringToInt(slice []string) (springLengths, error) {
	intSlice := make(springLengths, len(slice))
	for i, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
