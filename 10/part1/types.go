package main

import (
	"errors"
)

type Person struct {
	currentDirection Direction
	i, j             int
}

type Direction string
type PipeType string

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

type DirectionChanger interface {
	NextDirection(currentDirection Direction) Direction
	GetPipe() *Pipe
}

type Pipe struct {
	i, j                              int
	Type                              PipeType
	Traversed                         bool
	ForwardDistance, BackwardDistance int
} // could add validMoves after traversed

type PipeNS struct{ Pipe }
type PipeEW struct{ Pipe }
type PipeNE struct{ Pipe }
type PipeNW struct{ Pipe }
type PipeSW struct{ Pipe }
type PipeSE struct{ Pipe }
type PipeStart struct{ Pipe }

func (p *Pipe) GetPipe() *Pipe {
	return p
}

func NewPipeNS(i, j int) *PipeNS {
	return &PipeNS{Pipe{i, j, NS, false, 0, 0}}
}

func NewPipeEW(i, j int) *PipeEW {
	return &PipeEW{Pipe{i, j, EW, false, 0, 0}}
}

func NewPipeNE(i, j int) *PipeNE {
	return &PipeNE{Pipe{i, j, NE, false, 0, 0}}
}

func NewPipeNW(i, j int) *PipeNW {
	return &PipeNW{Pipe{i, j, NW, false, 0, 0}}
}

func NewPipeSW(i, j int) *PipeSW {
	return &PipeSW{Pipe{i, j, SW, false, 0, 0}}
}

func NewPipeSE(i, j int) *PipeSE {
	return &PipeSE{Pipe{i, j, SE, false, 0, 0}}
}

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
	case GROUND: // Ground and Starting Position
		return nil, nil
	default:
		return nil, errors.New("Could not generate DirectionChanger due to unexpected PipeType")
	}
}
