package main

const (
	square = "#"
	circle = "O"
	none   = "."
)

type grid [][]string

type direction string

const (
	N direction = "north"
	E direction = "east"
	S direction = "south"
	W direction = "west"
)
