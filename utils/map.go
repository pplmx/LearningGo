package utils

import (
	"runtime"
)

func Map(strArr []string, f func(string) string) []string {
	mappedStrArr := make([]string, len(strArr))
	for i, v := range strArr {
		mappedStrArr[i] = f(v)
	}
	return mappedStrArr
}

// MapParallel applies function f to each element of strArr in parallel and returns a new slice containing the results.
func MapParallel(strArr []string, f func(string) string) []string {
	// Check if input is valid
	if strArr == nil || len(strArr) == 0 || f == nil {
		return nil
	}

	// Create a new slice to store the results
	mappedStrArr := make([]string, len(strArr))

	// Create a channel to synchronize goroutines
	ch := make(chan struct{})

	// Apply f to each element of strArr in parallel
	for i, v := range strArr {
		go func(i int, v string) {
			mappedStrArr[i] = f(v)
			ch <- struct{}{}
		}(i, v)
	}

	// Wait for all goroutines to complete
	for range strArr {
		<-ch
	}

	// Return the result
	return mappedStrArr
}

func MapParallel2(strArr []string, f func(string) string) []string {
	// Check if input is valid
	if strArr == nil || len(strArr) == 0 || f == nil {
		return nil
	}

	// Create a new slice to store the results
	mappedStrArr := make([]string, len(strArr))

	// Create a channel to synchronize goroutines
	ch := make(chan struct{})

	// Number of blocks
	numBlocks := runtime.NumCPU()

	blockSize := len(strArr) / numBlocks

	for i := 0; i < numBlocks; i++ {
		go func(i int) {
			start := i * blockSize
			end := start + blockSize
			if i == numBlocks-1 {
				end = len(strArr)
			}
			for j := start; j < end; j++ {
				mappedStrArr[j] = f(strArr[j])
			}
			ch <- struct{}{}
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < numBlocks; i++ {
		<-ch
	}

	return mappedStrArr
}
