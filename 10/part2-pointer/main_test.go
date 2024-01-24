package main

import "testing"

func TestEvaluateScore(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "executeMain - test example input 1 - no squeezed through pipes",
			input: `...........
			.S-------7.
			.|F-----7|.
			.||.....||.
			.||.....||.
			.|L-7.F-J|.
			.|..|.|..|.
			.L--J.L--J.
			...........`,
			expected: 4,
		},
		{
			name: "executeMain - test example input 2 - squeezed through pipes included",
			input: `..........
				.S------7.
				.|F----7|.
				.||....||.
				.||....||.
				.|L-7F-J|.
				.|..||..|.
				.L--JL--J.
				..........`,
			expected: 4,
		},
		{
			name: "executeMain - test example input 3",
			input: `.F----7F7F7F7F-7....
				.|F--7||||||||FJ....
				.||.FJ||||||||L7....
				FJL7L7LJLJ||LJ.L-7..
				L--J.L7...LJS7F-7L7.
				....F-J..F7FJ|L7L7L7
				....L7.F7||L7|.L7L7|
				.....|FJLJ|FJ|F7|.LJ
				....FJL-7.||.||||...
				....L---J.LJ.LJLJ...`,
			expected: 8,
		},
		{
			name: "executeMain - test example input 4",
			input: `FF7FSF7F7F7F7F7F---7
				L|LJ||||||||||||F--J
				FL-7LJLJ||||||LJL-77
				F--JF--7||LJLJ7F7FJ-
				L---JF-JLJ.||-FJLJJ7
				|F|F-JF---7F7-L7L|7|
				|FFJF7L7F-JF7|JL---7
				7-L-JL7||F7|L7F-7F7|
				L.L7LFJ|||||FJL7||LJ
				L7JLJL-JLJLJL--JLJ.L`,
			expected: 10,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := executeMain(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
