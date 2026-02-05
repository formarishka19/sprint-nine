package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const size = 100

var testSlice = []int{1, 4, 3, 0, 4}

func TestGenerateRandomElements(t *testing.T) {
	assert.Len(t, generateRandomElements(size), size)
}

func TestMaximum(t *testing.T) {
	assert.Equal(t, 4, maximum(testSlice))
}

func TestMaxChunks(t *testing.T) {
	testSlice := generateRandomElements(SIZE)
	assert.Equal(t, maximum(testSlice), maxChunks(testSlice))
}
