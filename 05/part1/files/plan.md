so input = 79

seed-to-soil map: 
50 98 2
52 50 48

if (input) falls in range:
 (col 2) to  (col 2 + col 3 - 1),
 then (ouput) = (input + (col 1 - col 2))

otherwise, output = input 

 79 -> 79 + 2 

 ## Using a Binary Search

 I will represent the maps as a slices of slices, with each slice of length 3 e.g. above:

 var soilToSeed = [][]int{
    {50, 98, 2},
    {52, 50, 48}
 }

 Then sort by col2. This will only work if the ranges (col 2) to  (col 2 + col 3 - 1) DO NOT OVERLAP

 Do they? No. Because if they overlapped there wouldn't be the required 1:1 relationship between then maps.

 So we can sort by and apply the search based on col2

 then apply the binary search

 need: 
 1. A function to read the input and split into each map
 2. a function to conver tte map into [][]int
 3. A function to reorder
 4. A binary search function to pass the inputs through 1 by 1

## Binary Search Algorithm

Let min = 0 and max = n-1.
If max < min, then stop: target is not present in array. Return -1.
Compute guess as the average of max and min, rounded down (so that it is an integer).
If array[guess] equals target, then stop. You found it! Return guess.
If the guess was too low, that is, array[guess] < target, then set min = guess + 1.
Otherwise, the guess was too high. Set max = guess - 1.
Go back to step 2.

