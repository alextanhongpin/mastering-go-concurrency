package main

import (
	"fmt"
	"math/rand"
)

func main() {

	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			// This will not be closed, as it will be an infinite loop
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int():
			}
		}()
		return randStream
	}
	randStream := newRandStream()
	fmt.Println("3 random integer")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
