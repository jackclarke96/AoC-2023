# Problem 1

Your toy boat has a starting speed of zero millimeters per millisecond. For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond.

so speed = time spent charging

if race time = T, and charging time = t, with T fixed integer and t variable integer

distance = speed * time in motion = t*(T-t)

For an input e.g.

Time:      7  15   30
Distance:  9  40  200

we want t*(T-t) > D e.g. for the first input, t*(7-t) > 9 i.e. when t is in 2, 3, 4 or 5

This will always be symmetrical with maximum in the middle, so we can just find the minimum value of t for which t*(T-t) > D. Call this t0. Then we have t0, T-t0 as the range.

T - 2*t0 + 1 is the number of vals in the range.

So the task is really only to find the minimum.

## Making this a little more complex

Make the main loop calculating number of ways to win run concurrently with other calculations.

Attempt 1 failed. It didn't work due to GoRoutine Closures, leading to inconsistent outcomes.

This is the "closure loop variable" problem.

below, `i` is shared across all goRoutines. Because goroutines may execute after the loop iteration in which they were launched, they might see same value of i.

To fix, pass `i` as a variable:

```go
for i := 0; i < len(Tvals); i++ {
    go func(idx int) {
        c <- getRange(convertStringToInt(Tvals[idx]), convertStringToInt(Dvals[idx]), c)
    }(i)
}
```

In a `for range` loop, the variables declared in the loop are re-declared in each iteration of the loop.  This means each goroutine launched within a for range loop gets its own copy of the variables, which is not affected by subsequent iterations. This is a safer alternative, therefore.




```go
for i := 0; i < len(Tvals); i++ {
    go getRange(convertStringToInt(Tvals[i]), convertStringToInt(Dvals[i]), c)
}
```


```go
func getMarginOfError(Tvals, Dvals []string) {
	total := 0
	c := make(chan Result)

	for i, T := range Tvals {
		go func(index int, T, D string) {
			c <- getRange(convertStringToInt(T), convertStringToInt(D), c)
		}(i, T, Dvals[i])
	}
	for range Tvals {
		result := <-c
		fmt.Println(result)
		if result.Err != nil {
			fmt.Println(result.Err)
		} else if total == 0 {
			total = result.Number
		} else {
			total *= result.Number
		}
	}
	fmt.Println(total)
}

func getRange(T, D int, c chan Result) Result {
	t := 1
	for t <= T/2 {
		if performCalculation(t, T, D) {
			return Result{T - 2*t + 1, nil}
		}
		t += 1
	}
	return Result{0, errors.New("It is not possible to guarantee a win of this race")}
}
```

## Part 2

Time:      71530
Distance:  940200

Basically the same problem but it should be optimised. 

I can use a binary search to find t0 and t0-1 such that:

t0*(T-t0) > D

but

(t0-1)*(T-(t0-1)) <= D