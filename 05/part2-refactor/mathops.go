package main

import "math"

// a range has values within another range if any of the following true:
// (a) the min s tarts in the range
// (b) the max starts in the range
// (c) min and max both outside the range
// or conversely, if inputMin > composedFunction[j].maxInput || inputMax < composedFunction[j].minInput
// then test the max of the mins
func Intersect(input []InputRange, composedFunction PiecewiseFunction) int {
	minOutput := -1
	for i := 0; i < len(input); i++ {
		inputMin := input[i].minInput
		inputMax := input[i].maxInput
		for j := 0; j < len(composedFunction); j++ {
			if inputMin > composedFunction[j].maxInput || inputMax < composedFunction[j].minInput {
				continue
			}
			output := max(inputMin, composedFunction[j].minInput) + composedFunction[j].transform
			if minOutput == -1 || output < minOutput {
				minOutput = output
			}
		}
	}
	return minOutput
}

// performs full composition from layer 0 to layer 6
// gives back x6 in terms of x0.
func ComposeFullMap(pws FullPiecewiseMap) PiecewiseFunction {
	compositionSlice := []PiecewiseFunction{pws[0]}
	for i := 0; i < len(pws)-1; i++ {
		composed := PiecewiseFunction{}
		for _, uncomposed := range pws[i+1] {
			composed = append(composed, CombineLinearEquations(uncomposed, compositionSlice[i])...)
		}
		compositionSlice = append(compositionSlice, composed)
	}
	return compositionSlice[6]
}

// takes function for xn+2 = xn+1 + a, as well as the entire piecewise function for xn+1 = f(xn)
// and composes them to return a new piecewise function xn+2 = f(xn)
func CombineLinearEquations(le LinearEquation, previousLayer PiecewiseFunction) PiecewiseFunction {
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
	return pw
}

// after composing functions, calculate new boundaries for the xn in terms of xo
func CalculateMinMax(current, previous LinearEquation) (int, int) {
	transformedMin := current.minInput - previous.transform
	transformedMax := current.maxInput - previous.transform
	newMinFloat := math.Max(float64(transformedMin), float64(previous.minInput))
	newMaxFloat := math.Min(float64(transformedMax), float64(previous.maxInput))

	return int(newMinFloat), int(newMaxFloat)
}
