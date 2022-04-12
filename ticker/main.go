package main

import (
	"fmt"
	"sync"
	"time"
)

func Tick1() {
	// create a ticker
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		fmt.Println("hi, tick1")
	}
}

func Tick2() {
	// create a ticker
	ticker := time.NewTicker(time.Second * 3)
	for range ticker.C {
		fmt.Println("hi, tick2")
	}
}

func Tick3() {
	// create a ticker
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		fmt.Println("hi, tick3")
		// if return, the ticker 3 will be shutdown
		// return
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		defer wg.Done()
		Tick1()
	}()

	go func() {
		defer wg.Done()
		Tick2()
	}()

	go func() {
		defer wg.Done()
		Tick3()
	}()

	wg.Wait()

}
