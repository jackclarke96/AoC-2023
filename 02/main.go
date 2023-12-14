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
