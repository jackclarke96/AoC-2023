package main

import (
	"math"
	"testing"
)

func TestSafeSubtract(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive a positive b",
			a:        3,
			b:        2,
			expected: 1,
		},
		{
			name:     "positive a negative b",
			a:        3,
			b:        -2,
			expected: 5,
		},
		{
			name:     "negative a positive b",
			a:        -3,
			b:        2,
			expected: -5,
		},
		{
			name:     "negative a negative b",
			a:        -3,
			b:        -2,
			expected: -1,
		},
		{
			name:     "Near MaxInt with result above maxInt returns MaxInt",
			a:        math.MaxInt - 5,
			b:        -6,
			expected: math.MaxInt,
		},
		{
			name:     "Near MaxInt with result below maxInt returns Int",
			a:        math.MaxInt - 5,
			b:        -4,
			expected: math.MaxInt - 1,
		},
		{
			name:     "Near MinInt with result below MinInt returns MinInt",
			a:        math.MinInt + 5,
			b:        6,
			expected: math.MinInt,
		},
		{
			name:     "Near MinInt with result above MinInt returns Int",
			a:        math.MinInt + 5,
			b:        4,
			expected: math.MinInt + 1,
		},
		{
			name:     "MaxInt with result above maxInt returns MaxInt",
			a:        math.MaxInt,
			b:        -6,
			expected: math.MaxInt,
		},
		{
			name:     "MaxInt with result below maxInt returns Int",
			a:        math.MaxInt,
			b:        1,
			expected: math.MaxInt - 1,
		},
		{
			name:     "MinInt with result below MinInt returns MinInt",
			a:        math.MinInt,
			b:        6,
			expected: math.MinInt,
		},
		{
			name:     "MinInt with result above MinInt returns Int",
			a:        math.MinInt,
			b:        -1,
			expected: math.MinInt + 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := safeSubtract(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}

}
