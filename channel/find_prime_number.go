package main

import (
	"fmt"
	"math"
	"sync"
)

// isPrime checks if a number is prime and sends it to the channel if it is.
func isPrime(num int, primeChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	if num <= 1 {
		return // Not a prime number
	}
	if num == 2 {
		primeChan <- num // The only even prime number
		return
	}
	if num%2 == 0 {
		return // Exclude even numbers
	}

	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 3; i <= sqrtNum; i += 2 {
		if num%i == 0 {
			return // Not a prime number
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
	fmt.Println("Primes in the given slice:", primes)

	ps := findPrimes(100)
	fmt.Println("Primes up to 100:", ps)
}
