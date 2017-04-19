// This program demonstrates how channels are created
// how to send/receive values through channel

package main

import (
	"fmt"
)

func main() {
	// Unbuffered channel of integers.
	// unbuffered := make(chan int)

	// Buffered channel of strings.
	buffered := make(chan string, 10)

	// Send a string through the channel.
	buffered <- "Gopher"

	// Receive a string from the channel.
	value := <-buffered
	fmt.Println(value)
}
