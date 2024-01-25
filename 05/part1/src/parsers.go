package main

import (
	"sort"
	"strconv"
	"strings"
)

func parseInput(inputString string) ([]int, error) {
	stringSlice := strings.Fields(strings.Split(inputString, ":")[1])
	return ConvertSliceStringToInt(stringSlice)
}

func ConvertSliceStringToInt(slice []string) ([]int, error) {
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

// convert to numbers
func parseMappings(unformatted []string) (combinedMaps, error) {
	fms := combinedMaps{}
	for _, mapping := range unformatted {
		formatted := singleMap{}
		splitMap := strings.Split(mapping, "\n")[1:]
		for _, split := range splitMap {
			strSlice := strings.Fields(split)
			intSlice, err := ConvertSliceStringToInt(strSlice)
			if err != nil {
				return nil, err
			}
			formatted = append(formatted, intSlice)
		}

		sort.Slice(formatted, func(i, j int) bool {
			return formatted[i][1] < formatted[j][1]
		})
		fms = append(fms, formatted)
	}
	return fms, nil
}
