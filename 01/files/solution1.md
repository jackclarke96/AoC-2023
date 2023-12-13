# Initial Solution:

The below works.


```Go
package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	stringSlice := strings.Split(input, "\n")
	total := 0
	for _, str := range stringSlice {
		bss := parseNumbers(str)
		combinedNumbers, _ := combineNumbers(bss)
		total += combinedNumbers
	}
}

func combineNumbers(bss [][]byte) (int, error) {
	if len(bss) == 0 {
		return 0, nil
	} else {
		bs1 := string(bss[0])
		bs2 := string(bss[len(bss)-1])
		return strconv.Atoi(string(bs1[0]) + string(bs2[len(bs2)-1]))
	}

}

func getInput() string {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(content)

}

func parseNumbers(s string) [][]byte {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAll([]byte(s), -1)
}
```

## Improvements and Optimisations

1. Use `strings.builder` instead of '+' with lots of type conversions for more efficient concatenation.
2. In parseNumbers, the regular expression is compiled every time the function is called. Compile once and reuse instead.
3. Add error handling.
4. Rename some variables and add some comments. Clean things up a bit


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

var numberRegexp = regexp.MustCompile("[0-9]+")

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	stringSlice := strings.Split(input, "\n")
	total := 0
	for _, str := range stringSlice {
		numbers := parseNumbers(str)
		combinedNumbers, err := combineNumbersStringBuilder(numbers)
		if err != nil {
			log.Printf("Failed to combine numbers for %s: %v", str, err)
			continue
		}
		total += combinedNumbers
	}
	fmt.Println(total)
}

// take a slice of byte slices representing numbers,
// combine the first digit of the first number with the last digit of the last number
// return the combined number.
func combineNumbers(numbers [][]byte) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	}
	first := string(numbers[0])
	last := string(numbers[len(numbers)-1])
	return strconv.Atoi(string(first[0]) + string(last[len(last)-1]))

}

// string builder more efficient that +. Concat bytes then convert
func combineNumbersStringBuilder(numbers [][]byte) (int, error) {
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

	content, err := os.ReadFile("input1.txt")
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
```