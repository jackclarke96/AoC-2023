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

	fmt.Println(iterations)
	return iterations / 2
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
