package main

type galaxyStruct struct {
	i, j int
}

type expansions map[int]int

type calculationInput struct {
	galaxyStart, galaxyEnd                   galaxyStruct
	horizontalExpansions, verticalExpansions expansions
}
