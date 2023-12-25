package main

// each map can be rewritten as f(x) where a < x <=b
// rewriting in this way we can find ranges
import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SingleMapping []int
type formattedMap []SingleMapping

type LinearEquation struct {
	minInput  int
	maxInput  int
	transform int
}

type PiecewiseFunction []LinearEquation
type FullPiecewiseMap []PiecewiseFunction

func generateLinearEquation(mp SingleMapping) LinearEquation {
	return LinearEquation{
		mp[1],
		mp[1] + mp[2] - 1,
		mp[0] - mp[1],
	}
}

func main() {
	file, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read the input file: %v", err)
	}

	fileStr := string(file)
	executeMain(fileStr)
}

func executeMain(fileStr string) int {
	fileSlice := strings.Split(fileStr, "\n\n")

	inputSlice, errInput := parseInput(fileSlice[0])
	if errInput != nil {
		log.Fatalf("could not parse input: %v", errInput)
	}

	mapsSlice, errMappings := parseMappings(fileSlice[1:])
	if errMappings != nil {
		log.Fatalf("could not parse input: %v", errInput)
	}

	mapsSlice = mapsSlice.fillInGaps()
	composed := mapsSlice.compose()

	output := intersect(inputSlice, composed)
	fmt.Println(output)
	return output

}

// a range has values within another range if any of the following true:
// (a) the min starts in the range
// (b) the max starts in the range
// (c) min and max both outside the range
// or conversely, if inputMin > composedFunction[j].maxInput || inputMax < composedFunction[j].minInput
// then test the max of the mins
func intersect(input []int, composedFunction PiecewiseFunction) int {
	minOutput := -1
	for i := 0; i < len(input); i += 2 {
		inputMin := input[i]
		inputMax := inputMin + input[i+1] - 1
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
func parseMappings(unformatted []string) (FullPiecewiseMap, error) {
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
			piecewise = append(piecewise, generateLinearEquation(intSlice))
		}
		sort.Slice(piecewise, func(i, j int) bool {
			return piecewise[i].minInput < piecewise[j].minInput
		})
		pwfs = append(pwfs, piecewise)
	}
	return pwfs, nil
}

// performs full composition from layer 0 to layer 6
// gives back x6 in terms of x0.
func (pws FullPiecewiseMap) compose() PiecewiseFunction {
	compositionSlice := []PiecewiseFunction{pws[0]}
	for i := 0; i < len(pws)-1; i++ {
		composed := PiecewiseFunction{}
		for _, uncomposed := range pws[i+1] {
			composed = append(composed, combineLinearEquations(uncomposed, compositionSlice[i])...)
		}
		compositionSlice = append(compositionSlice, composed)
	}
	return compositionSlice[6]
}

// takes function for xn+2 = xn+1 + a, as well as the entire piecewise function for xn+1 = f(xn)
// and composes them to return a new piecewise function xn+2 = f(xn)
func combineLinearEquations(le LinearEquation, previousLayer PiecewiseFunction) PiecewiseFunction {
	pw := PiecewiseFunction{}
	for _, equation := range previousLayer {
		combinedTransformation := le.transform + equation.transform
		newMin, newMax := calculateMinMax(le, equation)
		if newMin > newMax {
			continue
		}
		combined := LinearEquation{newMin, newMax, combinedTransformation}
		pw = append(pw, combined)
	}
	return pw
}

// after composing functions, calculate new boundaries for the xn in terms of xo
func calculateMinMax(current, previous LinearEquation) (int, int) {
	transformedMin := current.minInput - previous.transform
	transformedMax := current.maxInput - previous.transform
	newMinFloat := math.Max(float64(transformedMin), float64(previous.minInput))
	newMaxFloat := math.Min(float64(transformedMax), float64(previous.maxInput))

	return int(newMinFloat), int(newMaxFloat)
}

// The original map definition does not include boundaries or maps in which f(x) = x
// account for this by filling in piecewise gaps with struct containing one to one mapping
func (pws FullPiecewiseMap) fillInGaps() FullPiecewiseMap {
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
