package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

type concurrentSlice struct {
	mu     sync.Mutex
	values []int
}

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size > 0 {
		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		r := rand.New(source)

		elements := make([]int, size)
		for i := range elements {
			elements[i] = r.Int()
		}
		return elements
	}
	log.Fatal("slice size must be greater than 0")
	return nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) <= 1 {
		log.Fatal("slice size must be greater than 1")
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
	if CHUNKS < 1 {
		log.Fatal("number of chunks must be greater than 0")
		return 0
	}
	if CHUNKS > SIZE {
		log.Fatal("number of elements must be greater or equal than number of chunks")
		return 0
	}
	if len(data) != SIZE {
		log.Fatal("number of elements must equal SIZE constant")
		return 0
	}
	var wg sync.WaitGroup
	var maxValues concurrentSlice
	chankSize := SIZE / CHUNKS

	for i := range CHUNKS {
		startIndex := i * chankSize
		endIndex := (1 + i) * chankSize
		if i == (CHUNKS - 1) {
			endIndex = (1+i)*chankSize + SIZE%CHUNKS
		}
		wg.Add(1)
		go func(sourceData []int, resultData *concurrentSlice, startIndex int, endIndex int) {
			defer wg.Done()
			maxNumber := 0
			for j := startIndex; j < endIndex; j++ {
				if sourceData[j] > maxNumber {
					maxNumber = sourceData[j]
				}
			}
			resultData.mu.Lock()
			defer resultData.mu.Unlock()
			resultData.values = append(resultData.values, maxNumber)
		}(data, &maxValues, startIndex, endIndex)

	}
	wg.Wait()
	return maximum(maxValues.values)
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
