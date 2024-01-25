package main

import (
	"fmt"
	"math"
)

// pass each input through composed map to get min output across all inputs
func (pw piecewiseFunction) findMinimumOutput(inputs []int) (int, error) {
	minOutput := math.MaxInt
	for _, input := range inputs {
		output, err := pw.performMapping(input)
		if err != nil {
			return 0, err
		}

		if output < minOutput {
			minOutput = output
		}
	}
	return minOutput, nil
}

// Use a binary search to efficiently pass inputs through the composed map
func (orderedPiecewise piecewiseFunction) performMapping(input int) (int, error) {

	min := float64(0)
	max := float64(len(orderedPiecewise) - 1)

	for max >= min {
		// guess that input falls within boundaries of the most central function of the ordereded set of piecewise functions
		guess := int(math.Floor(0.5 * (max + min)))
		if orderedPiecewise[guess].minInput <= input && input <= orderedPiecewise[guess].maxInput {
			return input + orderedPiecewise[guess].transform, nil
			// if input is smaller than lower boundary of our guess, the piecewise function immediately before the current guess becomes the upper bound
		} else if input < orderedPiecewise[guess].minInput {
			max = float64(guess - 1)
			// if input is larger than than upper boundary of our guess, the piecewise function immediately after the current guess becomes the upper lower bound
		} else if input > orderedPiecewise[guess].maxInput {
			min = float64(guess + 1)
		}
	}
	return 0, fmt.Errorf("Error performing binary search. Input %v did not fall in any ranges", input)
}

// Intersect Boundaries of piecewise functions and input ranges to get test inputs
func (pw piecewiseFunction) getTestInputs(input []inputRange) []int {
	testInputs := []int{}
	for i := 0; i < len(input); i++ {
		// get bounds of pair of inputs
		inputMin := input[i].minInput
		inputMax := input[i].maxInput
		for j := 0; j < len(pw); j++ {
			// if input min is greater than the piecewise boundary max
			// or input max is smaller than piecewise boundary min
			// then the ranges do not overlap so there is no intersection
			if inputMin > pw[j].maxInput || inputMax < pw[j].minInput {
				continue
			}

			// else, we have an intersection. Get the minimum value in the intersection by taking the max of the lower bounds
			testInputs = append(testInputs, max(inputMin, pw[j].minInput))
		}
	}
	return testInputs
}
