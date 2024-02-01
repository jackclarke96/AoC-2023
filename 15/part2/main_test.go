package main

import (
	"testing"
)

func TestEvaluateScore(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "executeMain - test example input",
			input:    []byte("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"),
			expected: 145,
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

func TestExtractLensOperation(t *testing.T) {
	testCases := []struct {
		name                string
		input               []byte
		expectedLabel       label
		expectedBoxNumber   int
		expectedFocalLength int
		expectedAction      string
	}{
		{
			name:                "extractLensOperation - test example input 1: rn=1",
			input:               []byte("rn=1"),
			expectedLabel:       "rn",
			expectedBoxNumber:   0,
			expectedFocalLength: 1,
			expectedAction:      "=",
		},
		{
			name:                "extractLensOperation - test example input 2: cm-",
			input:               []byte("cm-"),
			expectedLabel:       "cm",
			expectedBoxNumber:   0,
			expectedFocalLength: 0,
			expectedAction:      "-",
		},
		{
			name:                "extractLensOperation - test example input 3: qp=3",
			input:               []byte("qp=3"),
			expectedLabel:       "qp",
			expectedBoxNumber:   1,
			expectedFocalLength: 3,
			expectedAction:      "=",
		},
		{
			name:                "extractLensOperation - test example input 4: cm=2",
			input:               []byte("cm=2"),
			expectedLabel:       "cm",
			expectedBoxNumber:   0,
			expectedFocalLength: 2,
			expectedAction:      "=",
		},
		{
			name:                "extractLensOperation - test example input 5: qp-",
			input:               []byte("qp-"),
			expectedLabel:       "qp",
			expectedBoxNumber:   1,
			expectedFocalLength: 0,
			expectedAction:      "-",
		},
		{
			name:                "extractLensOperation - test example input 6: pc=4",
			input:               []byte("pc=4"),
			expectedLabel:       "pc",
			expectedBoxNumber:   3,
			expectedFocalLength: 4,
			expectedAction:      "=",
		},
		{
			name:                "extractLensOperation - test example input 7: ot=9",
			input:               []byte("ot=9"),
			expectedLabel:       "ot",
			expectedBoxNumber:   3,
			expectedFocalLength: 9,
			expectedAction:      "=",
		},

		{
			name:                "extractLensOperation - test example input 8: ab=5",
			input:               []byte("ab=5"),
			expectedLabel:       "ab",
			expectedBoxNumber:   3,
			expectedFocalLength: 5,
			expectedAction:      "=",
		},
		{
			name:                "extractLensOperation - test example input 9: pc-",
			input:               []byte("pc-"),
			expectedLabel:       "pc",
			expectedBoxNumber:   3,
			expectedFocalLength: 0,
			expectedAction:      "-",
		},
		{
			name:                "extractLensOperation - test example input 10: pc=6",
			input:               []byte("pc=6"),
			expectedLabel:       "pc",
			expectedBoxNumber:   3,
			expectedFocalLength: 6,
			expectedAction:      "=",
		},
		{
			name:                "extractLensOperation - test example input 11: ot=7",
			input:               []byte("ot=7"),
			expectedLabel:       "ot",
			expectedBoxNumber:   3,
			expectedFocalLength: 7,
			expectedAction:      "=",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			label, boxNumber, action, focalLength := extractLensOperation(&HASHMAP{}, tc.input)
			if label != tc.expectedLabel {
				t.Errorf("Failed %s: expected label %v, got %v", tc.name, tc.expectedLabel, label)
			}
			if boxNumber != tc.expectedBoxNumber {
				t.Errorf("Failed %s: expected box number %v, got %v", tc.name, tc.expectedBoxNumber, boxNumber)
			}
			if focalLength != tc.expectedFocalLength {
				t.Errorf("Failed %s: expected focal length %v, got %v", tc.name, tc.expectedFocalLength, focalLength)
			}
			if action != tc.expectedAction {
				t.Errorf("Failed %s: expected action %v, got %v", tc.name, tc.expectedAction, action)
			}
		})
	}
}
