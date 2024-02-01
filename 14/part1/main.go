package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	fmt.Println(executeMain(string(input)))
}

func executeMain(s string) int {
	grid := generateGrid(s)
	grid.performTilt()
	total := 0
	gridHeight := len(grid)
	for i, row := range grid {
		for _, element := range row {
			if element == circle {
				total += gridHeight - i
			}
		}
	}
	return total

}

func (g *grid) tiltColumnNorth(j int) {

	minFreeIndex := 0
	for i := 0; i < len(*g); i++ {
		element := (*g)[i][j]
		if element == square {
			minFreeIndex = i + 1
		} else if element == circle {
			if i != minFreeIndex {
				(*g)[i][j], (*g)[minFreeIndex][j] = (*g)[minFreeIndex][j], (*g)[i][j]
			}
			minFreeIndex += 1
		}
	}
}

func (g *grid) performTilt() {
	width := len((*g)[0])
	done := make(chan bool, width)

	for j := range (*g)[0] {
		go func(col int) {
			g.tiltColumnNorth(col)
			done <- true
		}(j)
	}

	for j := 0; j < width; j++ {
		<-done
	}
	close(done)
}
