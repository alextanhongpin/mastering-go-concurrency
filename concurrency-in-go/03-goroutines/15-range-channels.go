package main

import "log"

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		log.Println("populating channel")
		for i := 1; i <= 5; i++ {
			log.Println("<- wrote", i)
			intStream <- i
		}
		log.Println("done, closing channel")
	}()

	log.Println("reading from channel")
	for integer := range intStream {
		log.Println(integer, "<- read")
	}
}
