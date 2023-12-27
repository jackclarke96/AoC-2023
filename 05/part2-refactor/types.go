package main

type SingleMapping []int
type FormattedMap []SingleMapping

type LinearEquation struct {
	minInput  int
	maxInput  int
	transform int
}

type InputRange struct {
	minInput int
	maxInput int
}

type PiecewiseFunction []LinearEquation
type FullPiecewiseMap []PiecewiseFunction
