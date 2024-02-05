package main

import "strconv"

func (s springCombination) checkAllCharactersHash() bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "#" {
			return false
		}
	}
	return true
}

func (s springCombination) areAllRemainingDots(i int) bool {
	if (i + 1) < len(s) {
		for j := i + 1; j < len(s); j++ {
			if s[j] != "." {
				return false
			}
		}
	}
	return true
}

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

func checkProposedSpringFits(springLength, combinationLength, currentIndex int, finalSpring bool) bool {
	if finalSpring {
		return currentIndex+springLength < combinationLength+1
	}
	return currentIndex+springLength < combinationLength
}

func checkCharacterIsDot(s string) bool {
	return s == "."
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
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
