package main

import (
	"log"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			log.Println("begun", i)
		}(i)
	}

	log.Println("unblocking goroutines")
	close(begin)
	wg.Wait()
}
