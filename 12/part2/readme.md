# Part Two

As you look out at the field of springs, you feel like there are way more springs than the condition records list. When you examine the records, you discover that they were actually folded up this whole time!

To unfold the records, on each row, replace the list of spring conditions with five copies of itself (separated by ?) and replace the list of contiguous groups of damaged springs with five copies of itself (separated by ,).

So, this row:

```
.# 1
```

Would become:

```
.#?.#?.#?.#?.# 1,1,1,1,1
```

The first line of the above example would become:

```
???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3
```

In the above example, after unfolding, the number of possible arrangements for some rows is now much larger:

```
???.### 1,1,3 - 1 arrangement
.??..??...?##. 1,1,3 - 16384 arrangements
?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
????.#...#... 4,1,1 - 16 arrangements
????.######..#####. 1,6,5 - 2500 arrangements
?###???????? 3,2,1 - 506250 arrangements
```

After unfolding, adding all of the possible arrangement counts together produces 525152.

Unfold your condition records; what is the new sum of possible arrangement counts?

# Planned Solution

This is substantially more complex. Old solution not feasible. Some maps have 10 "?"s so 2^50 paths to explore.

## Using Memoization

If we define the value of a node to be the number of valid combinations possible from that point onwards, then the value of a node is equal to the sum of the value of its children.

If we define the state to be a combination of:
* The index of the spring currently being evaluated
* number of hashes placed up to and including the current index of the string slice
* The number of consecutive hashes leading up to the current index of the string slice

Then the number of combinations from that point onwards is the same regardless of how the previous hashes are ordered.
This means that the value of the children will be the same as the value of the children from the first exploration.

Defining the base case of the problem to be a fully arranged string, we can then essentially backpropagate its value so that equivalent states need never be explored more than once.

## Data Structures

This key allows for the unique representation of equivalent states.

```go
type memoKey struct {
	index                    int
	numHashes                int
	currentConsecutiveHashes int
}
```

# Running the code

* `cd` into `src`.
* `go run main.go isValidPartial.go isValidCombination.go helpers.go generateCombinations.go`