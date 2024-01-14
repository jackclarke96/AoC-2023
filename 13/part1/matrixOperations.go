package main

func (g grid) transpose() grid {
	numRows := len(g)
	numCols := len(g[0])

	newGrid := make(grid, numCols)
	for i := range newGrid {
		newGrid[i] = make([]int, numRows)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			newGrid[j][i] = g[i][j]
		}
	}

	return newGrid
}

func (g grid) reflectInHorizontalPlane() grid {

	numRows := len(g)
	numCols := len(g[0])

	newGrid := make(grid, numRows)
	for i := range newGrid {
		newGrid[i] = make([]int, numCols)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			newGrid[i][j] = g[numRows-1-i][j]
		}
	}

	return newGrid
}

// impure but memory efficient as we directly edit the ints themselves but the pointers to the address in memory that g[i][j] references remains the same
// to make a pure version i would need to make an empty 2d slice and fill it with its own set of integers
func (g grid) reflectInVerticalPlaneInPlace() grid {
	rowLength := len(g[0])
	for i, _ := range g {
		for j, k := 0, rowLength-1; j < k; j, k = j+1, k-1 {
			g[i][j], g[i][rowLength-1-j] = g[i][rowLength-1-j], g[i][j]
		}
	}
	return g
}

func compareRows(row1, row2 []int) bool {
	for i := range row1 {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}
