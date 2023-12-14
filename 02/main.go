package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type colourMap map[string]int

var myColourMap = colourMap{
	"red":   12,
	"green": 13,
	"blue":  14,
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
	fmt.Println(c)
	fmt.Println(c["red"] * c["green"] * c["blue"])
	return c["red"] * c["green"] * c["blue"]
}

func checkGameValid(coloursDrawn []string, numberOfTimes []string) bool {
	for j, colour := range coloursDrawn {
		num, err := strconv.Atoi(numberOfTimes[j])
		if err != nil {
			log.Println("Error extracting number")
			continue
		}
		if (myColourMap[colour]) < num {
			return false
		}
	}
	return true
}
