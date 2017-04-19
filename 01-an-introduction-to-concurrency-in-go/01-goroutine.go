package main

import (
	"fmt"
	"time"
)

// Job is the task to be done
type Job struct {
	i, max int
	text   string
}

// outputText prints n text
func outputText(j *Job) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
}

func main() {
	// Create a new job
	// Alternative way of writing: new(Job)
	hello := &Job{0, 3, "hello"}
	world := &Job{0, 5, "world"}

	// Run a goroutine that prints text every millisecond
	go outputText(hello)
	outputText(world)
}
