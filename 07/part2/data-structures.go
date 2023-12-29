package main

var cardToScoreMap = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

var handScoreMap = map[string]int{
	"FiveOfAKind":  7,
	"FourOfAKind":  6,
	"FullHouse":    5,
	"ThreeOfAKind": 4,
	"TwoPair":      3,
	"OnePair":      2,
	"HighCard":     1,
}

type SortedHandsResult struct {
	index int
	hands []Hand
}

type CountMap map[string]int

type Hand struct {
	cards       string
	cardsScores []int
	handScore   int
	bid         int
}
