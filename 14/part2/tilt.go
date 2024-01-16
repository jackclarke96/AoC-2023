package main

import (
	"log"
)

/* ---------------------------------------- Column/Row ops ------------------------------------- */

// KEY FUNCTION HERE. This function tilts exactly one row or column (to allow for concurrent tilting)

// When tilting north, all circular rocks move as far north as possible. All square rocks are fixed in position.
// The maximum free index to which a circle rock could roll is therefore index(lastEncounteredSquareRock)+1 or index(lastEncounteredCircleRock+1)
// Therefore, we can initialise a 'targetIndex'. This will initially be 0, most northern point.
// Each time we encounter a circle rock, move it to the target Index, and set targetIndex = targetIndex +1 (because the circle rock now occupies that index)
// Each time we encounter a square rock, set targetIndex = indexOfSquareRock+1

// This is exactly the same for tilting west, but operate on rows instead of columns

// we can do the exact reverse for tilting south. Start with targetIndex = len(column)-1.
// Each time we encounter a circle rock, move it to the targetIndex, and set targetIndex = targetIndex -1 (because the circle rock now occupies that index)
// Each time we encounter a square rock, set targetIndex = indexOfSquareRock-1

// This is exactly the same for east

func (g *grid) tilt(fixedIndex int, d direction) {
	rows := false
	size := len(*g)
	if d == E || d == W {
		size = len((*g)[0])
		rows = true
	}

	// start at beginning for north, west. Start and end for south, east
	targetIndex := getTargetIndexStart(size-1, d)

	for iterator := 0; iterator < size; iterator++ {
		// traverse forwards for north, west. Traverse backwards for south, east
		nextElementIndex := getNextElementIndex(iterator, size-1, d)
		var element string
		if rows {
			element = (*g)[fixedIndex][nextElementIndex]
		} else {
			element = (*g)[nextElementIndex][fixedIndex]
		}

		if element == circle && nextElementIndex != targetIndex { // move the rock to the target
			// swap nextElementIndex and targetIndex (since target must always be a ".")
			if rows {
				(*g)[fixedIndex][nextElementIndex], (*g)[fixedIndex][targetIndex] = (*g)[fixedIndex][targetIndex], (*g)[fixedIndex][nextElementIndex]
			} else {
				(*g)[nextElementIndex][fixedIndex], (*g)[targetIndex][fixedIndex] = (*g)[targetIndex][fixedIndex], (*g)[nextElementIndex][fixedIndex]
			}
		}

		// update target index to be higher for north, west. Update to be lower for south, east
		targetIndex = updateTargetIndex(element, nextElementIndex, targetIndex, d)
	}
}

func updateTargetIndex(element string, index, targetIndex int, d direction) int {
	if element == circle {
		return updateTargetIndexCircle(targetIndex, d)
	} else if element == square {
		return updateTargetIndexSquare(index, d)
	}
	return targetIndex
}

func getNextElementIndex(index, indexMax int, d direction) int {
	if d == N || d == W {
		return index
	}
	return indexMax - index
}

func getTargetIndexStart(indexMax int, d direction) int {
	if d == N || d == W {
		return 0
	}
	return indexMax
}

func updateTargetIndexCircle(currentTargetIndex int, d direction) int {
	if d == N || d == W {
		return currentTargetIndex + 1
	}
	return currentTargetIndex - 1
}

func updateTargetIndexSquare(index int, d direction) int {
	if d == N || d == W {
		return index + 1
	}
	return index - 1
}

func (g *grid) tiltColumnNorth(column int) {
	g.tilt(column, N)
}

func (g *grid) tiltColumnSouth(column int) {
	g.tilt(column, S)
}

func (g *grid) tiltRowEast(row int) {
	g.tilt(row, E)
}

func (g *grid) tiltRowWest(row int) {
	g.tilt(row, W)
}

/* ------------------------------------------- grid ops ---------------------------------------- */

func (g *grid) tiltGridEastWest(d direction) {
	numRows := len(*g)
	// A BUFFERED CHANNEL we basically fill up and empty the buffer to a capacity. In this case numRows is capacity
	// Sends to a buffered channel block only when the buffer is full.
	done := make(chan bool, numRows)
	for i := range *g {
		go func(row int) {
			if d == W {
				g.tiltRowWest(row)
			} else if d == E {
				g.tiltRowEast(row)
			} else {
				log.Panicf("function called with direction %v. Direction should be East or West", d)
			}
			done <- true
		}(i)
	}

	//  Receives block when the buffer is empty.
	for i := 0; i < numRows; i++ {
		<-done
	}
	close(done)
}

func (g *grid) tiltGridNorthSouth(d direction) {
	numCols := len((*g)[0])
	done := make(chan bool, numCols)
	for i := range (*g)[0] {
		go func(col int) {
			if d == N {
				g.tiltColumnNorth(col)
			} else if d == S {
				g.tiltColumnSouth(col)
			} else {
				log.Panicf("function called with direction %v. Direction should be North or South", d)
			}
			done <- true
		}(i)
	}

	for i := 0; i < numCols; i++ {
		<-done
	}
	close(done)
}

func (g *grid) tiltGridNorth() {
	g.tiltGridNorthSouth(N)
}

func (g *grid) tiltGridSouth() {
	g.tiltGridNorthSouth(S)
}

func (g *grid) tiltGridEast() {
	g.tiltGridEastWest(E)
}

func (g *grid) tiltGridWest() {
	g.tiltGridEastWest(W)
}
