package main

import (
	"fmt"
	"sync"
)

// isPrime checks if a number is prime and sends it to the channel if it is.
func isPrime(num int, primeChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 2 and 3 are prime numbers
	if num == 2 || num == 3 {
		primeChan <- num
		return
	}

	// Prime-Number 6k ± 1 optimization:
	// All prime numbers are of the form 6k ± 1, except 2 and 3.
	// 6k + 1 means the number is 1 greater than a multiple of 6, hence num % 6 == 1.
	// 6k - 1 means the number is 1 less than a multiple of 6, hence num % 6 == 5.
	// Hence, if a number is not 6k ± 1, then it is not a prime number.
	if num%6 != 1 && num%6 != 5 {
		return
	}

	// Optimize loop for 6k ± 1 optimization:
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return
		}
	}
	primeChan <- num // It's a prime number
}

// findPrimesWithNumbers finds prime numbers from a given slice of numbers.
func findPrimesWithNumbers(numbers []int) []int {
	primeChan := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go isPrime(num, primeChan, &wg)
	}

	wg.Wait()
	close(primeChan)

	var primes []int
	for prime := range primeChan {
		primes = append(primes, prime)
	}
	return primes
}

// findPrimes finds all prime numbers up to a given limit.
func findPrimes(limit int) []int {
	primeChan := make(chan int)
	var wg sync.WaitGroup

	for num := 2; num <= limit; num++ {
		wg.Add(1)
		go isPrime(num, primeChan, &wg)
	}

	go func() {
		wg.Wait()
		close(primeChan)
	}()

	var primes []int
	for prime := range primeChan {
		primes = append(primes, prime)
	}
	return primes
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	primes := findPrimesWithNumbers(numbers)
	fmt.Printf("Primes in the given slice%+v: %v\n", numbers, primes)

	// https://en.wikipedia.org/wiki/Prime-counting_function
	limit := 100000
	ps := findPrimes(limit)
	fmt.Printf("Count of prime numbers up to %d: %d\n", limit, len(ps))
}
