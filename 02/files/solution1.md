## Summary of Problem 

There is a finite number of cubes - red, green, blue.

Each time, a random number of cubes is shown and put back. This is done a few times per game

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
...

Whuich games are possible if there are 12 red, 13 green, and 14 blue?

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