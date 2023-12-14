package main

import (
	"testing"
)

var resultsMap = map[string]string{
	"abc1def2hij3": "13",
	"abone23four":  "14",
	"abctwonedef":  "21",
}

func TestIterateThroughString(t *testing.T) {
	for input, _ := range resultsMap {
		result := findFirstLastNum(input, "2")
		if result != resultsMap[input] {
			t.Errorf("For input '%v', expected output '%v', but got '%v'", input, resultsMap[input], result)
		}
	}
}
