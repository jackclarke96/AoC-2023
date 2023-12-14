package main

import (
	"regexp"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var numberWordsPattern = `one|two|three|four|five|six|seven|eight|nine`
var numberDigitsPattern = `1|2|3|4|5|6|7|8|9`
var reWords = regexp.MustCompile(numberWordsPattern)
var reDigits = regexp.MustCompile(numberDigitsPattern)

func findFirstLastNum(s string, problem string) string {

	myNum := findFirstNumberForward(s, problem)
	if myNum != "" {
		return myNum + findFirstNumberBackward(s, problem)
	}
	return ""
}

func findFirstNumberInSubstring(substr string, problem string) string {
	numberMatch := reDigits.FindString(substr)
	if numberMatch != "" {
		return numberMatch
	}
	if problem == "2" {
		// alternative combined syntax for conditionals and assigining vars
		if stringMatch := reWords.FindString(substr); stringMatch != "" {
			return numberMap[stringMatch]
		}
	}

	return ""
}

func findFirstNumberForward(s string, problem string) string {
	for i := 0; i < len(s); i++ {
		substr := s[:i+1]
		if num := findFirstNumberInSubstring(substr, problem); num != "" {
			return num
		}
	}
	return ""
}

func findFirstNumberBackward(s string, problem string) string {
	for i := len(s); i >= 0; i-- {
		substr := s[i:]
		if num := findFirstNumberInSubstring(substr, problem); num != "" {
			return num
		}
	}
	return ""
}
