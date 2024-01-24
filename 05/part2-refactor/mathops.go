package main

import (
	"fmt"
	"math"
	"sort"
)

// a range has values within another range if any of the following true:
// (a) the min s tarts in the range
// (b) the max starts in the range
// (c) min and max both outside the range
// or conversely, if inputMin > composedFunction[j].maxInput || inputMax < composedFunction[j].minInput
// then test the max of the mins
func Intersect(input []InputRange, composedFunction PiecewiseFunction) int {
	testedVals := 0
	minOutput := -1
	for i := 0; i < len(input); i++ {
		inputMin := input[i].minInput
		inputMax := input[i].maxInput // could apply binary search here
		for j := 0; j < len(composedFunction); j++ {
			if inputMin > composedFunction[j].maxInput || inputMax < composedFunction[j].minInput {
				continue
			}
			testedVals++
			output := max(inputMin, composedFunction[j].minInput) + composedFunction[j].transform
			if minOutput == -1 || output < minOutput {
				minOutput = output
			}
		}
	}
	fmt.Println("number of tested values = ", testedVals)
	return minOutput
}

// performs full composition from layer 0 to layer 6
// gives back x6 in terms of x0.
func ComposeFullMap(pws FullPiecewiseMap) PiecewiseFunction {
	compositionSlice := []PiecewiseFunction{pws[0]}
	for i := 0; i < len(pws)-1; i++ {
		compositionSlice = append(compositionSlice, composeMapping(pws[i+1], compositionSlice[i]))
	}
	fmt.Println(compositionSlice[0])
	fmt.Println("map 1")
	fmt.Println(compositionSlice[1])
	fmt.Println("map 2")

	fmt.Println(compositionSlice[2])
	fmt.Println("map 3")

	fmt.Println(compositionSlice[3])

	fmt.Println("map 4")

	fmt.Println(compositionSlice[4])

	fmt.Println("map 5")

	fmt.Println(compositionSlice[5])

	fmt.Println("map 6")

	fmt.Println(len(compositionSlice[6]))

	return compositionSlice[6]
}

func composeMapping(pws, composedSoFar PiecewiseFunction) PiecewiseFunction {
	composed := PiecewiseFunction{}
	for _, uncomposed := range pws {
		composed = append(composed, CombineLinearEquations(uncomposed, composedSoFar)...)
		sort.Slice(composed, func(i, j int) bool {
			return composed[i].minInput < composed[j].minInput
		})
	}
	return composed
}

// takes function for xn+2 = xn+1 + a, as well as the entire piecewise function for xn+1 = f(xn)
// and composes them to return a new piecewise function xn+2 = f(xn)
func CombineLinearEquations(le LinearEquation, previousLayer PiecewiseFunction) PiecewiseFunction {
	fmt.Println("combining equation: ", le)
	fmt.Println("with ", previousLayer)
	pw := PiecewiseFunction{}
	for _, equation := range previousLayer {
		combinedTransformation := le.transform + equation.transform
		newMin, newMax := CalculateMinMax(le, equation)
		if newMin > newMax {
			continue
		}
		combined := LinearEquation{newMin, newMax, combinedTransformation}
		pw = append(pw, combined)
	}
	fmt.Println("gives")
	fmt.Println(pw)
	return pw
}

// after composing functions, calculate new boundaries for the xn in terms of xo
func CalculateMinMax(current, previous LinearEquation) (int, int) {
	safeTransformedMin := safeSubtract(int64(current.minInput), int64(previous.transform))
	safeTransformedMax := safeSubtract(int64(current.maxInput), int64(previous.transform))
	newMinFloat := math.Max(float64(safeTransformedMin), float64(previous.minInput))
	newMaxFloat := math.Min(float64(safeTransformedMax), float64(previous.maxInput))

	return int(newMinFloat), int(newMaxFloat)
}

func safeSubtract(a, b int64) int64 {
	if b < 0 && a > math.MaxInt64+b {
		// Prevent overflow
		return math.MaxInt64
	} else if b > 0 && a < math.MinInt64+b {
		// Prevent underflow
		return math.MinInt64
	}
	return a - b
}
