package main

import (
	"log"
	"strings"
)

func parseInputToStructMatrix(s string) ([][]DirectionChanger, int, int) {
	ss := strings.Split(s, "\n")
	height := len(ss)
	width := len(strings.TrimSpace(ss[0]))
	var iStart, jStart int
	matrix := make([][]DirectionChanger, height) // surround matrix with nil values incase the border is around the outside
	for i := range matrix {
		matrix[i] = make([]DirectionChanger, width)
	}

	for i := range matrix {
		str := strings.TrimSpace(ss[i])
		for j := range matrix[i] {
			pt := PipeType(str[j])
			if pt == START {
				iStart, jStart = i, j
				continue
			}
			p, err := generatePipeStruct(pt, i, j)
			if err != nil {
				log.Fatalf("Could not parse input into matrix: %v", err)
			}
			matrix[i][j] = p
		}
	}
	p := NewPipeStart(iStart, jStart, matrix)
	matrix[iStart][jStart] = p

	return matrix, iStart, jStart
}
