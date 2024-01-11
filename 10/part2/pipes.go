package main

import (
	"errors"
)

/*---------------------------------- Basic Type Definitions ---------------------------------- */

const (
	North Direction = "north"
	East  Direction = "east"
	South Direction = "south"
	West  Direction = "west"
)

const (
	NS     PipeType = "|"
	EW     PipeType = "-"
	NE     PipeType = "L"
	NW     PipeType = "J"
	SW     PipeType = "7"
	SE     PipeType = "F"
	GROUND PipeType = "."
	START  PipeType = "S"
)

type Direction string
type PipeType string

type DirectionChanger interface {
	NextDirection(currentDirection Direction) Direction
	GetPipe() *Pipe
}

type Pipe struct {
	i, j      int
	Type      PipeType
	Traversed bool
}

type PipeNS struct{ Pipe }
type PipeEW struct{ Pipe }
type PipeNE struct{ Pipe }
type PipeNW struct{ Pipe }
type PipeSW struct{ Pipe }
type PipeSE struct{ Pipe }
type PipeStart struct{ Pipe }

/*--------------------------------- Pipe Constructor Functions --------------------------------- */

func NewPipeNS(i, j int) *PipeNS {
	return &PipeNS{Pipe{i, j, NS, false}}
}

// Return pointer to the pipes so that underlying values can be accessed
func NewPipeEW(i, j int) *PipeEW {
	return &PipeEW{Pipe{i, j, EW, false}}
}

func NewPipeNE(i, j int) *PipeNE {
	return &PipeNE{Pipe{i, j, NE, false}}
}

func NewPipeNW(i, j int) *PipeNW {
	return &PipeNW{Pipe{i, j, NW, false}}
}

func NewPipeSW(i, j int) *PipeSW {
	return &PipeSW{Pipe{i, j, SW, false}}
}

func NewPipeSE(i, j int) *PipeSE {
	return &PipeSE{Pipe{i, j, SE, false}}
}

func NewPipeStart(i, j int) *PipeStart {
	return &PipeStart{Pipe{i, j, START, false}}
}

func (p PipeStart) NextDirection(currentDirection Direction) Direction {
	return North // doesn't really matter as this will change when we edit S to be a proper pipe
}

/*--------------------------------------- Getter for Pipes --------------------------------------- */

/*
When we store a value (or a pointer) in an interface, the interface holds a copy of that value (or a pointer).
when you access a method on an interface, you're actually calling a method on the value that the interface holds.
For example, grid[p.i][p.j] gives you a DirectionChanger interface. When we call pipe.Traversed,
we attempt to access a field on the interface itself, not on the Pipe struct that the interface might be holding.
This would result in a compilation error because Traversed is not a method or field of the DirectionChanger interface.

If DirectionChanger holds a pointer to a Pipe, we need a way to access that pointer.
This is what GetPipe() does. It returns the pointer to the Pipe struct that the interface is holding.
Without GetPipe(), you don't have a direct way to access the Pipe pointer from the interface.

This is why we cant use &pipe.Traversed or pipe.traversed
*/

func (p *Pipe) GetPipe() *Pipe {
	return p
}

/*-------------------------- NextDirection Functions For each PipeType -------------------------- */

func (p PipeNS) NextDirection(currentDirection Direction) Direction {
	if currentDirection == North {
		return North
	}
	return South
}

func (p PipeEW) NextDirection(currentDirection Direction) Direction {
	if currentDirection == East {
		return East
	}
	return West
}

func (p PipeNE) NextDirection(currentDirection Direction) Direction {
	if currentDirection == South {
		return East
	}
	return North
}

func (p PipeNW) NextDirection(currentDirection Direction) Direction {
	if currentDirection == South {
		return West
	}
	return North
}

func (p PipeSW) NextDirection(currentDirection Direction) Direction {
	if currentDirection == North {
		return West
	}
	return South
}

func (p PipeSE) NextDirection(currentDirection Direction) Direction {
	if currentDirection == North {
		return East
	}
	return South
}

// Orchestrator to invoke constructors of different Pipes
func generatePipeStruct(pt PipeType, i, j int) (DirectionChanger, error) {
	switch pt {
	case NS:
		return NewPipeNS(i, j), nil
	case EW:
		return NewPipeEW(i, j), nil
	case NE:
		return NewPipeNE(i, j), nil
	case NW:
		return NewPipeNW(i, j), nil
	case SW:
		return NewPipeSW(i, j), nil
	case SE:
		return NewPipeSE(i, j), nil
	case START:
		return NewPipeStart(i, j), nil
	case GROUND:
		return nil, nil
	default:
		return nil, errors.New("Could not generate DirectionChanger due to unexpected PipeType")
	}
}
