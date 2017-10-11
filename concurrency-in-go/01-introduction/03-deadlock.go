package main

import (
	"log"
	"sync"
	"time"
)

type value struct {
	sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.Lock()
		defer v1.Unlock()

		time.Sleep(1 * time.Second)
		v2.Lock()
		defer v2.Unlock()
		log.Printf("sum is: %d", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
