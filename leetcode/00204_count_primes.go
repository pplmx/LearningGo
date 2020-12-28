package main

import (
	"fmt"
	"math"
	"time"
)

func Multiplier() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}

func FindAllPrimes(n int) []int {
	/*
		find all primes in n (not include n)
	*/
	loopTimes := int(math.Sqrt(float64(n))) + 1
	var primeArr []int
	if n > 2 {
		primeArr = append(primeArr, 2)
	}
	for i := 3; i < n; i += 2 {
		isPrime := true
		for j := 3; j < loopTimes; j++ {
			if i != j && i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeArr = append(primeArr, i)
		}
	}
	return primeArr
}

func CountPrimes(n int) int {
	/*
		count all primes in n (not include n)
	*/
	loopTimes := int(math.Sqrt(float64(n))) + 1
	lenPrimes := 0
	if n > 2 {
		lenPrimes++
	}
	for i := 3; i < n; i += 2 {
		isPrime := true
		for j := 3; j < loopTimes; j++ {
			if i != j && i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			lenPrimes++
		}
	}
	return lenPrimes
}

func countPrimesUltimate(n int) int {
	var sum int
	// create a slice whose length is n
	f := make([]bool, n)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if f[i] == false {
			for j := i * i; j < n; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if f[i] == false {
			sum++
		}
	}
	return sum
}

func countPrimes3(n int) int {
	if n < 2 {
		return 0
	}
	dict := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !dict[i] {
			count++
			for j := 2; i*j < n; j++ {
				dict[i*j] = true
			}
		}
	}
	return count
}

func main() {
	NumMax := 16777216
	//start := time.Now()
	//primes := FindAllPrimes(NumMax)
	//elapsed := time.Since(start)
	//
	//start1 := time.Now()
	//lenPrimes := CountPrimes(NumMax)
	//elapsed1 := time.Since(start1)

	start2 := time.Now()
	lenPrimes2 := countPrimesUltimate(NumMax)
	elapsed2 := time.Since(start2)

	//start3 := time.Now()
	//lenPrimes3 := countPrimes3(NumMax)
	//elapsed3 := time.Since(start3)

	//fmt.Println(len(primes))
	//fmt.Println(elapsed)
	//
	//fmt.Println(lenPrimes)
	//fmt.Println(elapsed1)

	fmt.Println(lenPrimes2)
	fmt.Println(elapsed2)

	//fmt.Println(lenPrimes3)
	//fmt.Println(elapsed3)
}
