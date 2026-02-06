package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	elements := make([]int, size)
	for i := range elements {
		elements[i] = r.Int()
	}
	return elements

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) < 1 {
		return 0
	}
	maxNumber := 0
	for i := range data {
		if data[i] > maxNumber {
			maxNumber = data[i]
		}
	}
	return maxNumber

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	dataLen := len(data)
	if dataLen < 1 {
		return 0
	}

	maxValues := make([]int, CHUNKS)
	chankSize := dataLen / CHUNKS

	for i := range CHUNKS {
		startIndex := i * chankSize
		endIndex := (1 + i) * chankSize
		if i == (CHUNKS - 1) {
			endIndex = (1+i)*chankSize + dataLen%CHUNKS
		}
		wg.Add(1)
		go func(sourceData []int, resultData []int, startIndex int, endIndex int) {
			defer wg.Done()
			maxNumber := maximum(data[startIndex:endIndex])
			resultData[i] = maxNumber
		}(data, maxValues, startIndex, endIndex)

	}
	wg.Wait()
	return maximum(maxValues)
}

func main() {

	elapsed := 0
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(numbers)
	end := time.Now()
	elapsed = int(end.Sub(start).Microseconds())

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(numbers)
	end = time.Now()
	elapsed = int(end.Sub(start).Microseconds())

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
