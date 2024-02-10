# --- Day 11: Cosmic Expansion ---

You continue following signs for "Hot Springs" and eventually come across an observatory. The Elf within turns out to be a researcher studying cosmic expansion using the giant telescope here.

He doesn't know anything about the missing machine parts; he's only visiting for this research project. However, he confirms that the hot springs are the next-closest area likely to have people; he'll even take you straight there once he's done with today's observation analysis.

Maybe you can help him with the analysis to speed things up?

The researcher has collected a bunch of data and compiled the data into a single giant image (your puzzle input). The image includes empty space (.) and galaxies (#). For example:

```
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
```

The researcher is trying to figure out the sum of the lengths of the shortest path between every pair of galaxies. However, there's a catch: the universe expanded in the time it took the light from those galaxies to reach the observatory.

Due to something involving gravitational effects, only some space expands. In fact, the result is that any rows or columns that contain no galaxies should all actually be twice as big.

In the above example, three columns and two rows contain no galaxies:

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

These rows and columns need to be twice as big; the result of cosmic expansion therefore looks like this:

```
....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......
```

Equipped with this expanded universe, the shortest path between every pair of galaxies can be found. It can help to assign every galaxy a unique number:

```
....1........
.........2...
3............
.............
.............
........4....
.5...........
............6
.............
.............
.........7...
8....9.......
```

In these 9 galaxies, there are 36 pairs. Only count each pair once; order within the pair doesn't matter. For each pair, find any shortest path between the two galaxies using only steps that move up, down, left, or right exactly one . or # at a time. (The shortest path between two galaxies is allowed to pass through another galaxy.)

For example, here is one of the shortest paths between galaxies 5 and 9:

```
....1........
.........2...
3............
.............
.............
........4....
.5...........
.##.........6
..##.........
...##........
....##...7...
8....9.......
```

This path has length 9 because it takes a minimum of nine steps to get from galaxy 5 to galaxy 9 (the eight locations marked # plus the step onto galaxy 9 itself). Here are some other example shortest path lengths:

Between galaxy 1 and galaxy 7: 15
Between galaxy 3 and galaxy 6: 17
Between galaxy 8 and galaxy 9: 5
In this example, after expanding the universe, the sum of the shortest path between all 36 pairs of galaxies is 374.

Expand the universe, then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?

# Plan

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


# --- Part Two ---
The galaxies are much older (and thus much farther apart) than the researcher initially estimated.

Now, instead of the expansion you did before, make each empty row or column one million times larger. That is, each empty row should be replaced with 1000000 empty rows, and each empty column should be replaced with 1000000 empty columns.

(In the example above, if each empty row or column were merely 10 times larger, the sum of the shortest paths between every pair of galaxies would be 1030. If each empty row or column were merely 100 times larger, the sum of the shortest paths between every pair of galaxies would be 8410. However, your universe will need to expand far beyond these values.)

Starting with the same initial image, expand the universe according to these new rules, then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?

# Plan

No change needed other than rather than adding 1 for an expansion, add 999999
