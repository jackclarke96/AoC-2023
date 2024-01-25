package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read the input file: %v", err)
	}

	fmt.Println(executeMain(string(file)))

	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
}

func executeMain(fileStr string) (int, error) {
	fileSlice := strings.Split(fileStr, "\n\n")

	inputString, mapsString := fileSlice[0], fileSlice[1:]

	// 1: Parse input string into slice of inputRange type.
	inputRanges, errInput := ParseInputRanges(inputString)
	if errInput != nil {
		log.Fatalf("could not parse input: %v", errInput)
	}

	// 2: Parse maps string into slice of piecewise functions
	piecewiseFuncs, errPiecewise := ParseMappings(mapsString)
	if errPiecewise != nil {
		log.Fatalf("could not parse input: %v", errPiecewise)
	}

	// 3: Fill in gaps of the ranges defined in the piecewise functions with xn+1 = xn.
	for i := range piecewiseFuncs {
		piecewiseFuncs[i] = piecewiseFuncs[i].setMinBoundaryToZero().setUpperBoundaryToMaxInt().fillInRemainingGaps()
	}

	// 4: Compose each and every provided map to get x7 in terms of x0 only
	composed := orchestrateComposition(piecewiseFuncs)

	// 5: Intersect the boundary conditions of the input ranges with the piecewise map boundary ranges to filter the test inputs
	testInputs := composed.getTestInputs(inputRanges)

	// 6: Pass the filtered test inputs through the composed map to get the minimum
	return composed.findMinimumOutput(testInputs)
}
