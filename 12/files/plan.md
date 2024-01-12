# Problem 1

## Breaking down Problem 1

### Damaged Springs

```
#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1
```

#s are damaged springs. we list 1,1,3 for #.#.### as we have 1 spring damaged, gap, another damaged, gap, 3 damaged. ### is always 3, never 2,1 or 1,2.

### Unknown Springs

Some are unknown, however (we don't know if they're damaged or functional)

```
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
```

* In the first group there is exactly one way we can get 1,1,3 and that is #.#.###
* In the second group .??..??...?##. 1,1,3 we can have four arrangements the last ? must be a # and then we jusr have four spaces for 2 remaining hashes


### Plan

some sort of binomial modification? For the second example, we have (2 choose 1) and (2 choose 1) and (1 choose 1) multiplied e.g. in second group is 4

The first one is a little more awkward

need to fit 2 broken 1 functional... 3 choose 3? = 1

The final one has 10 combinations

?###???????? 3,2,1.

We have 3 hashed. Therefore the first and fifth must be '.'s
Then we have ??????? 7 places to fit in a pair of hashes, then a gap, then a hash
If we count the pair of hashes as a single space, and take away the gap that must appear, we get 5 spaces to choose 2 things. 5 choose 2 is 10.

Or:

Any length of hashes should actually be treated as #s followed by a . since the split is required.

So step 1 is to find any string of hashes of any of the defined lengths
best way to deal with this?

1. Work out what we know.

Maybe it's best to treat as though everything is a question mark,, then remove combinations that do not fit what is passed?  No. This would require actually generating each combination

e.g. more challenging ?###???????? 4,2,1. or  ?###???????? 4,2,1. or 

?###?????????? 4,3,1

then the first 3 hashes couldmake up a 4 OR they could make up a 3.

Perhaps treat as 2 separate problems. #### and ### then multiple the results. Will start coding this and see how it goes