package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	result := executeMain(string(input))
	fmt.Println(result)
}

func executeMain(s string) int {
	sum := 0
	grids := parseStringIntoGridsSlice(s)
	for _, grid := range grids {
		sum += grid.getGridScore(100, 1)
	}

	return sum
}

func (g grid) getGridScore(horizontalWeight, verticalWeight int) int {
	return verticalWeight*g.getVerticalReflectionLength() + horizontalWeight*g.getHorizontalReflectionLength()
}

func (g grid) getHorizontalReflectionLength() int {
	for i := 0; i < len(g)-1; i++ {
		if compareRows(g[i], g[i+1]) {
			if g.checkForHorizontalSymmetry(i, i+1) {
				return i + 1
			}
		}
	}
	return 0
}

func (g grid) getVerticalReflectionLength() int {
	return g.transpose().getHorizontalReflectionLength()
}

func (g grid) checkForHorizontalSymmetry(iOneStart, iTwoStart int) bool {
	for iOne, iTwo := iOneStart, iTwoStart; iOne >= 0 && iTwo < len(g); iOne, iTwo = iOne-1, iTwo+1 {
		if !compareRows(g[iOne], g[iTwo]) {
			return false
		}
	}
	return true
}
