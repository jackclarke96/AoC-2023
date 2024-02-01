# Part 1

## Understabding the problem

* Rounded rocks (O), square rocks (#) and empty places (.)
* Can tilt the mirror north, south, east or west.
* Tilt north. All round rocks roll upwards as far as possible. #s stay in place
* sum up the distances from the south edge of the round rocks i.e. len(grid) - index(0) for each rock

OOOO.#.O.. 10
OO..#....#  9
OO..O##..O  8
O..#.OO...  7
........#.  6
..#....#.#  5
..O..#.O.O  4
..O.......  3
#....###..  2
#....#....  1

## Plan

Tilting north - each rock is only interested in what is above it - not what is to the side. Maybe can execute concurrently using this.

Data structures: could just use slices with values

alternatively, structs 

```{
    type string // "#", "O", "."
    fixed bool // true if #
    i, j int
}
```

swap if fixed !=true and type == "."

When iterating through in initial parse could track 'maxNorth'. Update it when encounter a # and set it. Maybe could have a map[i,j]int to store nrothMax

Will just create slices of the columns to start. Keep simple. 

1. Parse input into 2d slice. `parseInputIntogrid`.
2. for each column, rearrange. To do this efficiently, find indices of #s and just move beneath. `tiltColumnNorth`. Since each column is independent of each other column, can use pointers to avoid creating 100 copies of grid.
3. Use orchestrator function, `orchestrateTilt`. This will handle channel and creating go routines

# Part 2

In part 2, need to perform full spins, N, W, S E.

Need to find the grid score after 1,000,000,000 spins.

To do this, make use of the fact that the grid will eventually enter a cycle.

I can detect this cycle by checking whether a grid has appeared before. If it has, we have entered the cycle.

I will create a slice, `encounteredGrids` to which I append the current representation of the grid after each subsequent spin. After each spin, iterate through t `encounteredGrids` checking for the current representation of the grid.

If the grid is found, get the index `startIndex`  of the grid as it appears in `encounteredGrids`.

Use this, together with `currentIndex`, the second index at which the duplicate grid has been found, to get `cycleLength = currentIndex-startIndex`

The grid after 1,000,000,000 spins will match the grid whose:

* index is between `startIndex` and `currentIndex`
* index, when divided by `cycleLength`, has the same remainder as when 1,000,000,000 is divided by `cycleLength` (since adding on cycleLength repeatedly to it will give 1,000,000,000). Actually, should add one onto the index first to account for indexing starting at 0
* i.e. when  `(cycleIndex+1)%cycleLength == offset` for `cycleIndex := startIndex; cycleIndex < currentIndex; cycleIndex++` and `offset := 1000000000 % cycleLength`