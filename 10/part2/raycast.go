package main

// mark any untraversed pipes as nil to make comparisons in Raycast more simple.
func markUntraversedAsNil(grid [][]DirectionChanger) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != nil && !grid[i][j].GetPipe().Traversed {
				grid[i][j] = nil
			}
		}
	}
}

// Perform RayCast Algorithm
func rayCast(grid [][]DirectionChanger) int {
	inside := 0
	for _, path := range grid {
		boundariesCrossed := 0
		boundary := []PipeType{}
		for _, cell := range path {
			if cell == nil {
				if boundariesCrossed%2 == 1 {
					inside += 1
				}
				boundary = []PipeType{}
			} else {
				boundary, boundariesCrossed = processBoundary(cell, boundary, boundariesCrossed)
			}
		}
	}
	return inside
}

// THEORY HERE:
// L--J doesnt count as a border cross since you never actually make it to the gap between them. The same applies to F----7. Sort of like going around a U shape outside remains outside
// L--7 has a north and south so you would have to actually cross over to get past. So it is a proper boundary. Same applies to F----J
// | is always a boundary
// - is never a boundary

func processBoundary(cell DirectionChanger, boundary []PipeType, boundariesCrossed int) ([]PipeType, int) {
	if cell.GetPipe().Type == NS {
		boundariesCrossed++
	} else {
		boundary = append(boundary, cell.GetPipe().Type)
		if (boundary[0] == NE && boundary[len(boundary)-1] == SW) || (boundary[0] == SE && boundary[len(boundary)-1] == NW) {
			boundariesCrossed++
			boundary = []PipeType{}
		} else if (boundary[0] == NE && boundary[len(boundary)-1] == NW) || (boundary[0] == SE && boundary[len(boundary)-1] == SW) {
			boundary = []PipeType{}
		}
	}
	return boundary, boundariesCrossed
}
