package main

import (
	"fmt"
)

func main() {
	// declare two bidirectional channels without buffer
	out := make(chan int)
	in := make(chan int)

	// Create 3 `multiplyByTwo` goroutines.
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	// Up till this point, none of the created goroutines actually do
	// anything, since they are all waiting for the `in` channel to
	// receive some data, we can send this in another goroutine
	go func() {
		in <- 1
		in <- 2
		in <- 3
		in <- 4
	}()

	// Now we wait for each result to come in
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

// multiplyByTwo the parameters are the two unidirectional channels for the type-safety of the program
func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	for {
		num := <-in
		result := num * 2
		out <- result
	}
}
