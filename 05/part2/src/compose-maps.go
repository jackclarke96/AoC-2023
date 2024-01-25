package main

import "math"

// Performs full composition from layer 0 to layer 6 to return a single piecewise function describing x7 in terms of x0.
func orchestrateComposition(pws []piecewiseFunction) piecewiseFunction {
	composedSoFar := pws[0]
	for i := 0; i < len(pws)-1; i++ {
		ithMap := pws[i+1]
		composedSoFar = composedSoFar.subInPiecewiseFunction(ithMap)
	}
	return composedSoFar
}

// subs an entire piecewise function into another piecewise function
func (fn1 piecewiseFunction) subInPiecewiseFunction(fn2 piecewiseFunction) piecewiseFunction {
	composed := piecewiseFunction{}
	for _, uncomposed := range fn2 {
		composed = append(composed, uncomposed.subInPiecewiseFunction(fn1)...)
	}
	composed.orderByLowerBound()
	return composed
}

// subs an entire piecewise function into a single linear function; in our case a single 'piece' of a piecewise function
func (lf linearFunction) subInPiecewiseFunction(piecewise piecewiseFunction) piecewiseFunction {
	pw := piecewiseFunction{}
	for _, piece := range piecewise {
		combinedTransformation := lf.transform + piece.transform
		newMin, newMax := CalculateMinMax(lf, piece)
		if newMin > newMax {
			continue
		}
		combined := linearFunction{newMin, newMax, combinedTransformation}
		pw = append(pw, combined)
	}
	return pw
}

// After composing functions, calculate new boundaries for the xn in terms of x0
func CalculateMinMax(current, previous linearFunction) (int, int) {
	safeTransformedMin := safeSubtract(current.minInput, previous.transform)
	safeTransformedMax := safeSubtract(current.maxInput, previous.transform)
	newMinFloat := math.Max(float64(safeTransformedMin), float64(previous.minInput))
	newMaxFloat := math.Min(float64(safeTransformedMax), float64(previous.maxInput))
	return int(newMinFloat), int(newMaxFloat)
}

// Required to ensure nothing weird happens caused by upper bound of each map being math.maxInt
func safeSubtract(a, b int) int {
	if b < 0 && a > math.MaxInt+b {
		// Prevent overflow
		return math.MaxInt
	} else if b > 0 && a < math.MinInt+b {
		// Prevent underflow
		return math.MinInt
	}
	return a - b
}
