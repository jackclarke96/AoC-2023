package main

import "strconv"

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

func findQuestionMarkIndices(s []string) map[int]bool {
	indices := map[int]bool{}
	for i, elem := range s {
		if elem == "?" {
			indices[i] = true
		}
	}
	return indices
}
