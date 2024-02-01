package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	fmt.Println(executeMain(input))
}

func executeMain(bs []byte) int {
	bss := parseInputIntoByteSlices(bs)
	total := 0
	for _, byteSlice := range bss {
		total += processByteSlice(byteSlice)
	}
	return total
}

func processByteSlice(bs []byte) int {
	total := 0
	for _, b := range bs {
		total += int(b)
		total *= 17
		total = total % 256
	}
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
