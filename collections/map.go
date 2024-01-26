package collections

import (
	"runtime"
	"sync"
)

// mapFunc is a function that takes an item of type T and returns an item of type T.
type mapFunc[T any] func(T) T

// Map is a function that takes a slice of items and a mapFunc and returns a slice of items.
// Modify the items in place to avoid allocation.
func Map[T any](items []T, mf mapFunc[T]) []T {
	for i, v := range items {
		items[i] = mf(v)
	}
	return items
}

// ConcurrentMap is the Map function with concurrency.
func ConcurrentMap[T any](items []T, mf mapFunc[T]) []T {
	// Determine optimal chunk size based on CPU cores and workload
	numChunks := runtime.NumCPU()

	// Adjust chunk size if the number of items is small
	if len(items) < numChunks*1000 { // Adjust a threshold as needed
		numChunks = 1
	}

	chunkSize := (len(items) + numChunks - 1) / numChunks // Ensure even distribution

	// Create a buffered channel to avoid blocking, explicitly specifying its capacity
	resultCh := make(chan []T, numChunks) // Buffer channel to avoid deadlock

	var wg sync.WaitGroup
	for i := 0; i < numChunks; i++ {
		wg.Add(1)
		go func(startIndex, endIndex int) {
			defer wg.Done()
			// Use the original Map instead of creating a new one for each goroutine
			chunk := Map(items[startIndex:endIndex], mf)
			// Avoid sending nil if the chunk is empty
			if len(chunk) > 0 {
				resultCh <- chunk // Send processed chunk only if it's not empty
			}
		}(i*chunkSize, min(i*chunkSize+chunkSize, len(items)))
	}

	// Wait for all goroutines to finish before closing the channel
	wg.Wait()
	close(resultCh) // Explicitly close the channel after all sends are done

	results := make([]T, 0, len(items)) // Pre-allocate results slice for efficiency
	for chunk := range resultCh {
		results = append(results, chunk...)
	}

	return results
}
