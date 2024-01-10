package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	executeMain(string(input))
}

func executeMain(s string) int {
	pipeMap, iStart, jStart := parseInputToStructMatrix(s)
	pipeMap[iStart][jStart].GetPipe().markTraversed()
	startDirection := pipeMap[iStart][jStart].NextDirection(West)
	myPerson := Person{startDirection, iStart, jStart}
	iterations := 1
	traversed := false
	for traversed == false {

		myPerson.move()
		pipeFw := pipeMap[myPerson.i][myPerson.j]
		if pipeFw.GetPipe().Traversed {
			break
		}
		myPerson.rotate(pipeFw)
		pipeFw.GetPipe().markTraversed()
		iterations++

	}

	markUntraversedAsNil(pipeMap)
	return rayCast(pipeMap)
}

func markUntraversedAsNil(pipeMap [][]DirectionChanger) {
	for i := 0; i < len(pipeMap); i++ {
		for j := 0; j < len(pipeMap[0]); j++ {
			if pipeMap[i][j] != nil && !pipeMap[i][j].GetPipe().Traversed {
				pipeMap[i][j] = nil
			}
		}
	}
}

// THEORY HERE L--J doesnt count as a border cross since you never actually make it to the gap between them. The same applies to F----7
// L--7 has a north and south so you would have to actually cross over to get past. So it is a proper boundary. Same applies to F----J
// | is always a boundary
// - is never a boundary

func rayCast(pipeMap [][]DirectionChanger) int {
	inside := 0
	for i, path := range pipeMap {
		boundariesCrossed := 0
		boundary := []PipeType{}
		for _, cell := range path {
			if cell == nil {
				if boundariesCrossed%2 == 1 {
					inside += 1
				}
				boundary = []PipeType{}
			} else {
				if cell.GetPipe().Type == NS {
					boundariesCrossed++
				} else {
					boundary = append(boundary, cell.GetPipe().Type)
					if i == 6 {
					}
					if (boundary[0] == NE && boundary[len(boundary)-1] == SW) || (boundary[0] == SE && boundary[len(boundary)-1] == NW) {
						boundariesCrossed++
						boundary = []PipeType{}
					} else if (boundary[0] == NE && boundary[len(boundary)-1] == NW) || (boundary[0] == SE && boundary[len(boundary)-1] == SW) {
						boundary = []PipeType{}
					}
				}
			}

		}
	}
	fmt.Println(inside)
	return inside
}

func (p *Pipe) markTraversed() {
	p.Traversed = true
}

func (p *Person) move() {
	switch p.currentDirection {
	case North:
		p.i -= 1
	case East:
		p.j += 1
	case South:
		p.i += 1
	case West:
		p.j -= 1
	}
}

func (p *Person) rotate(next DirectionChanger) {
	switch pipe := next.(type) {
	case *PipeNS:
		p.currentDirection = pipe.NextDirection(p.currentDirection)
	case *PipeEW:
		p.currentDirection = pipe.NextDirection(p.currentDirection)
	case *PipeNE:
		p.currentDirection = pipe.NextDirection(p.currentDirection)
	case *PipeNW:
		p.currentDirection = pipe.NextDirection(p.currentDirection)
	case *PipeSW:
		p.currentDirection = pipe.NextDirection(p.currentDirection)
	case *PipeSE:
		p.currentDirection = pipe.NextDirection(p.currentDirection)
	}
}
