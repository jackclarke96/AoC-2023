package main

/*---------------------------------- Basic Type Definitions ---------------------------------- */

const (
	north direction = "north"
	east  direction = "east"
	south direction = "south"
	west  direction = "west"
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
	pipeType  pipeType
	traversed bool
}

type pipeNS struct{ pipe }
type pipeEW struct{ pipe }
type pipeNE struct{ pipe }
type pipeNW struct{ pipe }
type pipeSW struct{ pipe }
type pipeSE struct{ pipe }
type pipeStart struct{ pipe }

/*--------------------------------- Pipe Constructor Functions --------------------------------- */

// Orchestrator to invoke constructors of different Pipes
func generatePipeStruct(pt pipeType) directionChanger {
	switch pt {
	case NS:
		return &pipeNS{pipe{NS, false}}
	case EW:
		return &pipeEW{pipe{EW, false}}
	case NE:
		return &pipeNE{pipe{NE, false}}
	case NW:
		return &pipeNW{pipe{NW, false}}
	case SW:
		return &pipeSW{pipe{SW, false}}
	case SE:
		return &pipeSE{pipe{SE, false}}
	case START:
		return &pipeStart{pipe{START, false}}
	default: // must be ground
		return nil
	}
}

/*--------------------------------------- Getter for Pipes --------------------------------------- */

/*
When we store a value (or a pointer) in an interface, the interface holds a copy of that value (or a pointer).
when you access a method on an interface, you're actually calling a method on the value that the interface holds.
For example, grid[p.i][p.j] gives you a directionChanger interface. When we call grid[p.i][p.j].Traversed,
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

// These introduced to learn interfaces. Could just use a simple switch loop but below allows for polymorphism

func (p pipeStart) nextDirection(currentDirection direction) direction {
	return north // doesn't really matter as this will change when we edit S to be a proper pipe
}

func (p pipeNS) nextDirection(currentDirection direction) direction {
	if currentDirection == north {
		return north
	}
	return south
}

func (p pipeEW) nextDirection(currentDirection direction) direction {
	if currentDirection == east {
		return east
	}
	return west
}

func (p pipeNE) nextDirection(currentDirection direction) direction {
	if currentDirection == south {
		return east
	}
	return north
}

func (p pipeNW) nextDirection(currentDirection direction) direction {
	if currentDirection == south {
		return west
	}
	return north
}

func (p pipeSW) nextDirection(currentDirection direction) direction {
	if currentDirection == north {
		return west
	}
	return south
}

func (p pipeSE) nextDirection(currentDirection direction) direction {
	if currentDirection == north {
		return east
	}
	return south
}

// Pointer for Modification: Using a pointer receiver allows us to modify the Traversed field
// of the actual Pipe instance in the grid, rather than a copy.
func (p *pipe) markTraversed() {
	p.traversed = true
}
