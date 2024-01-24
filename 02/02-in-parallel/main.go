package main

import (
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

	filePath := os.Args[1]

	total := executeMain(filePath)

	log.Println(total)
}

func executeMain(filePath string) int {

	input, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read input file.")
	}

	total := 0
	chunkSize := 10
	inputSlice := strings.Split(string(input), "\n")
	inputChunks := chunkSlice(inputSlice, chunkSize)

	// Create a channel to which each goRoutine can communicate scores achieved by its chunk of 10 games
	gameScores := make(chan int)

	// Use a function literal with index of the chunk passed into it as a variable.
	// This means each 'processChunk' gets its own copy of the index variable
	for index := range inputChunks {
		go func(i int) {
			// pass the result of processChunk function into the gameScores channel
			gameScores <- processChunk(inputChunks[i])
		}(index)
	}

	// wait for each goroutine to send its computed score to the channel and then add it to the total.
	for chanResponses := 0; chanResponses < len(inputChunks); chanResponses++ {
		total += <-gameScores
	}

	return total
}

func processChunk(chunk []string) int {
	total := 0
	for _, game := range chunk {
		numberMatch := numberRegexp.FindAllString(game, -1)[1:]
		colourMatch := colourRegExp.FindAllString(game, -1)
		total += getMaximumOfEachColour(colourMatch, numberMatch).getCubes()
	}
	return total
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

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
