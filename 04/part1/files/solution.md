# Problem 1

For each row of the input, numbers on the right are our numbers. Numbers on the left are winning numbers. 
Find how many numbers `n` on the right match the numbers on the left, then take `2^(n-1)` to get points.
Sum across all rows.

What is the best data structure to use for this?

A single row from the inout looks like this.

Card   1: 74  8  2 86 40 25 93 17 61 32 | 65 25 73 55 75 94 54 99 53 17 89  4 44 13 15 32 57 92  8 21 74 64  5 87 24

## Options for data structures

### Slices/Arrays

**Maps** are the most efficient in terms of lookup performance. They provide constant-time lookups which are essential when dealing with large sets of data where you need to quickly check if an item exists. Maps provide the most efficient lookup with O(1) average time complexity. This is because maps are implemented as hash tables, allowing for fast retrieval based on keys. 

**Arrays and slices** are less efficient for lookups as they require linear time to search through the elements. Lookup involves iterating through each element to find a match, resulting in O(n) time complexity per lookup.

**Direct string manipulation** inefficient and complex

Will use two maps.
