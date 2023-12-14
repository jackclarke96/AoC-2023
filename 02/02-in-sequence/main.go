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
	input, err := os.ReadFile("./files/input.txt")
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
