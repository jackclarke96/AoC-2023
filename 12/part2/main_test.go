package main

import (
	"strings"
	"testing"
)

var mainInput = `???.### 1,1,3 
.??..??...?##. 1,1,3 
?#?#?#?#?#?#?#? 1,3,1,6 
????.#...#... 4,1,1 
????.######..#####. 1,6,5 
?###???????? 3,2,1`

func TestExecuteMain(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "executeMain test input",
			input:    mainInput,
			expected: 525152,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := executeMain((tc.input))
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestGetCombinations(t *testing.T) {
	testCases := []struct {
		name     string
		arg1     []string
		arg2     []int
		expected int
	}{
		{
			name:     "TestGetCombinations - test example input 1",
			arg1:     strings.Split("???.###", ""),
			arg2:     []int{1, 1, 3},
			expected: 1,
		},
		{
			name:     "TestGetCombinations - test example input 2",
			arg1:     copySpringSlice(strings.Split(".??..??...?##.", "")),
			arg2:     copySpringLengthSlice([]int{1, 1, 3}),
			expected: 16384,
		},
		{
			name:     "TestGetCombinations - test example input 3",
			arg1:     copySpringSlice(strings.Split("?#?#?#?#?#?#?#?", "")),
			arg2:     copySpringLengthSlice([]int{1, 3, 1, 6}),
			expected: 1,
		},
		{
			name:     "TestGetCombinations - test example input 4",
			arg1:     copySpringSlice(strings.Split("????.#...#...", "")),
			arg2:     copySpringLengthSlice([]int{4, 1, 1}),
			expected: 16,
		},
		{
			name:     "TestGetCombinations - test example input 5",
			arg1:     copySpringSlice(strings.Split("????.######..#####.", "")),
			arg2:     copySpringLengthSlice([]int{1, 6, 5}),
			expected: 2500,
		},
		{
			name:     "TestGetCombinations - test example input 6",
			arg1:     copySpringSlice(strings.Split("?###????????", "")),
			arg2:     copySpringLengthSlice([]int{3, 2, 1}),
			expected: 506250,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateCombinations(tc.arg1, tc.arg2)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
