package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	size    int
	maximum int
	numbers []int
}

var testValues = []data{
	{size: 0, maximum: 0, numbers: []int{}},
	{size: 5, maximum: 4, numbers: []int{1, 4, 3, 0, 4}},
	{size: 5, maximum: 1, numbers: []int{1, 1, 1, 1, 1}},
	{size: 8, maximum: 100, numbers: []int{1, 2, 3, 4, 100, 7, 100, 1}},
	{size: 9, maximum: 1, numbers: []int{1, 1, 1, 1, 1, 0, 0, 0, 0}},
}

var incorrectSizeValues = []int{-1, 0}

var testSizes = []int{0, 1, 2, 8, 10, 100_000_001}

func TestGenerateRandomElements(t *testing.T) {
	for _, v := range testValues {
		assert.Len(t, generateRandomElements(v.size), v.size)
	}
	for _, v := range incorrectSizeValues {
		assert.Nil(t, generateRandomElements(v))
	}
}

func TestMaximum(t *testing.T) {
	for _, v := range testValues {
		assert.Equal(t, maximum(v.numbers), v.maximum)
	}
}

func TestMaxChunks(t *testing.T) {
	for _, size := range testSizes {
		testSlice := generateRandomElements(size)
		assert.Equal(t, maximum(testSlice), maxChunks(testSlice))
	}
	for _, v := range incorrectSizeValues {
		testSlice := generateRandomElements(v)
		assert.Equal(t, maxChunks(testSlice), 0)
	}

}
