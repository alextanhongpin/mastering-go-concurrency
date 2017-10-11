package main

import "log"

func main() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	log.Println(integer, ok)
}
