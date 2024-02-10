package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Failed to execute. Please provide expansion size in args e.g. go \"run main.go types.go 999999\"")
	}
	fmt.Println("Expansion size =", os.Args[1])
	expansionSize, errStrConv := strconv.Atoi(os.Args[1])
	if errStrConv != nil {
		log.Fatalf("failed to parse expansionSize into int: %v", errStrConv)
	}

	bs, errReadFile := os.ReadFile("../input.txt")
	if errReadFile != nil {
		log.Fatalf("Could not read input file: %v", errReadFile)
	}

	fmt.Println(executeMain(bs, expansionSize))
}

func executeMain(bs []byte, expansionSize int) int {
	galaxyWidth, galaxyHeight := getGalaxyDimensions(bs)
	horizontalExpansions, verticalExpansions, galaxyStructSlice := expandGalaxy(bs, galaxyHeight, galaxyWidth, expansionSize)

	combinations := combin.Combinations(len(galaxyStructSlice), 2)

	total := 0
	for _, combination := range combinations {
		galaxyStart := galaxyStructSlice[combination[0]]
		galaxyEnd := galaxyStructSlice[combination[1]]
		total += calculateDistance(calculationInput{galaxyStart, galaxyEnd, horizontalExpansions, verticalExpansions})
	}
	return total
}

// If for example have height=10, width 10 galaxy, we can check, e.g.
// whether, 1,11,21,31,41,51,61,71,81,91 have a hash. If not, column 1 has no galaxy. our width becomes 11, so we instert at 2, 13, 24,...
// if, e,g==.g. 10,11,12,13,14,15,16,17,18,19 have no hash, we insert at 20, 21,... 29,
func expandGalaxy(bs []byte, height, width, expansionSize int) (expansions, expansions, []galaxyStruct) {
	bs = bytes.ReplaceAll(bs, []byte("\n"), []byte(""))
	reGalaxy := regexp.MustCompile(`#`)
	horizontalExpansions := getHorizontalExpansions(bs, height, width, expansionSize, reGalaxy)
	verticalExpansions := getVerticalExpansions(bs, height, width, expansionSize, reGalaxy)
	galaxyStructSlice := getAllGalaxyStructs(bs, height, width, reGalaxy)

	return horizontalExpansions, verticalExpansions, galaxyStructSlice
}

func getHorizontalExpansions(bs []byte, height, width, expansionSize int, reGalaxy *regexp.Regexp) expansions {
	horizontalExpansions := make(expansions, height)
	numberExpansions := 0
	for colStart := 0; colStart < width; colStart++ {
		colHasGalaxy := false
		for row := 0; row < height; row++ {
			if reGalaxy.Match(bs[row*width+colStart : row*width+colStart+1]) {
				colHasGalaxy = true
				break
			}
		}
		if !colHasGalaxy {
			numberExpansions += expansionSize
		}
		horizontalExpansions[colStart] = numberExpansions
	}
	return horizontalExpansions
}

func getAllGalaxyStructs(bs []byte, height, width int, reGalaxy *regexp.Regexp) []galaxyStruct {
	galaxyAllIndex := reGalaxy.FindAllIndex(bs, -1)
	galaxyStructSlice := make([]galaxyStruct, len(galaxyAllIndex))
	for index, position := range galaxyAllIndex {
		col := position[0] % width
		row := int(math.Floor(float64(position[0] / height)))
		galaxyStructSlice[index] = galaxyStruct{row, col}
	}
	return galaxyStructSlice
}

func getVerticalExpansions(bs []byte, height, width, expansionSize int, reGalaxy *regexp.Regexp) expansions {
	verticalExpansions := make(expansions, height)
	numberExpansions := 0
	for rowStart := 0; rowStart < len(bs); rowStart += width {
		if !reGalaxy.Match(bs[rowStart : rowStart+width]) {
			numberExpansions += expansionSize
		}
		verticalExpansions[int(rowStart/width)] = numberExpansions
	}
	return verticalExpansions
}

func getGalaxyDimensions(bs []byte) (int, int) {
	reNewline := regexp.MustCompile(`\n`)
	newlineAllIndex := reNewline.FindAllIndex(bs, -1)
	galaxyWidth := newlineAllIndex[0][0]
	galaxyHeight := len(newlineAllIndex) + 1

	return galaxyWidth, galaxyHeight
}

func calculateDistance(input calculationInput) int {
	iDiff := int(math.Abs(float64(input.galaxyEnd.i-input.galaxyStart.i))) + int(math.Abs(float64(input.verticalExpansions[input.galaxyEnd.i]-input.verticalExpansions[input.galaxyStart.i])))
	jDiff := int(math.Abs(float64(input.galaxyEnd.j-input.galaxyStart.j))) + int(math.Abs(float64(input.horizontalExpansions[input.galaxyEnd.j]-input.horizontalExpansions[input.galaxyStart.j])))
	return iDiff + jDiff
}
