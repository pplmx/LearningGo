package collections

import (
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
	// Define the number of chunks
	numChunks := 4 // You can adjust this based on your needs

	// Define the size of each chunk
	chunkSize := len(items) / numChunks

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to receive processed chunks
	resultCh := make(chan []T, numChunks) // Buffer the channel to avoid blocking

	// Split the items into chunks and process each chunk concurrently
	for i := 0; i < numChunks; i++ {
		wg.Add(1)
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if i == numChunks-1 {
			endIndex = len(items)
		}

		go func(slice []T) {
			defer wg.Done()
			// Process the chunk using the provided mapFunc
			for j, v := range slice {
				slice[j] = mf(v)
			}
			// Send the processed chunk to the channel
			resultCh <- slice
		}(items[startIndex:endIndex])
	}

	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect the processed chunks from the channel and concatenate them
	var result []T
	for chunk := range resultCh {
		result = append(result, chunk...)
	}

	return result
}
