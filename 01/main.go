package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegexp = regexp.MustCompile("[0-9]+")

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	stringSlice := strings.Split(input, "\n")
	total := 0
	for _, str := range stringSlice {
		// time.Sleep(10 * time.Millisecond)

		numbersString := iterateThroughString(str)
		numbers := parseNumbers(numbersString)
		fmt.Println("numbers", numbers)
		firstLast, err := getFirstLastNumStringBuilder(numbers)
		fmt.Println("firstLast", firstLast)

		if err != nil {
			log.Printf("Failed to combine numbers for %s: %v", str, err)
			continue
		}
		total += firstLast
		fmt.Println("total = ", total)

	}
	fmt.Println(total)
}

// take a slice of byte slices representing numbers,
// combine the first digit of the first number with the last digit of the last number
// return the combined number.
func getFirstLastNum(numbers [][]byte) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	}
	first := string(numbers[0])
	last := string(numbers[len(numbers)-1])
	return strconv.Atoi(string(first[0]) + string(last[len(last)-1]))
}

// string builder more efficient that +. Concat bytes then convert
func getFirstLastNumStringBuilder(numbers [][]byte) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	}
	first := numbers[0]
	last := numbers[len(numbers)-1]

	var builder strings.Builder

	builder.WriteByte(first[0])
	builder.WriteByte(last[len(last)-1])

	return strconv.Atoi(builder.String())

}

func getInput() (string, error) {

	content, err := os.ReadFile("./files/input.txt")
	if err != nil {
		return "", err
	}

	return string(content), nil

}

// use a regular expression to find all sequences of digits in a string
// return them as a slice of byte slices.

func parseNumbers(s string) [][]byte {
	return numberRegexp.FindAll([]byte(s), -1)
}
