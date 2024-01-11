# Part 1

## Data Structures

I will represent as a 2d array of structs.

Each struct will look like the below 

```go
type Pipe struct {
    Type         string
    Traversed    bool
    ForwardDistance, BackwardDistance     int
}
```

Will also create function(s) to evaluate which types can interact under different x,y conditions e.g. `"7L"` cannot but `"L7"` can

## Algorithm

Just start at the beginning of the loop and work round back to the start. Count how many pipes traversed divide by 2 for furthest distance away

# Part 2

Use the code from part 1 to mark out the loop.

Then iterate through the matriox marking anything non-traversed as of `nil` type. This will mean fewer lookups on pipes to see if they have been traversed.

# Things learnt

A ton on memory, pointers and interfaces

