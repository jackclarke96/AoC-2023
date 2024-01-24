package main

import (
	"errors"
)

/*---------------------------------- Basic Type Definitions ---------------------------------- */

const (
	North direction = "north"
	East  direction = "east"
	South direction = "south"
	West  direction = "west"
)

const (
	NS     pipeType = "|"
	EW     pipeType = "-"
	NE     pipeType = "L"
	NW     pipeType = "J"
	SW     pipeType = "7"
	SE     pipeType = "F"
	GROUND pipeType = "."
	START  pipeType = "S"
)

type direction string
type pipeType string

type directionChanger interface {
	nextDirection(currentDirection direction) direction
	getPipe() *pipe
}

type pipe struct {
	Type      pipeType
	Traversed bool
}

type pipeNS struct{ pipe }
type pipeEW struct{ pipe }
type pipeNE struct{ pipe }
type pipeNW struct{ pipe }
type pipeSW struct{ pipe }
type pipeSE struct{ pipe }
type pipeStart struct{ pipe }

/*--------------------------------- Pipe Constructor Functions --------------------------------- */

// Return pointer to the pipes so that underlying values can be accessed

func newPipeNS() *pipeNS {
	return &pipeNS{pipe{NS, false}}
}

func newPipeEW() *pipeEW {
	return &pipeEW{pipe{EW, false}}
}

func newPipeNE() *pipeNE {
	return &pipeNE{pipe{NE, false}}
}

func newPipeNW() *pipeNW {
	return &pipeNW{pipe{NW, false}}
}

func newPipeSW() *pipeSW {
	return &pipeSW{pipe{SW, false}}
}

func newPipeSE() *pipeSE {
	return &pipeSE{pipe{SE, false}}
}

func newPipeStart() *pipeStart {
	return &pipeStart{pipe{START, false}}
}

/*--------------------------------------- Getter for Pipes --------------------------------------- */

/*
When we store a value (or a pointer) in an interface, the interface holds a copy of that value (or a pointer).
when you access a method on an interface, you're actually calling a method on the value that the interface holds.
For example, grid[p.i][p.j] gives you a directionChanger interface. When we call pipe.Traversed,
we attempt to access a field on the interface itself, not on the Pipe struct that the interface might be holding.
This would result in a compilation error because Traversed is not a method or field of the directionChanger interface.

If directionChanger holds a pointer to a Pipe, we need a way to access that pointer.
This is what GetPipe() does. It returns the pointer to the Pipe struct that the interface is holding.
Without GetPipe(), you don't have a direct way to access the Pipe pointer from the interface.

This is why we cant use &pipe.Traversed or pipe.traversed
*/

func (p *pipe) getPipe() *pipe {
	return p
}

/*-------------------------- NextDirection Functions For each pipeType -------------------------- */

func (p pipeStart) nextDirection(currentDirection direction) direction {
	return North // doesn't really matter as this will change when we edit S to be a proper pipe
}

func (p *pipeNS) nextDirection(currentDirection direction) direction {
	if currentDirection == North { // person will continue north
		return North
	} // else they will continue south
	return South
}

func (p *pipeEW) nextDirection(currentDirection direction) direction {
	if currentDirection == East {
		return East
	}
	return West
}

func (p *pipeNE) nextDirection(currentDirection direction) direction {
	if currentDirection == South {
		return East
	}
	return North
}

func (p *pipeNW) nextDirection(currentDirection direction) direction {
	if currentDirection == South {
		return West
	}
	return North
}

func (p *pipeSW) nextDirection(currentDirection direction) direction {
	if currentDirection == North {
		return West
	}
	return South
}

func (p *pipeSE) nextDirection(currentDirection direction) direction {
	if currentDirection == North {
		return East
	}
	return South
}

// Orchestrator to invoke constructors of different Pipes
func generatePipeStruct(pt pipeType, i, j int) (directionChanger, error) {
	switch pt {
	case NS:
		return newPipeNS(), nil
	case EW:
		return newPipeEW(), nil
	case NE:
		return newPipeNE(), nil
	case NW:
		return newPipeNW(), nil
	case SW:
		return newPipeSW(), nil
	case SE:
		return newPipeSE(), nil
	case START:
		return newPipeStart(), nil
	case GROUND:
		return nil, nil
	default:
		return nil, errors.New("Could not generate directionChanger due to unexpected pipeType")
	}
}
