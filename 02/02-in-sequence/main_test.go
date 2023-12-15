package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, executeMain("1", "../files/input.txt"), 2439)
	assert.Equal(t, executeMain("2", "../files/input.txt"), 63711)
}
