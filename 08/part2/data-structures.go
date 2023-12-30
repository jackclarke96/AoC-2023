package main

type LeftRight struct {
	L string
	R string
}

type NodeMap map[string]LeftRight

type TraverseGraphResult struct {
	startLocation string
	steps         int
	endLocation   string
}
