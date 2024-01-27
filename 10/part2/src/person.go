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
func (p person) traverseMap(grid *[][]directionChanger) {
	for {
		p.move()
		pipe := (*grid)[p.i][p.j]
		if pipe.getPipe().traversed {
			break
		}
		p.rotate(pipe)
		// getPipe() retrieves the original pointer to Pipe from the DirectionChanger interface,
		// allowing us to modify the actual Pipe in the grid.
		pipe.getPipe().markTraversed()
	}
}

// Pointer for Modification: Using a pointer receiver allows us to modify the Traversed field
// of the actual Pipe instance in the grid, rather than a copy.
func (p *pipe) markTraversed() {
	p.traversed = true
}
func (p *person) move() {
	switch p.currentDirection {
	case north:
		p.i -= 1
	case east:
		p.j += 1
	case south:
		p.i += 1
	case west:
		p.j -= 1
	}
}

func (p *person) rotate(next directionChanger) {
	p.currentDirection = next.nextDirection(p.currentDirection)
}
