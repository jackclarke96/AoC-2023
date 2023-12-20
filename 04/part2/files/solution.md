# Problem 2

For each row of the input, numbers on the right are our numbers. Numbers on the left are winning numbers. 
Find how many numbers `N` on the right match the numbers on the left.

Then, if we are evaluating the `i`th card, we add 1 card up to card `i+N`, card `i+1`, `i+2`,... `i + N`

# Solution:

Before round 1: 

| Round | Card 1 | Card 2 | Card 3 | Card 4 | card 5 | Card 6 |
|-------|--------|--------|--------|--------|--------|--------|
|   0   |   1    |   1    |   1    |   1    |   1    |   1    |

In round 1:

1 copy of card 1.
4 winners.
cards 2,3,4,5 get 1 copy each

| Round | Card 1 | Card 2 | Card 3 | Card 4 | card 5 | Card 6 |
|-------|--------|--------|--------|--------|--------|--------|
|   1   |   1    |   2    |   2    |   2    |   2    |   1    |

In round 2:

2 copies of card 2
2 winners
cards 3, 4 get 2 copies each

| Round | Card 1 | Card 2 | Card 3 | Card 4 | card 5 | Card 6 |
|-------|--------|--------|--------|--------|--------|--------|
|   2   |   2    |   2    |   4    |   4    |   1    |   1    |

In round 3:
4 copies of card 3
2 winners.
cards 4, 5 get 4 copies each

formula: card[i+{1,2,... #winners}] += card[i].quantity

Apply this formula iteratively.

Will update the gameStruct like below:

```go
type game struct {
	winningNumbers gameNumbers
	ourNumbers     gameNumbers
	numberOfCards  int
}
```

so that the number of cards can be tracked.

## Initial attempt

The initial (successful) attempt can be found in the archive folder

Trying to get to grips with receivers and pointers etc. so used a lot

One thing to not is that the below function:

```go
func (gs *gamesSlice) appendGameStruct(gameString string) {
	numbers := strings.Split(strings.Split(gameString, ":")[1], "|")
	winningNumberMap, ourNumberMap := make(gameNumbers), make(gameNumbers)
	winningNumberMap.storeNumbers(numbers[0])
	ourNumberMap.storeNumbers(numbers[1])
	*gs = append(*gs, game{winningNumberMap, ourNumberMap, 1})
}
```

requires a pointer to the gamesSlice be passed to the receiver. 

This is because we are modifying the slice itself by appending. If we were modifying the underlying array e.g. by modifying an element this would not be the case.

## Attempt 2

Found in `main.go`

Refactored for clarity of code. Perhaps went overboard on the receivers. Also promotes pure functions but largely the same algorithm.