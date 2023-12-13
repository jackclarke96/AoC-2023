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

func iterateThroughString(s string) string {
	numberWordsPattern := `one|two|three|four|five|six|seven|eight|nine`
	numberDigitsPattern := `1|2|3|4|5|6|7|8|9`
	reWords := regexp.MustCompile(numberWordsPattern)
	reDigits := regexp.MustCompile(numberDigitsPattern)

	myNum := findFirstNumberForward(s, reWords, reDigits)

	if myNum != "" {
		return myNum + findFirstNumberBackward(s, reWords, reDigits)
	}
	return ""
}

func findFirstNumberInSubstring(substr string, reWords, reDigits *regexp.Regexp) string {
	stringMatch := reWords.FindString(substr)
	if stringMatch != "" {
		return numberMap[stringMatch]
	}
	// alternative combined syntax
	if numberMatch := reDigits.FindString(substr); numberMatch != "" {
		return numberMatch
	}
	return ""
}

func findFirstNumberForward(s string, reWords, reDigits *regexp.Regexp) string {
	for i := 0; i < len(s); i++ {
		substr := s[:i+1]
		if num := findFirstNumberInSubstring(substr, reWords, reDigits); num != "" {
			return num
		}
	}
	return ""
}

func findFirstNumberBackward(s string, reWords, reDigits *regexp.Regexp) string {
	for i := len(s); i >= 0; i-- {
		substr := s[i:]
		if num := findFirstNumberInSubstring(substr, reWords, reDigits); num != "" {
			return num
		}
	}
	return ""
}

// func reverseString(s string) string {}
