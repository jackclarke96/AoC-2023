package main

type galaxyStruct struct {
	i, j int
}

type galaxyCoordinatesMap map[galaxyStruct]bool

type expansionColumns map[int]int
type expansionRows map[int]int // perhaps don't need both but does give clear separation

type calculationInput struct {
	galaxyStart, galaxyEnd                   galaxyStruct
	horizontalExpansions, verticalExpansions int
}
