package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	log.Println("blocking on read...")
	select {
	case <-c:
		log.Printf("unblocked %v later", time.Since(start))
	}
}
