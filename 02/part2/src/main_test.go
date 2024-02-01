package main

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestExecuteMain(t *testing.T) {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Could not read input file for test")
	}
	assert.Equal(t, 63711, executeMain(string(input)))
	assert.Equal(t, 2286, executeMain(testInput))
}

func TestChunkSlice(t *testing.T) {
	tests := []struct {
		name       string
		slice      []string
		chunkSize  int
		wantChunks [][]string
	}{
		{
			name:      "Normal Case",
			slice:     []string{"a", "b", "c", "d", "e", "f"},
			chunkSize: 2,
			wantChunks: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e", "f"},
			},
		},
		{
			name:      "Uneven Division",
			slice:     []string{"a", "b", "c", "d", "e"},
			chunkSize: 2,
			wantChunks: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e"},
			},
		},
		{
			name:       "Empty Slice",
			slice:      []string{},
			chunkSize:  3,
			wantChunks: [][]string{},
		},
		{
			name:      "Chunk Size Larger Than Slice",
			slice:     []string{"a", "b", "c"},
			chunkSize: 5,
			wantChunks: [][]string{
				{"a", "b", "c"},
			},
		},
		{
			name:      "Chunk Size One",
			slice:     []string{"a", "b"},
			chunkSize: 1,
			wantChunks: [][]string{
				{"a"},
				{"b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChunks := chunkSlice(tt.slice, tt.chunkSize); !reflect.DeepEqual(gotChunks, tt.wantChunks) {
				t.Errorf("chunkSlice() = %v, want %v", gotChunks, tt.wantChunks)
			}
		})
	}
}
