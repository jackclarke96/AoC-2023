package main

import (
	"testing"
)

var baseGrid = grid{
	[]int{1, 0, 1, 1},
	[]int{0, 0, 0, 1},
	[]int{0, 0, 1, 0},
}

var expectedTransposed = grid{
	[]int{1, 0, 0},
	[]int{0, 0, 0},
	[]int{1, 0, 1},
	[]int{1, 1, 0},
}

var expectedReflected = grid{
	[]int{1, 1, 0, 1},
	[]int{1, 0, 0, 0},
	[]int{0, 1, 0, 0},
}

func areSlicesEqual(a, b grid) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestTransposingOperations(t *testing.T) {

	testCases := []struct {
		name     string
		receiver grid
		expected grid
	}{
		{
			name:     "Transposes the baseGrid",
			receiver: baseGrid,
			expected: expectedTransposed,
		},
		{
			name:     "Transposes the transpose",
			receiver: expectedTransposed,
			expected: baseGrid,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transposed := (tc.receiver).transpose()
			if !areSlicesEqual(transposed, tc.expected) {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, transposed)
			}
		})
	}
}

func TestReflectionOperations(t *testing.T) {

	testCases := []struct {
		name     string
		receiver grid
		expected grid
	}{
		{
			name:     "Reflects the baseGrid",
			receiver: baseGrid,
			expected: expectedReflected,
		},
		{
			name:     "Reflects the reflected",
			receiver: expectedReflected,
			expected: baseGrid,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reflected := (tc.receiver).reflectInVerticalPlane()
			if !areSlicesEqual(reflected, tc.expected) {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, reflected)
			}
		})
	}
}
