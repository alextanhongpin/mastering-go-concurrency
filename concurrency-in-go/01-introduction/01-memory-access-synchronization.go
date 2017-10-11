package main

import (
	"log"
)

// Demonstrates the issue with data synchronization
// Three different possibilities
// 1. Nothing is printed, the data++ is called firs
// 2. The value 0 is printed, the if statement is called before goroutine
// 3. The value 1 is printed, the goroutine is called after if statement, and before print statement
//
// To detect data-race:
// $ go run -race main.go
func main() {
	var data int
	go func() { data++ }()

	if data == 0 {
		log.Printf("the value is %d", data)
	}
}
