package main

type linearFunction struct {
	minInput, maxInput, transform int
}

type piecewiseFunction []linearFunction

type inputRange struct {
	minInput, maxInput int
}
