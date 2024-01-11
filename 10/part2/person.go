package main

// A person in the grid essentially just moves forward one place, marks as traversed,
// then rotates in the direction defined by the type of pipe they arrive in

type person struct {
	currentDirection direction
	i, j             int
}

// Understanding Slices: Passing 'grid' to this function passes a copy of the slice header,
// but the underlying array it points to remains the same. Modifications to elements of the grid
// (like marking pipes as traversed) will affect the original grid.
func (p person) traverseMap(grid [][]directionChanger) {
	for true {
		p.move()
		pipe := grid[p.i][p.j]
		if pipe.getPipe().Traversed {
			break
		}
		p.rotate(pipe)
		// Pointer Chaining: GetPipe() retrieves the original pointer to Pipe from the DirectionChanger interface,
		// allowing us to modify the actual Pipe in the grid.
		pipe.getPipe().markTraversed()
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

func (p *person) rotate(next directionChanger) {
	switch pipe := next.(type) {
	case *pipeNS:
		p.currentDirection = pipe.nextDirection(p.currentDirection)
	case *pipeEW:
		p.currentDirection = pipe.nextDirection(p.currentDirection)
	case *pipeNE:
		p.currentDirection = pipe.nextDirection(p.currentDirection)
	case *pipeNW:
		p.currentDirection = pipe.nextDirection(p.currentDirection)
	case *pipeSW:
		p.currentDirection = pipe.nextDirection(p.currentDirection)
	case *pipeSE:
		p.currentDirection = pipe.nextDirection(p.currentDirection)
	}
}
