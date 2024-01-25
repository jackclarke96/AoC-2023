package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*------------------------------------- Parsers for Input Range and Maps -------------------------------------*/

func ParseInputRanges(inputString string) ([]inputRange, error) {

	// Split on whitespace to get numbers as string slice
	stringSlice := strings.Fields(strings.Split(inputString, ":")[1])

	// initialise inputSlice. It will hold one value per pair of numbers so is half the length of stringSlice
	inputSlice := make([]inputRange, len(stringSlice)/2)

	// Parse string slice into integer slice
	ints, err := stringSliceToIntSlice(stringSlice)
	if err != nil {
		return nil, err
	}

	// Convert each pair  of inyd into an InputRange struct with a lower and upper bound
	for i := 0; i < len(inputSlice); i += 1 {
		inputSlice[i] = generateInputRange(ints[2*i : 2*i+2])
	}

	return inputSlice, nil
}

func ParseMappings(stringMaps []string) ([]piecewiseFunction, error) {
	pwes := make([]piecewiseFunction, len(stringMaps))
	for i, stringMap := range stringMaps {
		// Remove name of map from stringMap variable
		stringMappings := strings.Split(stringMap, "\n")[1:]
		piecewise, err := stringMapToPiecewiseFunction(stringMappings)
		if err != nil {
			return nil, err
		}
		pwes[i] = piecewise.orderByLowerBound()
	}
	return pwes, nil
}

/*------------------------------------------- Helpers for Parsing -------------------------------------------- */

func generateInputRange(inputInts []int) inputRange {
	return inputRange{
		inputInts[0],
		inputInts[0] + inputInts[1] - 1, // since we have inclusive end of range
	}
}

func generateLinearFunction(inputInts []int) linearFunction {
	return linearFunction{
		inputInts[1],
		inputInts[1] + inputInts[2] - 1, // since we have inclusive end of range
		inputInts[0] - inputInts[1],
	}
}

func stringSliceToIntSlice(slice []string) ([]int, error) {
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

func stringMapToPiecewiseFunction(stringMap []string) (piecewiseFunction, error) {
	piecewise := make(piecewiseFunction, len(stringMap))
	for j, singleMapping := range stringMap {
		intSlice, err := stringSliceToIntSlice(strings.Fields(singleMapping))
		if err != nil {
			return nil, err
		}
		piecewise[j] = generateLinearFunction(intSlice)
	}
	return piecewise, nil
}

func (pw piecewiseFunction) orderByLowerBound() piecewiseFunction {
	sort.Slice(pw, func(i, j int) bool {
		return (pw)[i].minInput < (pw)[j].minInput
	})
	return pw
}

func (pw piecewiseFunction) setMinBoundaryToZero() piecewiseFunction {
	if pw[0].minInput == 0 {
		return pw
	}
	insert := linearFunction{0, pw[0].minInput - 1, 0}
	return append(piecewiseFunction{insert}, pw...)
}

func (pw piecewiseFunction) setUpperBoundaryToMaxInt() piecewiseFunction {
	insert := linearFunction{pw[len(pw)-1].maxInput + 1, math.MaxInt, 0}
	return append(pw, piecewiseFunction{insert}...)
}

func (pw piecewiseFunction) fillInRemainingGaps() piecewiseFunction {
	for j := 0; j < len(pw)-1; j++ {
		if pw[j].maxInput+1 < pw[j+1].minInput {
			insert := linearFunction{pw[j].maxInput + 1, pw[j+1].minInput - 1, 0}
			pw = append(pw[:j+1], append(piecewiseFunction{insert}, pw[j+1:]...)...)
			j++ // to skip over input we just added
		}
	}
	return pw
}
