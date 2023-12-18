package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type EngineSlice []EngineRow
type EngineRow []string

func isSymbol(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9.]`)
	return re.MatchString(s)
}

func isNumber(s string) bool {
	re := regexp.MustCompile(`[0-9]`)
	return re.MatchString(s)
}

func (es EngineSlice) isCellInBounds(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(es) && j < len(es[i])
}

func main() {
	problem, pathToFile := os.Args[1], os.Args[2]

	bs, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("failed to read file with path %v", pathToFile)
	}

	result := executeProblem(string(bs), problem)
	fmt.Println(result)
}

// makes main easier to test
func executeProblem(engine, problem string) int {

	engineSlice := createMatrix(engine)
	total := engineSlice.iterateThroughEngine(problem)
	return total
}

func createMatrix(input string) EngineSlice {
	matrix := EngineSlice{}

	row := strings.Split(input, "\n")
	for _, s := range row {
		matrix = append(matrix, strings.Split(s, ""))
	}

	return matrix
}

func (es EngineSlice) iterateThroughEngine(problem string) int {
	num := 0
	for i, row := range es {
		for j, char := range row {
			if (isSymbol(char) && problem == "1") || char == "*" {
				num += es.checkSurroundingEntries(i, j, problem)
			}
		}
	}

	return num
}

func (es EngineSlice) checkSurroundingEntries(i, j int, problem string) int {
	results, num := []int{}, 0

	for iAdd := -1; iAdd <= 1; iAdd++ {
		for jAdd := -1; jAdd <= 1; jAdd++ {

			if iAdd == 0 && jAdd == 0 {
				continue
			}

			iStart, jStart := i+iAdd, j+jAdd

			if es.isCellInBounds(iStart, jStart) && isNumber(es[iStart][jStart]) {
				if result, err := strconv.Atoi(es[iStart].getEntireNumber(jStart)); err == nil {
					if problem == "1" {
						num += result
					} else {
						results = append(results, result)
						if len(results) > 2 {
							return num
						}
					}
				}
			}

		}
	}
	if len(results) == 2 {
		return results[0] * results[1]
	}
	return num
}

func (esRow EngineRow) getNumberLeft(j int) string {
	num := ""
	for j >= 0 {
		if isNumber(esRow[j]) == false {
			break
		}
		num = esRow[j] + num
		esRow[j] = "."
		j--
	}
	return num
}

func (esRow EngineRow) getNumberRight(j int) string {
	num := ""
	for j < len(esRow) {
		if isNumber(esRow[j]) == false {
			break
		}
		num += esRow[j]
		esRow[j] = "."
		j++
	}
	return num
}

func (esRow EngineRow) getEntireNumber(j int) string {
	numbersToRight := esRow.getNumberRight(j)
	numberToLeft := esRow.getNumberLeft(j - 1)
	return numberToLeft + numbersToRight
}
