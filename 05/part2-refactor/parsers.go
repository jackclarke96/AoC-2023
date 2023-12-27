package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func generateInputRange(si []int) InputRange {
	return InputRange{
		si[0],
		si[0] + si[1] - 1,
	}
}

func GenerateLinearEquation(sm SingleMapping) LinearEquation {
	return LinearEquation{
		sm[1],
		sm[1] + sm[2] - 1,
		sm[0] - sm[1],
	}
}

func ParseInput(inputString string) ([]InputRange, error) {
	inputSlice := []InputRange{}
	stringSlice := strings.Fields(strings.Split(inputString, ":")[1])
	ints, err := ConvertSliceStringToInt(stringSlice)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ints); i += 2 {
		inputSlice = append(inputSlice, generateInputRange(ints[i:i+2]))
	}
	return inputSlice, nil
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
func ParseMappings(unformatted []string) (FullPiecewiseMap, error) {
	pwfs := []PiecewiseFunction{}
	for _, mapping := range unformatted {
		piecewise := PiecewiseFunction{}
		splitMap := strings.Split(mapping, "\n")[1:]
		for _, split := range splitMap {
			strSlice := strings.Fields(split)
			intSlice, err := ConvertSliceStringToInt(strSlice)
			if err != nil {
				return nil, err
			}
			piecewise = append(piecewise, GenerateLinearEquation(intSlice))
		}
		sort.Slice(piecewise, func(i, j int) bool {
			return piecewise[i].minInput < piecewise[j].minInput
		})
		pwfs = append(pwfs, piecewise)
	}
	return pwfs, nil
}

// The original map definition does not include boundaries or maps in which f(x) = x
// account for this by filling in piecewise gaps with struct containing one to one mapping
func FillInGaps(pws FullPiecewiseMap) FullPiecewiseMap {
	for i, fns := range pws {
		if fns[0].minInput != 0 {
			insert := LinearEquation{0, fns[0].minInput - 1, 0}
			fns = append(PiecewiseFunction{insert}, fns...)
		}
		insert := LinearEquation{fns[len(fns)-1].maxInput + 1, math.MaxInt, 0}
		fns = append(fns, PiecewiseFunction{insert}...)
		j := 0
		for j < len(fns)-1 {
			if fns[j].maxInput+1 < fns[j+1].minInput {
				insert := LinearEquation{fns[j].maxInput + 1, fns[j+1].minInput - 1, 0}
				fns = append(fns[:j+1], append(PiecewiseFunction{insert}, fns[j+1:]...)...)
				j++
			}
			j++
		}
		pws[i] = fns
	}
	return pws
}
