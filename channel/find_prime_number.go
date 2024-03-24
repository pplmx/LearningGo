package main

import (
	"fmt"
	"math"
	"sync"
)

func isPrime(num int, primeChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// if num is 2, it is a prime number
	if num == 2 {
		primeChan <- num
		return
	}

	// if num is less than or equal to 1 or even number, it is not a prime number
	if num <= 1 || num%2 == 0 {
		return
	}

	// check if num is divisible by any odd number from 3 to sqrt(num)
	sqrt := int(math.Sqrt(float64(num)))
	for i := 3; i <= sqrt; i += 2 {
		if num%i == 0 {
			return
		}
	}

	// num is a prime number
	primeChan <- num
}

// findPrimesWithNumbers find prime numbers from given numbers
func findPrimesWithNumbers(numbers []int) []int {
	ch := make(chan int, len(numbers)) // Buffered channel
	var wg sync.WaitGroup
	var primes []int

	for _, num := range numbers {
		wg.Add(1)
		go isPrime(num, ch, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	close(ch) // Close the channel

	for prime := range ch {
		if prime != 0 { // Ignore zero values
			primes = append(primes, prime)
		}
	}
	return primes
}

// findPrimes find prime numbers since one to n number
func findPrimes(n int) []int {
	ch := make(chan int) // Unbuffered channel
	var wg sync.WaitGroup
	var primes []int

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go isPrime(i, ch, &wg)
	}

	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(ch) // Close the channel
	}()

	for prime := range ch {
		if prime != 0 { // Ignore zero values
			primes = append(primes, prime)
		}
	}
	return primes
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	primes := findPrimesWithNumbers(numbers)

	ps := findPrimes(100)
	fmt.Println(primes)
	fmt.Println(ps)
}
