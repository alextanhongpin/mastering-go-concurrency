package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		log.Println("In default after:", time.Since(start))
	}
}
