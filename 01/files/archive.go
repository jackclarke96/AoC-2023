package archive

import (
	"regexp"
	"strconv"
	"strings"
)

// use a regular expression to find all sequences of digits in a string
// return them as a slice of byte slices.

func parseNumbers(s string) [][]byte {
	return numberRegexp.FindAll([]byte(s), -1)
}

var numberRegexp = regexp.MustCompile("[0-9]+")

// string builder more efficient than +. Concatenate bytes then convert
func getFirstLastNum(numbers [][]byte) (int, error) {
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
