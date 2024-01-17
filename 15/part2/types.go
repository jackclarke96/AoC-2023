package main

// map for memoization linking lens labels to box numbers
type HASHMAP map[string]int

// use lensStruct with position field rather than a slice so that don't have to repeatedly (and inefficiently) append to slices
type lensStruct struct {
	focalLength int
	position    int
}

type label string

// Map each individual lens label contained within a box to its corresponding struct
type lensMap map[label](*lensStruct)

// Map each box number to the corresponding box containing each lens
type boxMap map[int]lensMap
