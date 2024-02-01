package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type colourCounts struct {
	Red   int
	Green int
	Blue  int
}

// Declare globally to avoid reinitialising
var numberRegexp = regexp.MustCompile("[0-9]+")
var colourRegExp = regexp.MustCompile(`red|blue|green`)

func main() {

	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	total := executeMain(string(input))
	log.Println(total)
}

func executeMain(input string) int {
	total := 0
	chunkSize := 10
	inputSlice := strings.Split(string(input), "\n")
	inputChunks := chunkSlice(inputSlice, chunkSize)

	// Create channel to which each goRoutine can communicate scores across its 10 games
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

// For each game  use isGameValid method to compare the maximum possible counts of each cube with the number drawn in the game to determine whether the game is valid.
// Sum the IDs of the valid games and return it by initialising a counter to 0 and summing each valid games ID
func processChunk(chunk []string) int {
	total := 0

	maxCounts := colourCounts{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	for _, game := range chunk {
		orderedColours := colourRegExp.FindAllString(game, -1)
		orderedNumbers := numberRegexp.FindAllString(game, -1)
		IDString, orderedQuantities := orderedNumbers[0], orderedNumbers[1:]

		if isGameValid(orderedColours, orderedQuantities, maxCounts) {
			if id, err := strconv.Atoi(IDString); err != nil {
				log.Fatalf("Error extracting int ID from string: %v", err)
			} else {
				total += id
			}
		}
	}
	return total
}

// Iterate through coloursDrawn slice e.g. []string{"Red", "Green", "Green", "Blue"}
// and numberOfTimes slice []string{"10", "12", "20", "2"}
// and return map with maximum number of drawn cubes for each colour
func isGameValid(coloursDrawn, numberOfTimes []string, maxCounts colourCounts) bool {
	for j, colour := range coloursDrawn {
		timesDrawn, err := strconv.Atoi(numberOfTimes[j])
		if err != nil {
			log.Fatalf("Error extracting number: %v", err)
		}
		switch colour {
		case "red":
			if maxCounts.Red < timesDrawn {
				return false
			}
		case "green":
			if maxCounts.Green < timesDrawn {
				return false
			}
		case "blue":
			if maxCounts.Blue < timesDrawn {
				return false
			}
		}
	}
	return true
}

// We have a slice of length N and chunkSize n. Then we need of N/n chunks in case of perfect integer division or N/n + 1 otherwise
func chunkSlice(slice []string, chunkSize int) [][]string {
	sliceLength := len(slice)

	// Ensure that we round up
	numChunks := (sliceLength + chunkSize - 1) / chunkSize
	chunks := make([][]string, numChunks)

	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > sliceLength {
			end = sliceLength
		}
		chunks[i] = slice[start:end]
	}
	return chunks
}
