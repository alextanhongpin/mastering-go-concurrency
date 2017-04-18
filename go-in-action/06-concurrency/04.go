// This sample program demonstrates how the goroutine scheduler
// will time slice goroutines on a single thread
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

// main is the entry point for all Go programs
func main() {
	// Allocate 1 logical processors for the scheduler to use
	runtime.GOMAXPROCS(1)

	// Add a count of 2, 1 for each goroutine
	wg.Add(2)

	// Create two goroutines
	fmt.Println("Create goroutine")

	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime display prime numbers for the first 5000 numbers
func printPrime(prefix string) {
	// Schedule a call to Done to tell main we're done
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
