# Plan

1. Parse into 2d matrix
2. Transpose it so have mat and tMat
3. iterate top to bottom through mat and tMat. Check for symmetry. Can use same logic then for both horizontal and vertical planes of symmetry

## Part 2

Rather than checking for absolute equality in rows, I will return the number of differences between rows

```go
func compareRows(row1, row2 []int) int {
	differences := 0
	for i := range row1 {
		if row1[i] != row2[i] {
			differences += 1
		}
	}
	return differences
}
```

Then rather than finding two symmetrical rows and using them as the start point, i will test each pair of consecutive rows, and check that the differences between them is 1

```go
func (g grid) checkForHorizontalSymmetry(iOneStart, iTwoStart int) bool {
	differences := 0
	for iOne, iTwo := iOneStart, iTwoStart; iOne >= 0 && iTwo < len(g); iOne, iTwo = iOne-1, iTwo+1 {
		differences += compareRows(g[iOne], g[iTwo])
		if differences > 1 {
			return false
		}
	}
	return differences == 1
}
```

Then can use the same logic otherwise.