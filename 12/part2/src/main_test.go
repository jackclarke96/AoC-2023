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
			arg1:     springCombination(strings.Split("???.###", "")).unfoldSpringArrangement(),
			arg2:     springLengths{1, 1, 3}.unfoldSpringLengths(),
			expected: 1,
		},
		{
			name:     "TestGetCombinations - test example input 2",
			arg1:     springCombination(strings.Split(".??..??...?##.", "")).unfoldSpringArrangement(),
			arg2:     springLengths{1, 1, 3}.unfoldSpringLengths(),
			expected: 16384,
		},
		{
			name:     "TestGetCombinations - test example input 3",
			arg1:     springCombination(strings.Split("?#?#?#?#?#?#?#?", "")).unfoldSpringArrangement(),
			arg2:     springLengths{1, 3, 1, 6}.unfoldSpringLengths(),
			expected: 1,
		},
		{
			name:     "TestGetCombinations - test example input 4",
			arg1:     springCombination(strings.Split("????.#...#...", "")).unfoldSpringArrangement(),
			arg2:     springLengths{4, 1, 1}.unfoldSpringLengths(),
			expected: 16,
		},
		{
			name:     "TestGetCombinations - test example input 5",
			arg1:     springCombination(strings.Split("????.######..#####.", "")).unfoldSpringArrangement(),
			arg2:     springLengths{1, 6, 5}.unfoldSpringLengths(),
			expected: 2500,
		},
		{
			name:     "TestGetCombinations - test example input 6",
			arg1:     springCombination(strings.Split("?###????????", "")).unfoldSpringArrangement(),
			arg2:     springLengths{3, 2, 1}.unfoldSpringLengths(),
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
