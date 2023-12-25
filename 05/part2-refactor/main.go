package main

// each map can be rewritten as f(x) where a < x <=b
// rewriting in this way we can find ranges
import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SingleMapping []int
type FormattedMap []SingleMapping

type LinearEquation struct {
	minInput  int
	maxInput  int
	transform int
}

type InputRange struct {
	minInput int
	maxInput int
}

type PiecewiseFunction []LinearEquation
type FullPiecewiseMap []PiecewiseFunction

func main() {
	file, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read the input file: %v", err)
	}

	fileStr := string(file)
	executeMain(fileStr)
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
	fmt.Println(output)
	return output

}
