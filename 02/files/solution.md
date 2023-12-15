# Part 1

## Summary of Problem 

There is a finite number of cubes - red, green, blue.

Each time, a random number of cubes is shown and put back. This is done a few times per game

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
...

Which games are possible if there are 12 red, 13 green, and 14 blue?

Get the sum of the IDs of the games.

## Plan

1. Create a map `colourMap` with the max possible vlaue of each colour
2. Load the input and split into a slice of length 100 based on `\n`
3. For each element, create a slice of numbers `numbersMatch` and a slice of colours `coloursMatch` using `regExp.findAllStrings(element)`
4. iterate through the colour slice. If `numbersSlice[index] < `colourMap[colour]` the game is impossible

```go
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var colourMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var numberRegexp = regexp.MustCompile("[0-9]+")
var colourRegExp = regexp.MustCompile(`red|blue|green`)

func main() {
	total := 0
	input, err := os.ReadFile("./files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file.")
	}

	inputSlice := strings.Split(string(input), "\n")

	for i, game := range inputSlice {
		valid := true
		numberMatch := numberRegexp.FindAllString(game, -1)[1:]
		colourMatch := colourRegExp.FindAllString(game, -1)
		for j, colour := range colourMatch {
			num, err := strconv.Atoi(numberMatch[j])
			if err != nil {
				log.Println("Error extracting number")
				continue
			}
			if (colourMap[colour]) < num {
				valid = false
				break
			}

		}
		if valid == true {
			total += i + 1
		}
	}
	fmt.Println(total)
}
```

# Part 2

## Part 2: Summary of Problem

Want to work out what the fewest number of each colour needs to be in order for the games to be possible.

## Plan

1. Take maximum of each colour across each game and add it to a map
2. multiply the maximums

```go
type colourMap map[string]int

func getMaximumOfEachColour(coloursDrawn []string, numberOfTimes []string) colourMap {

	maxMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for j, colour := range coloursDrawn {
		timesDrawn, err := strconv.Atoi(numberOfTimes[j])
		if err != nil {
			log.Println("Error extracting number")
			continue
		}
		if (maxMap[colour]) < timesDrawn {
			maxMap[colour] = timesDrawn
		}
	}
	return maxMap
}

func (c colourMap) getCubes() int {
	return c["red"] * c["green"] * c["blue"]
}
```

## Optimising

1. Since we know all the colours in the maxMap in advance, a `struct` should be used.
2. getMaximumOfEachColour function can be reused in problem 1. 

```go
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ColourCounts struct {
	Red   int
	Green int
	Blue  int
}

var myColourMap ColourCounts = ColourCounts{
	Red:   12,
	Green: 13,
	Blue:  14,
}

var numberRegexp = regexp.MustCompile("[0-9]+")
var colourRegExp = regexp.MustCompile(`red|blue|green`)

func main() {

	problem := os.Args[1]
	if !(problem == "1" || problem == "2") {
		log.Fatalf("Please provide number indicating which problem is to be solved")
	}

	total := 0
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file.")
	}

	inputSlice := strings.Split(string(input), "\n")

	for i, game := range inputSlice {
		numberMatch := numberRegexp.FindAllString(game, -1)[1:]
		colourMatch := colourRegExp.FindAllString(game, -1)

		if problem == "1" {
			if checkGameValid(colourMatch, numberMatch) {
				total += i + 1
			}
		} else {
			total += getMaximumOfEachColour(colourMatch, numberMatch).getCubes()
		}
	}
	fmt.Println(total)
}

func getMaximumOfEachColour(coloursDrawn []string, numberOfTimes []string) ColourCounts {

	maxCounts := ColourCounts{
		Red:   0,
		Blue:  0,
		Green: 0,
	}
	for j, colour := range coloursDrawn {
		timesDrawn, err := strconv.Atoi(numberOfTimes[j])
		if err != nil {
			log.Println("Error extracting number")
			continue
		}
		switch colour {
		case "red":
			if maxCounts.Red < timesDrawn {
				maxCounts.Red = timesDrawn
			}
		case "green":
			if maxCounts.Green < timesDrawn {
				maxCounts.Green = timesDrawn
			}
		case "blue":
			if maxCounts.Blue < timesDrawn {
				maxCounts.Blue = timesDrawn
			}
		default:
			log.Printf("Unknown color: %s", colour)
		}
	}
	return maxCounts
}

func (c ColourCounts) getCubes() int {
	return c.Red * c.Blue * c.Green
}

func checkGameValid(coloursDrawn []string, numberOfTimes []string) bool {
	maxCounts := getMaximumOfEachColour(coloursDrawn, numberOfTimes)
	return maxCounts.Blue <= myColourMap.Blue &&
		maxCounts.Green <= myColourMap.Green &&
		maxCounts.Red <= myColourMap.Red
}
```

Note that even though `getMaximumOfEachColour` deals with the `maxCounts` struct, we do not need to use a pointer.
Had we edited a variable declared outside the function, we would need to edit the pointer so that the
actual values were changed, as opposed to a copy of them.

Editing a struct in the same function scope edits it directly. If we edit the struct by passing it into another 
function as a receiver or as a parameter, a copy is made, and so instead we would need to use a pointer.

## Making Concurrent

No real reason to do this other than to practice with it.

However, it might be worth it if the dataset were larger.

1. Ingest the input but split into chunks - smaller slices of length 10. One go routine will handle each chunk, and return its own total.
2. Create a channel for the `goRoutines` to communicate back to the main function. The channel uses type integer as the totals are of type integer  `c := make(chan int)`
3. Create function to handle a single chunk

```go
func processChunk(chunk []string, problem string, chunkNumber int, chunckSize int, c chan int) {
	total := 0
	for i, game := range chunk {
		numberMatch := numberRegexp.FindAllString(game, -1)[1:]
		colourMatch := colourRegExp.FindAllString(game, -1)

		if problem == "1" {
			if checkGameValid(colourMatch, numberMatch) {
				total += i + 1 + chunkNumber*chunckSize
			}
		} else {
			total += getMaximumOfEachColour(colourMatch, numberMatch).getCubes()
		}
	}
	c <- total
}
```

4. Call `processChunk` using `go` keyword to spawn new `goRoutines`, and wait for the channel to receive all responses

```go
for chunkNum, chunk := range inputChunks {
	go processChunk(chunk, problem, chunkNum, chunkSize, c)
}

for chanResponses := 0; chanResponses < len(inputChunks); chanResponses++ {
	total += <-c
}
```
