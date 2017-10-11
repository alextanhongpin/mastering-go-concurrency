package main

import "log"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "hello channel"
	}()

	salutation, ok := <-stringStream
	log.Println(ok, salutation)
}
