package main

import (
	"log"
	"sync"
)

func main() {
	var data int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		data++
		wg.Done()
	}()

	wg.Wait()
	if data == 0 {
		log.Printf("the value is %d", data)
	} else {
		log.Printf("the value is %d", data)
	}
}
