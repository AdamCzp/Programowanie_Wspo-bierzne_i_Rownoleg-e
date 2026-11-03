package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	const n = 1024

	// 1. Stworzenie tablicy i wypełnienie jej wartościami
	var tab [n]float64
	for i := 0; i < n; i++ {
		tab[i] = math.Sin(float64(i) * 2 * math.Pi / float64(n-1))
	}

	// Zadanie 1: Suma jednowątkowa
	sumaJednowatkova := 0.0
	for i := 0; i < n; i++ {
		sumaJednowatkova += tab[i]
	}
	fmt.Printf("Suma jednowątkowa: %e\n", sumaJednowatkova)

	// Zadanie 2: Suma wielowątkowa
	numThreads := 8
	chunkSize := n / numThreads

	var wg sync.WaitGroup
	wg.Add(numThreads)

	wynikiCzesciowe := make([]float64, numThreads)

	for t := 0; t < numThreads; t++ {
		threadIndex := t
		go func(id int) {
			defer wg.Done()

			start := id * chunkSize
			end := start + chunkSize

			sum := 0.0
			for i := start; i < end; i++ {
				sum += tab[i]
			}
			wynikiCzesciowe[id] = sum
		}(threadIndex)
	}

	wg.Wait()

	sumaWielowatkova := 0.0
	for _, val := range wynikiCzesciowe {
		sumaWielowatkova += val
	}

	fmt.Printf("Suma wielowątkowa (%d wątków): %e\n", numThreads, sumaWielowatkova)
}
