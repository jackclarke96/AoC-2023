# Problem 1

The input is 140 lines of 140 length strings.
Wherever there is a symbol, check for surrounding numbers and sum them.

# Solution 1

1. split strings into a slice based on newline. Make each string itself a slice so we have a 2d slice
2. iterate over the 2d slice. If `slice[i][j]` is a symbol, check for numbers in positions `slice[i][j-1]`, `slice[i][j+1]`, `slice[i-1][j-1]`, `slice[i-1][j]`, `slice[i-1][j+1]`, `slice[i+1][j-1]`, `slice[i+1][j]`, `slice[i+1][j+1]`
3. If find a number, check for number with string length > 1 by moving left and right
4. Once number added, change to '.'

# Problem 2

Add the product of two numbers if 
1. They're adjacent to the '*' symbol
2. They're the only two numbers adjacent to the symbol

# Solution 2

1. Add a check such that if symbol is *, call `checkSurroundingEntries` if so
2. In `checkSurroundingEntries`, rather thansum up all the numbers, add them to a slice. If the slices goes beyong length 2, break. Otherwise, multiply them and sum onto the running total.

# Improvements

use structs instead of strings as slice entries. 
Could also use an array since we know the size of the 2d slice up front.