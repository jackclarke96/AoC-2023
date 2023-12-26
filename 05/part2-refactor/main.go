package main

// each map can be rewritten as f(x) where a < x <=b
// rewriting in this way we can find ranges

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

	fileStr := string(file)
	fmt.Println(executeMain(fileStr))
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
}

func ParseIntoPiecewiseFunctions(fileStr []string) FullPiecewiseMap {
	mapsSlice, errMappings := ParseMappings(fileStr)
	if errMappings != nil {
		log.Fatalf("could not parse input: %v", errMappings)
	}

	mapsSlice = FillInGaps(mapsSlice)
	return mapsSlice
}

func executeMain(fileStr string) int {
	fileSlice := strings.Split(fileStr, "\n\n")

	inputSlice, errInput := ParseInput(fileSlice[0])
	if errInput != nil {
		log.Fatalf("could not parse input: %v", errInput)
	}

	piecewiseFuncs := ParseIntoPiecewiseFunctions(fileSlice[1:])

	composed := ComposeFullMap(piecewiseFuncs)
	output := Intersect(inputSlice, composed)
	return output
}
