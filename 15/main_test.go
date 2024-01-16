package main

import "testing"

func TestProcessByteSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "makePrediction - test example input 1: rn=1",
			input:    []byte("rn=1"),
			expected: 30,
		},
		{
			name:     "makePrediction - test example input 2: cm-",
			input:    []byte("cm-"),
			expected: 253,
		},
		{
			name:     "makePrediction - test example input 3: qp=3",
			input:    []byte("qp=3"),
			expected: 97,
		},
		{
			name:     "makePrediction - test example input 4: cm=2",
			input:    []byte("cm=2"),
			expected: 47,
		},
		{
			name:     "makePrediction - test example input 5: qp-",
			input:    []byte("qp-"),
			expected: 14,
		},
		{
			name:     "makePrediction - test example input 6: pc=4",
			input:    []byte("pc=4"),
			expected: 180,
		},
		{
			name:     "makePrediction - test example input 7: ot=9",
			input:    []byte("ot=9"),
			expected: 9,
		},
		{
			name:     "makePrediction - test example input 8: ab=5",
			input:    []byte("ab=5"),
			expected: 197,
		},
		{
			name:     "makePrediction - test example input 9: pc-",
			input:    []byte("pc-"),
			expected: 48,
		},
		{
			name:     "makePrediction - test example input 10: pc=6",
			input:    []byte("pc=6"),
			expected: 214,
		},
		{
			name:     "makePrediction - test example input 11: ot=7",
			input:    []byte("ot=7"),
			expected: 231,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := processByteSlice(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
