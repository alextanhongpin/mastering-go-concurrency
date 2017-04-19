package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i, max int
	text   string
}

func outputText(j *Job, wg *sync.WaitGroup) {
	// defer wg.Done()
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	// Done tells the program the operation is complete
	wg.Done()
}

func main() {
	// wg is used to wait for the program to finish
	wg := new(sync.WaitGroup)
	fmt.Println("Starting")

	hello := &Job{0, 3, "Hello"}
	world := &Job{0, 5, "World"}

	// Run two goroutines
	go outputText(hello, wg)
	go outputText(world, wg)

	// Add a count of two, one for each goroutine
	wg.Add(2)

	// Wait for the goroutine to complete
	wg.Wait()
}
