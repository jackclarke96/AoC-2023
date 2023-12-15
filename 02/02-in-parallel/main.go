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

	problem, filePath := os.Args[1], os.Args[2]

	if !(problem == "1" || problem == "2") {
		log.Fatalf("Please provide number indicating which problem is to be solved")
	}

	total := executeMain(problem, filePath)

	log.Println(total)
}

func executeMain(problem string, filePath string) int {
	total := 0
	input, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read input file.")
	}

	chunkSize := 10
	inputSlice := strings.Split(string(input), "\n")
	inputChunks := chunkSlice(inputSlice, chunkSize)

	c := make(chan int)

	for chunkNum, chunk := range inputChunks {
		go processChunk(chunk, problem, chunkNum, chunkSize, c)
	}

	for chanResponses := 0; chanResponses < len(inputChunks); chanResponses++ {
		total += <-c
	}

	return total

}

func processChunk(chunk []string, problem string, chunkNumber int, chunkSize int, c chan int) {
	total := 0
	for i, game := range chunk {
		numberMatch := numberRegexp.FindAllString(game, -1)[1:]
		colourMatch := colourRegExp.FindAllString(game, -1)

		if problem == "1" {
			if checkGameValid(colourMatch, numberMatch) {
				total += i + 1 + chunkNumber*chunkSize
			}
		} else {
			total += getMaximumOfEachColour(colourMatch, numberMatch).getCubes()
		}

	}
	c <- total
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
