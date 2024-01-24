package main

import "strings"

// A person in the grid essentially just moves forward one place, marks as traversed,
// then rotates in the direction defined by the type of pipe they arrive in

type person struct {
	currentDirection direction
	i, j             int
}

// Understanding Slices: Passing 'grid' to this function passes a copy of the slice header,
// but the underlying array it points to remains the same. Modifications to elements of the grid
// (like marking pipes as traversed) will affect the original grid.
func (p *person) traverseMap(grid *[][]directionChanger) {
	for true {
		// move one step forward
		p.move()
		pipe := (*grid)[p.i][p.j]
		// if Pipe already marked as traversed, loop complete
		if pipe.getPipe().Traversed {
			break
		}
		// Mark the pipe as traversed
		pipe.getPipe().markTraversed()
		// rotate person to new direction
		p.rotate(pipe)
	}
}

// Pointer for Modification: Using a pointer receiver allows us to modify the Traversed field
// of the actual Pipe instance in the grid, rather than a copy.
func (p *pipe) markTraversed() {
	p.Traversed = true
}

func (p *person) move() {
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

/*
The next.(type) syntax is a type switch, a special form of switch statement that performs actions based on the concrete type of an interface.

The 'next' variable is an interface (DirectionChanger). The type switch allows you to check what the underlying type of this interface is (e.g., *pipeNS, *pipeEW, etc.).
*/

func (p *person) rotate(pipe directionChanger) {
	p.currentDirection = pipe.nextDirection(p.currentDirection)
}

type stringSlice []string

var myStringSlice = stringSlice{"a", "b", "c"}

// Create pointer to myStringSlice
var myStringSlicePointer = &myStringSlice

func (slicePointer *stringSlice) capitaliseEntries() {

	// Dereference the pointer to get the actual slice
	slice := *slicePointer

	// Iterate over the slice and update each string
	for i, str := range slice {
		slice[i] = strings.ToUpper(str)
	}
}
