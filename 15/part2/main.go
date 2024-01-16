package main

import (
	"fmt"
	"log"
	"os"
)

var HASHMAP = map[string]int{}

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	fmt.Println(executeMain(input))
}

func asciiToDigit(ch byte) int {
	return int(ch) - '0'
}

func executeMain(bs []byte) int {
	// bss := parseInputIntoByteSlices(bs)
	total := 0
	// for _, byteSlice := range bss {
	// 	label, boxNumber, action, focalLength := splitByteSliceIntoComponents(byteSlice)
	// 	if action == "=" {
	// 		upsertBoxEntry(label, boxNumber, focalLength)
	// 	} else {
	// 		deleteBoxEntry(label, boxNumber)
	// 	}
	// }
	fmt.Println([]byte("="))
	return total
}

func splitByteSliceIntoComponents(bs []byte) (string, int, string, int) {
	var label string
	var boxNumber int
	var focalLength int
	var action string

	for i, b := range bs {
		// single quotes compares ascii values
		if b == '-' || b == '=' {
			label = string(bs[:i])
			boxNumber = getBoxNumber(bs[:i])
			action = string(b)
			break
		}
	}

	if action == "=" {
		fmt.Println(string(bs))
		focalLength = asciiToDigit(bs[len(bs)-1])
	}
	return label, boxNumber, action, focalLength
}

func getBoxNumber(bs []byte) int {
	if total, ok := HASHMAP[string(bs)]; ok == true {
		return total
	}
	total := 0
	for _, b := range bs {
		total += int(b)
		total *= 17
		total = total % 256
	}
	HASHMAP[string(bs)] = total

	return total
}

func parseInputIntoByteSlices(bs []byte) [][]byte {
	startIndex := 0
	divider := []byte(",")[0]
	bss := [][]byte{}
	for i, b := range bs {
		if b == divider {
			bss = append(bss, bs[startIndex:i])
			startIndex = i + 1
		}
	}
	// account for final string
	bss = append(bss, bs[startIndex:])
	fmt.Println(bss[0][0])
	return bss
}
