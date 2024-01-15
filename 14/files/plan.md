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
}```

swap if fixed !=true and type == "."

When iterating through in initial parse could track 'maxNorth'. Update it when encounter a # and set it. Maybe could have a map[i,j]int to store nrothMax

Will just create slices of the columns to start. Keep simple. 

1. Parse input into 2d slice. `parseInputIntogrid`.
2. for each column, rearrange. To do this efficiently, find indices of #s and just move beneath. `tiltColumnNorth`. Since each column is independent of each other column, can use pointers to avoid creating 100 copies of grid.
3. Use orchestrator function, `orchestrateTilt`. This will handle channel and creating go routines
