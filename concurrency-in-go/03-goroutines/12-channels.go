package main

import "log"

func main() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		stringStream <- "Hello channels!"
	}()

	// This will cause a deadlock
	log.Println(<-stringStream)
}
