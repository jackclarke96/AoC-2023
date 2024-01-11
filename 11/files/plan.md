# Part 1

## Breaking this down

1. The growth part basically says For any empty row or column, insert a copy of that empty row or column directly beneath or to the right of it.
2. shortest path is just steps down + steps across

* Can parse into a 2d slice (converting hashes to numbers), then add columns and rows in another iteration
* Then, write, or use, a function to find all pairs of numbers 
* Then calculate distance

## Data Structures 

Since already pretty confident with slices and iterating through them, I will use a maps and use string manipulation and regex to work with the grid

```go
type galaxyStruct struct {
	i, j int
}

type expansions map[int]int

type calculationInput struct {
	galaxyStart, galaxyEnd                   galaxyStruct
	horizontalExpansions, verticalExpansions expansions
}
```

I will 

* Use a slice of `galaxyStructs` to find all `i` and `j` at which a galaxy exists. Store only the coordinates in which there are actually galaxies.
* Store expansion rows and columns as maps of ints (1 = added column r row, 0 = havent), then loop through and add on the distance. or even, each row/column could store number of expansions til this point cumulatively.

e.g. here

```
   v  v  v
 ...#......
 .......#..
 #.........
>..........<
 ......#...
 .#........
 .........#
>..........<
 .......#..
 #...#.....
   ^  ^  ^
```

moving horizontally:

* There are 0 expansions for j = 0,1. 
* There is 1 expansion for j = 2,3,4
* There are 2 expansions for j = 5,6,7
* 3 expansions for j = 8, 9

So a movement from 0 to 6, for example, becomes a movement of 8 in the j direction instead of 6.

Therefore there is no need to alter the map.

Use `Abs(horizontal distance between points A and B in the original map) + Abs(expansion number at B - expansion number at A)`

and similar for vertical.


## Part 2

No change needed other than rather than adding 1 for an expansion, add 999999